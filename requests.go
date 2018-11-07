package rpcclient

// Request ...
type Request interface {
	Params() []interface{}
	Method() string
}

// GetBlockVerbosityRequest ...
type GetBlockVerbosityRequest struct {
	Hash      string
	Verbosity int
}

// Params ...
func (r *GetBlockVerbosityRequest) Params() []interface{} {
	params := make([]interface{}, 2)
	params[0] = r.Hash
	params[1] = r.Verbosity
	return params
}

// Method get rpc method
func (r *GetBlockVerbosityRequest) Method() string {
	return "getblock"
}
