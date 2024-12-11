package api

import (
	"Mrkonxyz/github.com/config"
	"Mrkonxyz/github.com/model"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

var BaseUrl = config.AppConfig.BaseUrl

type ApiService struct {
	Cfg *config.Config
}

func NewApiService(cgf *config.Config) *ApiService {
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
	url := fmt.Sprintf("%s%s", config.AppConfig.BaseUrl, path)
	req, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error making HTTP request: %v", err)
	}
	res := a.ReadResponse(req.Body)
	return string(res)
}
func (a *ApiService) Get() {

}
func (a *ApiService) ReadResponse(r io.Reader) []byte {
	body, err := io.ReadAll(r)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}
	return body
}
func (a *ApiService) Post(path string) (response model.Response) {
	ts := a.getTimestamp()
	url := a.Cfg.BaseUrl + path

	// Create a new GET request
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	var payload []string
	payload = append(payload, ts)
	payload = append(payload, "POST")
	payload = append(payload, path)
	payloadStr := strings.Join(payload, "")

	sig := a.genSign(a.Cfg.ApiSecret, payloadStr)
	// Optionally set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-BTK-APIKEY", config.AppConfig.ApiKey)
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
	body := a.ReadResponse(resp.Body)

	// Check the status code
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Unexpected status code: %d, response: %s", resp.StatusCode, string(body))
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatalf("Error Unmarshal: %v", err)
	}
	return response
}
