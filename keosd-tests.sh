#!/bin/bash

KEOSD_PID=0
function finish {
    if [ $KEOSD_PID -ne 0 ]; then
        echo "Exiting cleanly"
        kill -9 $KEOSD_PID
    fi
}
trap finish EXIT

##################################################

WORKDIR="/tmp/eos-compat-tests/keosd/"
KEOSD_ADDRESS="127.0.0.1:4545"

##################################################

rm -rf "$WORKDIR"
mkdir -p "$WORKDIR"
echo "Started KEOSD on address $KEOSD_ADDRESS"
keosd --http-server-address "$KEOSD_ADDRESS" --wallet-dir "$WORKDIR" > /dev/null 2>&1 &
KEOSD_PID=$!
echo "Started KEOSD with PID $KEOSD_PID"
sleep 1

echo "Creating KEOSD wallet named $WALLET_NAME inside of $WORKDIR"
WALLET_PASS=$(curl -s "http://$KEOSD_ADDRESS/v1/wallet/create" --data '"default"' | tr -d '"')
echo "Fresh wallet passphrase: $WALLET_PASS"

# echo "Creating fresh key in the new wallet"
# NEW_PUB_KEY=$(curl -s "http://$KEOSD_ADDRESS/v1/wallet/create_key" --data '["default",""]' | tr -d '"')
# echo "New pub key : $NEW_PUB_KEY"

# NEW_PRIV_KEY=$(curl -s "http://$KEOSD_ADDRESS/v1/wallet/list_keys" --data "[\"default\",\"$WALLET_PASS\"]" | awk -F'"' '{print $4}')
# echo "Exported fresh private key : $NEW_PRIV_KEY"

# echo "Signing message digest ($DEFAULT_MESSAGE_DIGEST) with fresh key"
# SIGNATURE=$(curl -s "http://$KEOSD_ADDRESS/v1/wallet/sign_digest" --data "[\"$DEFAULT_MESSAGE_DIGEST\",\"$NEW_PUB_KEY\"]" | tr -d '"')
# echo "Signature : $SIGNATURE"

#########################################################
echo "#################################################################"

# Generated with cleos wallet create_key
TOIMPORT_PRIV_KEY="5JFhynQnFBYNTPDA9TiKeE7TmujNYaExcbZi9bsRUjhVxwZF4Mt"
TOIMPORT_PUB_KEY="EOS5jSQLpKBHLaMtuzkftnYE6bCMA5Jxso8f22uZyKj6cDEp32eSj"

echo "------------------------------------"
echo "Importing existing private key to wallet : $TOIMPORT_PRIV_KEY"
echo "Matching existing public key             : $TOIMPORT_PUB_KEY"
curl -s "http://$KEOSD_ADDRESS/v1/wallet/import_key" --data "[\"default\",\"$TOIMPORT_PRIV_KEY\"]" 

echo "------------------------------------"
MESSAGE_DIGEST="89529cb031c69eccc92f3e8492393a8688bd3d071d7346677b6ff59d314d5121" # Generated with `openssl rand -hex 32`
echo "Signing message digest ($MESSAGE_DIGEST) with $TOIMPORT_PUB_KEY"
SIGNATURE=$(curl -s "http://$KEOSD_ADDRESS/v1/wallet/sign_digest" --data "[\"$MESSAGE_DIGEST\",\"$TOIMPORT_PUB_KEY\"]" | tr -d '"')
echo "Signature : $SIGNATURE"

echo "------------------------------------"
MESSAGE_DIGEST="0000000000000000000000000000000000000000000000000000000000000000"
echo "Signing message digest ($MESSAGE_DIGEST) with $TOIMPORT_PUB_KEY"
SIGNATURE=$(curl -s "http://$KEOSD_ADDRESS/v1/wallet/sign_digest" --data "[\"$MESSAGE_DIGEST\",\"$TOIMPORT_PUB_KEY\"]" | tr -d '"')
echo "Signature : $SIGNATURE"

echo "------------------------------------"
MESSAGE_DIGEST="1111111111111111111111111111111111111111111111111111111111111111"
echo "Signing message digest ($MESSAGE_DIGEST) with $TOIMPORT_PUB_KEY"
SIGNATURE=$(curl -s "http://$KEOSD_ADDRESS/v1/wallet/sign_digest" --data "[\"$MESSAGE_DIGEST\",\"$TOIMPORT_PUB_KEY\"]" | tr -d '"')
echo "Signature : $SIGNATURE"

#########################################################
echo "#################################################################"

# Generated with cleos wallet create_key
TOIMPORT_PRIV_KEY="5Kk2STsBpo6UkY5Uw8BQ1YeFjp2BGLiBEsC5h4TYYiRDb7y5BTR"
TOIMPORT_PUB_KEY="EOS5gWrScGTTMyieGGhFDAmrVtDCp3UYzwdE7VLoZQnFSiGcezE3H"

echo "------------------------------------"
echo "Importing existing private key to wallet : $TOIMPORT_PRIV_KEY"
echo "Matching existing public key             : $TOIMPORT_PUB_KEY"
curl -s "http://$KEOSD_ADDRESS/v1/wallet/import_key" --data "[\"default\",\"$TOIMPORT_PRIV_KEY\"]" 

echo "------------------------------------"
MESSAGE_DIGEST="89529cb031c69eccc92f3e8492393a8688bd3d071d7346677b6ff59d314d5121" # Generated with `openssl rand -hex 32`
echo "Signing message digest ($MESSAGE_DIGEST) with $TOIMPORT_PUB_KEY"
SIGNATURE=$(curl -s "http://$KEOSD_ADDRESS/v1/wallet/sign_digest" --data "[\"$MESSAGE_DIGEST\",\"$TOIMPORT_PUB_KEY\"]" | tr -d '"')
echo "Signature : $SIGNATURE"

echo "------------------------------------"
MESSAGE_DIGEST="0000000000000000000000000000000000000000000000000000000000000000"
echo "Signing message digest ($MESSAGE_DIGEST) with $TOIMPORT_PUB_KEY"
SIGNATURE=$(curl -s "http://$KEOSD_ADDRESS/v1/wallet/sign_digest" --data "[\"$MESSAGE_DIGEST\",\"$TOIMPORT_PUB_KEY\"]" | tr -d '"')
echo "Signature : $SIGNATURE"

echo "------------------------------------"
MESSAGE_DIGEST="1111111111111111111111111111111111111111111111111111111111111111"
echo "Signing message digest ($MESSAGE_DIGEST) with $TOIMPORT_PUB_KEY"
SIGNATURE=$(curl -s "http://$KEOSD_ADDRESS/v1/wallet/sign_digest" --data "[\"$MESSAGE_DIGEST\",\"$TOIMPORT_PUB_KEY\"]" | tr -d '"')
echo "Signature : $SIGNATURE"

#######################################################
echo "#################################################################"

# Generated with cleos wallet create_key
TOIMPORT_PRIV_KEY="5HxQKWDznancXZXm7Gr2guadK7BhK9Zs8ejDhfA9oEBM89ZaAru"
TOIMPORT_PUB_KEY="EOS7dwvuZfiNdTbo3aamP8jgq8RD4kzauNkyiQVjxLtAhDHJm9joQ"

echo "------------------------------------"
echo "Importing existing private key to wallet : $TOIMPORT_PRIV_KEY"
echo "Matching existing public key             : $TOIMPORT_PUB_KEY"
curl -s "http://$KEOSD_ADDRESS/v1/wallet/import_key" --data "[\"default\",\"$TOIMPORT_PRIV_KEY\"]" 

echo "------------------------------------"
MESSAGE_DIGEST="6cb75bc5a46a7fdb64b92efefca01ed7b060ab5e0d625226e8efbc0980c3ddc1" # Generated with `openssl rand -hex 32`
echo "Signing message digest ($MESSAGE_DIGEST) with $TOIMPORT_PUB_KEY"
SIGNATURE=$(curl -s "http://$KEOSD_ADDRESS/v1/wallet/sign_digest" --data "[\"$MESSAGE_DIGEST\",\"$TOIMPORT_PUB_KEY\"]" | tr -d '"')
echo "Signature : $SIGNATURE"

