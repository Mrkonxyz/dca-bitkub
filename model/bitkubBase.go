package model

type BaseResultPage[T any] struct {
	Error      int64      `json:"error"`
	Result     []T        `json:"result"`
	Pagination Pagination `json:"pagination"`
}

type Pagination struct {
	Last int64       `json:"last"`
	Page interface{} `json:"page"`
}
