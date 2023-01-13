#!/bin/sh

CONTRACT_ADDR="0x00806778f9b06746fffd6ca567e0cfea9b3515432d9ba39928201d18c8dc9fdf"

while true
do
    TX_HASH=$(curl -s --location --request POST 'http://localhost:9545' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "jsonrpc":"2.0",
        "method":"starknet_getBlockWithTxs",
        "params":["latest"],
        "id":1
    }' | jq -r '.result.transactions[] | .contract_address + " " + .transaction_hash' | grep $CONTRACT_ADDR | awk '{print $2}')

    if [ -z "$TX_HASH" ]; then
        echo "no transaction for contract"
    else
        echo "found transaction: $TX_HASH"

        echo "Eevents to Index:"

        curl --location --request POST 'http://localhost:9545' \
            --header 'Content-Type: application/json' \
            --data-raw '{
                "jsonrpc":"2.0",
                "method":"starknet_getTransactionReceipt",
                "params":["0x05dc7fd94d54b4b1bf4229fe3b306164d5bab5cf2f4ac59b1fd8bbd5605244c5"],
                "id":1
            }' | jq -r '.result.events[]'
        
        exit 0
    fi
    sleep 60
done