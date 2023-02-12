package schema

// Error is a common struct that holds the
// error response returned
type Error struct {
	StatusCode    int    `json:"status_code,string"`
	StatusMessage string `json:"status_message"`
}
