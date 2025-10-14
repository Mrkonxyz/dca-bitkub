package utils

import (
	"Mrkonxyz/github.com/config"
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type ApiService struct {
	Cfg config.Config
}

func NewApiService(cgf config.Config) *ApiService {
	return &ApiService{Cfg: cgf}
}
func (a *ApiService) genSign(secret string, payloadString string) string {
	// Create a new HMAC hash using SHA-256
	h := hmac.New(sha256.New, []byte(secret))

	// Write the payload string (as bytes) to the HMAC hash
	h.Write([]byte(payloadString))

	// Return the hexadecimal representation of the HMAC
	return hex.EncodeToString(h.Sum(nil))
}
func (a *ApiService) getTimestamp() string {
	path := "/api/v3/servertime"
	url := fmt.Sprintf("%s%s", a.Cfg.BaseUrl, path)
	req, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error making HTTP request: %v", err)
	}
	res := a.readResponse(req.Body)
	return string(res)
}
func (a *ApiService) Get(url string) ([]byte, error) {
	req, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	res := a.readResponse(req.Body)

	return res, nil
}
func (a *ApiService) readResponse(r io.Reader) []byte {
	body, err := io.ReadAll(r)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}
	return body
}

func (a *ApiService) Post(url string, body *bytes.Buffer) (res []byte, err error) {
	req, err := http.NewRequest("POST", url, body)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	if err != nil {
		return
	}
	client := &http.Client{}

	response, err := client.Do(req)

	if err != nil {
		return
	}

	defer response.Body.Close()

	return a.readResponse(response.Body), nil
}

func genQueryParam(queryParams map[string]string) string {
	// Create a new url.Values object. It's the standard way to hold query params.
	params := url.Values{}

	// Loop through the map and add each key-value pair.
	// The .Set() method handles the assignment.
	for key, value := range queryParams {
		params.Set(key, value)
	}

	// .Encode() automatically handles URL encoding (e.g., spaces become %20)
	// and formats the string as "key1=value1&key2=value2".
	encodedParams := params.Encode()

	if encodedParams == "" {
		return ""
	}

	return "?" + encodedParams
}

func (a *ApiService) PostWithSig(path string, b *bytes.Buffer, params map[string]string) (response []byte, err error) {
	ts := a.getTimestamp()
	url := a.Cfg.BaseUrl + path

	var payload []string
	payload = append(payload, ts)
	payload = append(payload, "POST")
	payload = append(payload, path)
	var queryParam = ""
	if params != nil {
		queryParam = genQueryParam(params)
		payload = append(payload, queryParam)
	}
	if b != nil {
		payload = append(payload, b.String())
	}
	payloadStr := strings.Join(payload, "")
	sig := a.genSign(a.Cfg.ApiSecret, payloadStr)

	// Create a new GET request
	var body1 io.Reader = nil
	if b != nil {
		body1 = b
	}

	req, err := http.NewRequest("POST", url+queryParam, body1)
	if err != nil {
		log.Printf("Error creating request: %v \n", err)
	}
	// Optionally set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-BTK-APIKEY", a.Cfg.ApiKey)
	req.Header.Set("X-BTK-TIMESTAMP", ts)
	req.Header.Set("X-BTK-SIGN", sig)

	// Send the request using http.DefaultClient
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error making HTTP request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body := a.readResponse(resp.Body)
	// Check the status code
	if resp.StatusCode != http.StatusOK {
		log.Printf("Unexpected status code: %d, response: %s\n", resp.StatusCode, string(body))
	}

	return body, nil
}
