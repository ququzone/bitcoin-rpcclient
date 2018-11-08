package rpcclient

// Client rpc http client
type Client struct {
	Host     string
	User     string
	Password string
}

// NewClient new http client
func NewClient(host, user, password string) *Client {
	return &Client{
		Host:     host,
		User:     user,
		Password: password,
	}
}

// GetBlockVerbosity ...
func (c *Client) GetBlockVerbosity() (*BlockVerbosity, error) {
	return nil, nil
}
