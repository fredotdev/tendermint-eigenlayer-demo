package main

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

// TxPayload represents a common payload for transactions
type TxPayload struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

// TxCompute represents the data needed to perform a computation
type TxCompute struct {
	Input []byte `json:"input"`
}

// sendTx sends the transaction to the Tendermint node
func sendTx(txPayload TxPayload) error {
	// Convert the txPayload to JSON
	txBytes, err := json.Marshal(txPayload)
	if err != nil {
		return err
	}

	// Encode the transaction data as base64
	txEncoded := base64.StdEncoding.EncodeToString(txBytes)

	// Encode the transaction data as a URL query parameter
	txQueryParam := url.QueryEscape(txEncoded)

	// Construct the URL with the query parameter
	urlString := fmt.Sprintf("http://localhost:26657/broadcast_tx_commit?tx=\"%s\"", txQueryParam)

	// Send the transaction to the Tendermint node's broadcast_tx_commit endpoint
	resp, err := http.Get(urlString)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Read the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Unmarshal the response JSON into a map
	var respJSON map[string]interface{}
	if err := json.Unmarshal(body, &respJSON); err != nil {
		return err
	}

	// Check if the transaction was successful
	if code, ok := respJSON["code"].(float64); ok && code != 0 {
		// Non-zero code indicates an error
		return fmt.Errorf("transaction failed with code %f: %s", code, string(body))
	}

	// Log the transaction hash if it's available
	if txHash, ok := respJSON["hash"].(string); ok {
		fmt.Printf("Transaction successful with hash: %s\n", txHash)
	}

	return nil
}

func main() {
	computeDataHex := os.Getenv("COMPUTE_DATA")
	if computeDataHex == "" {
		log.Fatal("Compute data not set in ENV")
	}
	computeData, err := hex.DecodeString(computeDataHex)
	if err != nil {
		log.Fatalf("Failed to decode compute data from hex: %v", err)
	}

	txCompute := TxCompute{
		Input: computeData,
	}
	txComputeData, err := json.Marshal(txCompute)
	if err != nil {
		log.Fatalf("Failed to marshal compute data: %v", err)
	}

	txPayload := TxPayload{
		Type: "compute",
		Data: json.RawMessage(txComputeData),
	}

	err = sendTx(txPayload)
	if err != nil {
		log.Fatalf("Failed to send compute transaction: %v", err)
	}
}
