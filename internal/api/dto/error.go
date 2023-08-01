package dto

// Error
//
// swagger:model
type Error struct {
	// Error code - for programmatic error handling
	//
	// example: internal_server_error
	// required: true
	Code string `json:"code"`
	// Error message - human readable
	//
	// example: Internal server error
	// required: true
	Message string `json:"message"`
	// Additional arguments
	//
	// example: {"arg1": "value1", "arg2": "value2"}
	// required: false
	Args map[string]string `json:"args,omitempty"`
}
