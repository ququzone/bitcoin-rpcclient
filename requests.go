package rpcclient

// Request ...
type Request interface {
	Params() []interface{}
	Method() string
	Response() interface{}
}

// GetBlockVerbosityRequest ...
type GetBlockVerbosityRequest struct {
	Hash      string
	Verbosity int
}

// Params get rpc params
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

// Response ...
func (r *GetBlockVerbosityRequest) Response() interface{} {
	return &BlockVerbosity{}
}

// GetBlockchainInfoRequest ...
type GetBlockchainInfoRequest struct {
}

// Params ...
func (*GetBlockchainInfoRequest) Params() []interface{} {
	return nil
}

// Method ...
func (*GetBlockchainInfoRequest) Method() string {
	return "getblockchaininfo"
}

// Response ...
func (*GetBlockchainInfoRequest) Response() interface{} {
	return &BlockchainInfo{}
}
