package rpcclient

// GetBlockVerbosity ...
func (c *Client) GetBlockVerbosity(hash string) (*BlockVerbosity, error) {
	result, err := c.sendRequest(&GetBlockVerbosityRequest{Hash: hash, Verbosity: 2})
	if err != nil {
		return nil, err
	}
	return result.(*BlockVerbosity), err
}

// GetBlockchainInfo ...
func (c *Client) GetBlockchainInfo() (*BlockchainInfo, error) {
	result, err := c.sendRequest(&GetBlockchainInfoRequest{})
	if err != nil {
		return nil, err
	}
	return result.(*BlockchainInfo), err
}

// GetBlockhash ...
func (c *Client) GetBlockhash(height uint64) (string, error) {
	result, err := c.sendRequest(&GetBlockhashRequest{height})
	if err != nil {
		return "", err
	}
	return result.(string), err
}
