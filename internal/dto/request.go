package dto

type ShortenRequest struct {
	LongURL string `json:"long_url" validate:"required,url"`
	Expiry  int64  `json:"expiry",omitempty"`
}
