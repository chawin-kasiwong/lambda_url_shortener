package dto

type ShortenResponse struct {
	ShortURL string `json:"short_url"`
}

func ErrorResponse(message string) map[string]string {
	return map[string]string{"error": message}
}
