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

// BlockchainInfo ...
type BlockchainInfo struct {
	Chain                string  `json:"chain"`
	Blocks               uint64  `json:"blocks"`
	Headers              uint64  `json:"headers"`
	BestBlockHash        string  `json:"bestblockhash"`
	Difficulty           float64 `json:"difficulty"`
	MedianTime           int64   `json:"mediantime"`
	VerificationProgress float64 `json:"verificationprogress"`
	InitialBlockDownload bool    `json:"initialblockdownload"`
	ChainWork            string  `json:"chainwork"`
	SizeOnDisk           uint64  `json:"size_on_disk"`
	Pruned               bool    `json:"pruned"`
	SoftForks            []struct {
		ID      string `json:"id"`
		Version int    `json:"version"`
		Reject  struct {
			Status bool `json:"status"`
		} `json:"reject"`
	} `json:"softforks"`
	BIP9SoftForks struct {
		CSV struct {
			Status    string `json:"string"`
			StartTime int64  `json:"startTime"`
			Timeout   int64  `json:"timeout"`
			Since     int64  `json:"since"`
		} `json:"csv"`
		Segwit struct {
			Status    string `json:"string"`
			StartTime int64  `json:"startTime"`
			Timeout   int64  `json:"timeout"`
			Since     int64  `json:"since"`
		} `json:"segwit"`
	} `json:"bip9_softforks"`
}
