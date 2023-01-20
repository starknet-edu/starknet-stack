#!/bin/sh

CONTRACT_ADDR="0x806778f9b06746fffd6ca567e0cfea9b3515432d9ba39928201d18c8dc9fdf"

while true
do
    BLOCK_NUM=$(curl -s --location --request POST 'http://localhost:9545' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "jsonrpc":"2.0",
        "method":"starknet_getBlockWithTxs",
        "params":["latest"],
        "id":1
    }' | jq -r '.result.block_number')
    echo "Pulled Block #: $BLOCK_NUM"

    TX_HASH=$(curl -s --location --request POST 'http://localhost:9545' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "jsonrpc":"2.0",
        "method":"starknet_getBlockWithTxs",
        "params":["latest"],
        "id":1
    }' | jq -r '.result.transactions[] | .calldata[] + " " + .transaction_hash' 2>/dev/null | grep $CONTRACT_ADDR | awk '{print $2}')

    if [ -z "$TX_HASH" ]; then
        echo "no transaction for contract\n"
    else
        echo "found transaction: $TX_HASH"

        echo "Events to Index:"

        curl -s --location --request POST 'http://localhost:9545' \
            --header 'Content-Type: application/json' \
            --data-raw '{
                "jsonrpc":"2.0",
                "method":"starknet_getTransactionReceipt",
                "params":["'$TX_HASH'"],
                "id":1
            }' | jq -r '.result.events'
        
        exit 0
    fi
    sleep 15
done
