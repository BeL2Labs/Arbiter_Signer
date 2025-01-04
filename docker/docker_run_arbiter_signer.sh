#!/bin/bash
read -p "please input Arbiter address: " arbiter_address
read -p "please input Arbiter btc private key: " arbiter_btc_private_key
read -p "please input Arbiter esc private key: " arbiter_esc_private_key

docker run \
    -e ARBITER_BTC_PRIVATE_KEY="$arbiter_btc_private_key" \
    -e ARBITER_ESC_PRIVATE_KEY="$arbiter_esc_private_key" \
    -e ARBITER_ADDRESS="$arbiter_address" \
    mollkeith/arbiter-signer:v0.0.1