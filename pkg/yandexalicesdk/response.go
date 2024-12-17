package yandexalicesdk

type Response struct {
	Version string   `json:"version"`
	Session struct{} `json:"session"`
	Result  Result   `json:"response"`
}
