package yeelight

// Notification represents notification response
type Notification struct {
	Method string            `json:"method"`
	Params map[string]string `json:"params"`
}
