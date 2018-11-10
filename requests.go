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

// GetBlockhashRequest ...
type GetBlockhashRequest struct {
	Height uint64
}

// Params ...
func (r *GetBlockhashRequest) Params() []interface{} {
	params := make([]interface{}, 1)
	params[0] = r.Height
	return params
}

// Method ...
func (*GetBlockhashRequest) Method() string {
	return "getblockhash"
}

// Response ...
func (*GetBlockhashRequest) Response() interface{} {
	s := ""
	return &s
}
