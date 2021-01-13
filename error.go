package yeelight

// Error struct represents error part of response
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
