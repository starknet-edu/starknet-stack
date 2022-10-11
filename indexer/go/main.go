package main

import (
	"fmt"
	"log"
  	"net/http"
	"encoding/json"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Block struct {
	BlockHash       string `json:"block_hash"`
	ParentBlockHash string `json:"parent_block_hash"`
	BlockNumber     int    `json:"block_number"`
	StateRoot       string `json:"state_root"`
	Status          string `json:"status"`
	GasPrice        string `json:"gas_price"`
	Transactions    []struct {
		Version             string   `json:"version"`
		ContractAddress     string   `json:"contract_address"`
		EntryPointSelector  string   `json:"entry_point_selector,omitempty"`
		EntryPointType      string   `json:"entry_point_type,omitempty"`
		Calldata            []string `json:"calldata,omitempty"`
		Signature           []string `json:"signature,omitempty"`
		TransactionHash     string   `json:"transaction_hash"`
		MaxFee              string   `json:"max_fee,omitempty"`
		Type                string   `json:"type"`
		ContractAddressSalt string   `json:"contract_address_salt,omitempty"`
		ClassHash           string   `json:"class_hash,omitempty"`
		ConstructorCalldata []string `json:"constructor_calldata,omitempty"`
		Nonce               string   `json:"nonce,omitempty"`
	} `json:"transactions"`
	Timestamp           int    `json:"timestamp"`
	SequencerAddress    string `json:"sequencer_address"`
	TransactionReceipts []struct {
		TransactionIndex int           `json:"transaction_index"`
		TransactionHash  string        `json:"transaction_hash"`
		L2ToL1Messages   []interface{} `json:"l2_to_l1_messages"`
		Events           []struct {
			FromAddress string   `json:"from_address"`
			Keys        []string `json:"keys"`
			Data        []string `json:"data"`
		} `json:"events"`
		ExecutionResources struct {
			NSteps                 int `json:"n_steps"`
			BuiltinInstanceCounter struct {
				PedersenBuiltin   int `json:"pedersen_builtin"`
				RangeCheckBuiltin int `json:"range_check_builtin"`
				EcdsaBuiltin      int `json:"ecdsa_builtin"`
			} `json:"builtin_instance_counter"`
			NMemoryHoles int `json:"n_memory_holes"`
		} `json:"execution_resources"`
		ActualFee             string `json:"actual_fee"`
		L1ToL2ConsumedMessage struct {
			FromAddress string   `json:"from_address"`
			ToAddress   string   `json:"to_address"`
			Selector    string   `json:"selector"`
			Payload     []string `json:"payload"`
			Nonce       string   `json:"nonce"`
		} `json:"l1_to_l2_consumed_message,omitempty"`
	} `json:"transaction_receipts"`
	StarknetVersion string `json:"starknet_version"`
}

func main() {
	db, err := sql.Open("sqlite3", "./index.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("select address from erc20")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var address string
		err = rows.Scan(&address)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("ERC20 address: ", address)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	block := fetchLatestBlock()

	for _, txReceipt := range block.TransactionReceipts {
		fmt.Println("Block: ", txReceipt)
	}
}


func fetchLatestBlock() Block {
	url := "http://localhost:5050/feeder_gateway/get_block?blockNumber=latest"
	method := "GET"
  
	client := &http.Client {
	}
	req, err := http.NewRequest(method, url, nil)
  
	if err != nil {
	  panic(err)
	}
	res, err := client.Do(req)
	if err != nil {
	  panic(err)
	}
	defer res.Body.Close()
  
	var block Block
	err = json.NewDecoder(res.Body).Decode(&block)
	if err != nil {
	  panic(err)
	}
	 
	return block
  }
