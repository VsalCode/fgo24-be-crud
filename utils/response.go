package utils

// Response struct untuk API response
type Response struct {
	Success  bool           `json:"success"`
	Message  string         `json:"message"`
	PageInfo map[string]any `json:"pageInfo,omitempty"`
	Result   interface{}    `json:"result,omitempty"`
}