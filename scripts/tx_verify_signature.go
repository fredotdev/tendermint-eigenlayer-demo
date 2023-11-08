package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/base64"
	"strconv"
	"io/ioutil"
	"os"
	"log"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/tendermint/tendermint/crypto/ed25519"
)

// TxPayload represents a common payload for transactions
type TxPayload struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

// TxVerifySignature represents the data needed to verify a signature
type TxVerifySignature struct {
	From      string       `json:"from"`
	Message   string       `json:"message"`
	Signature string       `json:"signature"`
	PubKey    ed25519.PubKey `json:"pubKey"` // Public key in the appropriate format
}

// sendTx sends the transaction to the Tendermint node
func sendTx(txPayload TxPayload) error {
    // Convert the txPayload to JSON
    txBytes, err := json.Marshal(txPayload)
    if err != nil {
        return err
    }

    // The transaction might need to be encoded to base64 or some other encoding
    // depending on what Tendermint expects.
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

	// TODO: Check the Tendermint response for success or failure

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
        fmt.Printf("Raw response body: %s\n", string(body))

    }

    // Optionally, handle additional logic for success or logging here.
    // For example, you might want to log the transaction hash:
    if txHash, ok := respJSON["hash"].(string); ok {
        fmt.Printf("Transaction successful with hash: %s\n", txHash)
    }

    return nil
}

func main() {
	outputDir := "scripts/output"
    if _, err := os.Stat(outputDir); os.IsNotExist(err) {
        os.MkdirAll(outputDir, os.ModePerm)
    }

	privKeyHex := os.Getenv("ETH_PRIV_KEY")
	if privKeyHex == "" {
		log.Fatal("Ethereum private key not set in ENV")
	}
	privKey, err := crypto.HexToECDSA(privKeyHex)
	if err != nil {
		log.Fatalf("Invalid private key: %v", err)
	}

	// Generate the corresponding Ethereum public key
	pubKeyECDSA, ok := privKey.Public().(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Error casting public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*pubKeyECDSA).Hex()

	message := "YourMessage"
	prefix := []byte("\x19Ethereum Signed Message:\n" + strconv.Itoa(len(message)))
	prefixedMsg := append(prefix, []byte(message)...)
	messageHash := crypto.Keccak256Hash(prefixedMsg)
	signatureBytes, err := crypto.Sign(messageHash.Bytes(), privKey)
	if err != nil {
		log.Fatalf("Error signing the message: %v", err)
	}

	signature := hex.EncodeToString(signatureBytes)


	// Convert the signature to a format that Tendermint can understand
	// Note: Tendermint uses ed25519, Ethereum uses ECDSA, they are not directly compatible
	// You'll have to find a way to convert or work with ECDSA signatures within Tendermint

	// Log the generated keys and other important information to a file
    ethLogData := fmt.Sprintf("Address: %s\nPrivate Key: %s\nPublic Key: %x\nSignature: %s\n",
        fromAddress, privKeyHex, crypto.FromECDSAPub(pubKeyECDSA), signature)

    ethLogPath := "scripts/output/key_sign_log.txt" // Path updated to write inside scripts/output directory
    err = ioutil.WriteFile(ethLogPath, []byte(ethLogData), 0644)
    if err != nil {
        log.Fatalf("Error writing to Ethereum log file: %v", err)
    }

		// Generate an ed25519 keypair
    tmPrivKey := ed25519.GenPrivKey() // Use the tendermint ed25519 package's GenPrivKey function
    tmPubKey := tmPrivKey.PubKey()

    // Encode the private key to hex format for logging
    privKeyHex = hex.EncodeToString(tmPrivKey.Bytes()) // Note: now using tmPrivKey which is of type ed25519.PrivKey
    // Encode the public key to hex format for logging
    pubKeyHex := hex.EncodeToString(tmPubKey.Bytes())

    // Log the generated keys and other important information to a file
    tmLogData := fmt.Sprintf("Tendermint Public Key: %s\nTendermint Private Key: %s\n", pubKeyHex, privKeyHex)

    tmLogPath := "scripts/output/ed25519_keys.log" // Path updated to write inside scripts/output directory
    err = ioutil.WriteFile(tmLogPath, []byte(tmLogData), 0644)
    if err != nil {
        log.Fatalf("Error writing to Tendermint log file: %v", err)
    }

	// Create the TxVerifySignature data
	txData := TxVerifySignature{
		From:      fromAddress,
		Message:   message,
		Signature: signature,
		PubKey: tmPubKey.(ed25519.PubKey),
	}

	// Marshal the TxVerifySignature into JSON
	txDataBytes, err := json.Marshal(txData)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Sending JSON payload: %s\n", string(txDataBytes))

	// Create the TxPayload
	txPayload := TxPayload{
		Type: "verify_signature",
		Data: txDataBytes,
	}

	// Send the transaction
	err = sendTx(txPayload)
	if err != nil {
		fmt.Println("Error sending transaction:", err)
		os.Exit(1) // Exit the program with a non-zero status code to indicate failure
	}

	fmt.Println("Verify signature transaction sent successfully")
}
