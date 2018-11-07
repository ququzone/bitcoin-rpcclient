package rpcclient

import "encoding/json"

// Response ...
type Response struct {
	ID     int32           `json:"id"`
	Result json.RawMessage `json:"result"`
	Error  struct {
		Code    int32  `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

// BlockVerbosity ...
type BlockVerbosity struct {
	Hash          string `json:"hash"`
	Confirmations int64  `json:"confirmations"`
	StrippedSize  int64  `json:"strippedsize"`
	Size          int64  `json:"size"`
	Weight        int64  `json:"weight"`
	Height        int64  `json:"height"`
	Version       int    `json:"version"`
	VersionHex    string `json:"versionHex"`
	MerkleRoot    string `json:"merkleroot"`
	Tx            []struct {
		TxID     string `json:"txid"`
		Hash     string `json:"hash"`
		Version  int    `json:"version"`
		Size     int64  `json:"size"`
		VSize    int64  `json:"vsize"`
		LockTime int64  `json:"locktime"`
		Vin      []struct {
			Coinbase  string `json:"coinbase"`
			TxID      string `json:"txid"`
			Vout      int    `json:"vout"`
			ScriptSig struct {
				Asm string `json:"asm"`
				Hex string `json:"hex"`
			} `json:"scriptSig"`
			Sequence int64 `json:"sequence"`
		} `json:"vin"`
		Vout []struct {
			Value        float64 `json:"value"`
			N            int     `json:"n"`
			ScriptPubKey struct {
				Asm       string   `json:"asm"`
				Hex       string   `json:"hex"`
				RegSigs   int      `json:"reqSigs"`
				Type      string   `json:"type"`
				Addresses []string `json:"addresses"`
			} `json:"scriptPubKey"`
		} `json:"vout"`
		Hex string `json:"hex"`
	} `json:"tx"`
	Time              int64   `json:"time"`
	MedianTime        int64   `json:"mediantime"`
	Nonce             uint64  `json:"nonce"`
	Bits              string  `json:"bits"`
	Difficulty        float64 `json:"difficulty"`
	ChainWork         string  `json:"chainwork"`
	NTx               int     `json:"nTx"`
	PreviousBlockHash string  `json:"previousblockhash"`
	NextBlockHash     string  `json:"nextblockhash"`
}
