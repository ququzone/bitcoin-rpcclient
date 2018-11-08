package rpcclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync/atomic"
)

// Client rpc http client
type Client struct {
	id         uint64
	Endpoint   string
	User       string
	Password   string
	httpClient *http.Client
}

type rawRequest struct {
	Jsonrpc string            `json:"jsonrpc"`
	Method  string            `json:"method"`
	Params  []json.RawMessage `json:"params"`
	ID      uint64            `json:"id"`
}

func marshalRawRequestJSON(id uint64, request Request) ([]byte, error) {
	rawParams := make([]json.RawMessage, 0, len(request.Params()))
	for _, param := range request.Params() {
		marshalledParam, err := json.Marshal(param)
		if err != nil {
			return nil, err
		}
		rawMessage := json.RawMessage(marshalledParam)
		rawParams = append(rawParams, rawMessage)
	}
	rawReq := rawRequest{
		Jsonrpc: "1.0",
		ID:      id,
		Method:  request.Method(),
		Params:  rawParams,
	}

	return json.Marshal(&rawReq)
}

type rpcError struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func (e rpcError) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}

type rawResponse struct {
	Result json.RawMessage `json:"result"`
	Error  *rpcError       `json:"error"`
}

func (r *rawResponse) result() (result []byte, err error) {
	if r.Error != nil {
		return nil, r.Error
	}
	return r.Result, nil
}

type response struct {
	result []byte
	err    error
}

// NewClient new http client
func NewClient(endpoint, user, password string) *Client {
	return &Client{
		Endpoint:   endpoint,
		User:       user,
		Password:   password,
		httpClient: &http.Client{},
	}
}

func (c *Client) nextID() uint64 {
	return atomic.AddUint64(&c.id, 1)
}

func (c *Client) sendPost(req Request) *response {
	rj, err := marshalRawRequestJSON(c.nextID(), req)
	if err != nil {
		return &response{result: nil, err: err}
	}
	bodyReader := bytes.NewReader(rj)
	httpReq, err := http.NewRequest("POST", c.Endpoint, bodyReader)
	if err != nil {
		return &response{result: nil, err: err}
	}
	httpReq.Close = true
	httpReq.Header.Set("Content-Type", "application/json")

	httpReq.SetBasicAuth(c.User, c.Password)

	httpResponse, err := c.httpClient.Do(httpReq)
	if err != nil {
		return &response{err: err}
	}

	respBytes, err := ioutil.ReadAll(httpResponse.Body)
	httpResponse.Body.Close()
	if err != nil {
		err = fmt.Errorf("error reading json reply: %v", err)
		return &response{err: err}
	}

	var resp rawResponse
	err = json.Unmarshal(respBytes, &resp)
	if err != nil {
		err = fmt.Errorf("status code: %d, response: %q",
			httpResponse.StatusCode, string(respBytes))
		return &response{err: err}
	}

	res, err := resp.result()
	return &response{result: res, err: err}
}

// GetBlockVerbosity ...
func (c *Client) GetBlockVerbosity(hash string, verbosity int) (*BlockVerbosity, error) {
	response := c.sendPost(&GetBlockVerbosityRequest{
		Hash:      hash,
		Verbosity: verbosity,
	})
	if response.err != nil {
		return nil, response.err
	}

	var result BlockVerbosity
	err := json.Unmarshal(response.result, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
