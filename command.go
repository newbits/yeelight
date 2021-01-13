package yeelight

// Command represents COMMAND request to Yeelight device
type Command struct {
	ID     int           `json:"id"`
	Method string        `json:"method"`
	Params []interface{} `json:"params"`
}

// CommandResult represents response from Yeelight device
type CommandResult struct {
	ID     int           `json:"id"`
	Result []interface{} `json:"result,omitempty"`
	Error  *Error        `json:"error,omitempty"`
}
