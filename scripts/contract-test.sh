set -x

KEY="mykey"
CHAINID="ethermint_9000-1"
MONIKER="localtestnet"

shopt -s expand_aliases
alias reql="curl -X POST --url http://127.0.0.1:8545  -H 'Content-Type: application/json;' --data ${2} "

main() {
    # stop and remove existing daemon and client data and process(es)
    rm -rf ~/.ethermint*
    pkill -f "ethermint*"

    make build-ethermint

    # if $KEY exists it should be override
    "$PWD"/build/ethermintd keys add $KEY --keyring-backend test --algo "eth_secp256k1"

    # Set moniker and chain-id for Ethermint (Moniker can be anything, chain-id must be an integer)
    "$PWD"/build/ethermintd init $MONIKER --chain-id $CHAINID

    # Change parameter token denominations to aphoton
    cat $HOME/.ethermint/config/genesis.json | jq '.app_state["staking"]["params"]["bond_denom"]="stake"' >$HOME/.ethermint/config/tmp_genesis.json && mv $HOME/.ethermint/config/tmp_genesis.json $HOME/.ethermint/config/genesis.json
    cat $HOME/.ethermint/config/genesis.json | jq '.app_state["crisis"]["constant_fee"]["denom"]="aphoton"' >$HOME/.ethermint/config/tmp_genesis.json && mv $HOME/.ethermint/config/tmp_genesis.json $HOME/.ethermint/config/genesis.json
    cat $HOME/.ethermint/config/genesis.json | jq '.app_state["gov"]["deposit_params"]["min_deposit"][0]["denom"]="aphoton"' >$HOME/.ethermint/config/tmp_genesis.json && mv $HOME/.ethermint/config/tmp_genesis.json $HOME/.ethermint/config/genesis.json
    cat $HOME/.ethermint/config/genesis.json | jq '.app_state["mint"]["params"]["mint_denom"]="aphoton"' >$HOME/.ethermint/config/tmp_genesis.json && mv $HOME/.ethermint/config/tmp_genesis.json $HOME/.ethermint/config/genesis.json

    # Allocate genesis accounts (cosmos formatted addresses)
    "$PWD"/build/ethermintd add-genesis-account "$("$PWD"/build/ethermintd keys show "$KEY" -a --keyring-backend test)" 100000000000000000000aphoton,10000000000000000000stake --keyring-backend test

    # Sign genesis transaction
    "$PWD"/build/ethermintd gentx $KEY 10000000000000000000stake --amount=100000000000000000000aphoton --keyring-backend test --chain-id $CHAINID

    # Collect genesis tx
    "$PWD"/build/ethermintd collect-gentxs

    # Run this to ensure everything worked and that the genesis file is setup correctly
    "$PWD"/build/ethermintd validate-genesis

    # Start the node (remove the --pruning=nothing flag if historical queries are not needed) in background and log to file
    "$PWD"/build/ethermintd start --pruning=nothing --rpc.unsafe --json-rpc.address="127.0.0.1:8545" --keyring-backend test >ethermintd.log 2>&1 &

    # Give ethermintd node enough time to launch
    sleep 5

    solcjs --abi "$PWD"/tests/solidity/suites/basic/contracts/Counter.sol --bin -o "$PWD"/tests/solidity/suites/basic/counter
    mv "$PWD"/tests/solidity/suites/basic/counter/*.abi "$PWD"/tests/solidity/suites/basic/counter/counter_sol.abi 2>/dev/null
    mv "$PWD"/tests/solidity/suites/basic/counter/*.bin "$PWD"/tests/solidity/suites/basic/counter/counter_sol.bin 2>/dev/null

    # Query for the account
    ACCT=$(curl --fail --silent -X POST --data '{"jsonrpc":"2.0","method":"eth_accounts","params":[],"id":1}' -H "Content-Type: application/json" http://localhost:8545 | grep -o '\0x[^"]*')
    echo "$ACCT"

    # Start testcases (not supported)
    # curl -X POST --data '{"jsonrpc":"2.0","method":"personal_unlockAccount","params":["'$ACCT'", ""],"id":1}' -H "Content-Type: application/json" http://localhost:8545

    #PRIVKEY="$("$PWD"/build/ethermintd keys export $KEY)"

    ## need to get the private key from the account in order to check this functionality.
    cd tests/solidity/suites/basic/ && go get && go run main.go $ACCT

    # After tests
    # kill test ethermintd
    echo "going to shutdown ethermintd in 3 seconds..."
    sleep 3
    pkill -f "ethermint*"
}

probe() {
    # reql '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}' | jq .
    reql '{"jsonrpc":"2.0","method":"eth_getTransactionReceipt","params":["0x4465be1dc6fe6cb8327049905ed1bab6b2838c4dd3390921fa17d34d6d829995"],"id":1}' | jq .
    # Query for the account
    # curl --fail --silent -X POST --data '{"jsonrpc":"2.0","method":"eth_accounts","params":[],"id":1}' -H "Content-Type: application/json" http://localhost:8545

    # ACCT=$(curl --fail --silent -X POST --data '{"jsonrpc":"2.0","method":"eth_accounts","params":[],"id":1}' -H "Content-Type: application/json" http://localhost:8545 | grep -o '\0x[^"]*')
    # echo "$ACCT"
    # Start testcases (not supported)
    # curl -X POST --data '{"jsonrpc":"2.0","method":"personal_unlockAccount","params":["'$ACCT'", ""],"id":1}' -H "Content-Type: application/json" http://localhost:8545

}

bal() {
    # for addr in \
    #     $ADDR_DEPLOYER \
    #     $ADDR1 \
    #     0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266; do
    for addr in \
        0XF39FD6E51AAD88F6F4CE6AB8827279CFFFB92266 \
        0X70997970C51812DC3A010C7D01B50E0D17DC79C8 \
        0X3C44CDDDB6A900FA2B585DD299E03D12FA4293BC; do
        reql "{\"jsonrpc\":\"2.0\",\"method\":\"eth_getBalance\",\"params\":[\"$addr\", \"latest\"],\"id\":1}" | jq .
        # num=$(reql2 "{\"jsonrpc\":\"2.0\",\"method\":\"eth_getBalance\",\"params\":[\"$addr\"],\"id\":1}" | jq .result | tr -d '"')
        # cast --to-dec $num
        # cast --from-wei $num
    done
    return
}

$@
