#!/bin/bash
HOME_DIR="$HOME/.ethermint"
CHAINID="bsqhub_1113-1"
KEYRING="test"
KEYALGO="eth_secp256k1"
NODE=$HOME_DIR/node

KEYS=(
    "node0"
    "node1"
    "node2"
    "node3"
)

# validate dependencies are installed
command -v jq > /dev/null 2>&1 || { echo >&2 "jq not installed. More info: https://stedolan.github.io/jq/download/"; exit 1; }

# remove existing daemon and client
rm -rf ~/.ethermint*

make install

ethermintd config keyring-backend $KEYRING
ethermintd config chain-id $CHAINID

# if $KEY exists it should be deleted
# ethermintd keys add $KEY --keyring-backend $KEYRING --algo $KEYALGO

# Generate 4 nodes config
ethermintd testnet init-files --v 4 --chain-id $CHAINID --node-dir-prefix node --keyring-backend $KEYRING  --output-dir $HOME_DIR --algo $KEYALGO

# Copy each node's config to the temp directory
echo ">>> Copying each node's config to the temp directory"
mkdir -p $NODE/config $NODE/keyring-test
cp $HOME_DIR/node0/ethermintd/config/genesis.json $NODE/config/genesis.json
for i in {0..3}; do
    echo "Copying node$i's config to the temp directory"
    cp $HOME_DIR/node$i/ethermintd/keyring-test/node$i.info $NODE/keyring-test/node$i.info
done

# Change parameter token denominations to bsq
echo ">>> Change parameter token denominations to bsq"
jq '.app_state["staking"]["params"]["bond_denom"]="bsq"' $NODE/config/genesis.json > $NODE/config/tmp_genesis.json && mv $NODE/config/tmp_genesis.json $NODE/config/genesis.json
jq '.app_state["crisis"]["constant_fee"]["denom"]="bsq"' $NODE/config/genesis.json > $NODE/config/tmp_genesis.json && mv $NODE/config/tmp_genesis.json $NODE/config/genesis.json
jq '.app_state["gov"]["deposit_params"]["min_deposit"][0]["denom"]="bsq"' $NODE/config/genesis.json > $NODE/config/tmp_genesis.json && mv $NODE/config/tmp_genesis.json $NODE/config/genesis.json
jq '.app_state["mint"]["params"]["mint_denom"]="bsq"' $NODE/config/genesis.json > $NODE/config/tmp_genesis.json && mv $NODE/config/tmp_genesis.json $NODE/config/genesis.json

# Set gas limit in genesis
echo "Set gas limit in genesis"
jq '.consensus_params["block"]["max_gas"]="20000000"' $NODE/config/genesis.json > $NODE/config/tmp_genesis.json && mv $NODE/config/tmp_genesis.json $NODE/config/genesis.json

# Remove existing genesis accounts
echo ">>> Remove existing genesis accounts"
jq '.app_state["auth"]["accounts"]=[]' $NODE/config/genesis.json > $NODE/config/tmp_genesis.json && mv $NODE/config/tmp_genesis.json $NODE/config/genesis.json

# Remove existing genesis balanaces
echo "Remove existing genesis balanaces"
jq '.app_state["bank"]["balances"]=[]' $NODE/config/genesis.json > $NODE/config/tmp_genesis.json && mv $NODE/config/tmp_genesis.json $NODE/config/genesis.json

# Allocate genesis accounts (cosmos formatted addresses)
echo ">>> Allocate genesis accounts"
for i in {0..3}; do
    echo "Allocate genesis accounts for ${KEYS[i]}"
    ethermintd add-genesis-account ${KEYS[i]} 100000000000000000000000000bsq --keyring-backend $KEYRING --home $NODE
done

# Remove existing genesis gentxs
echo "Remove existing genesis gentxs"
jq '.app_state["genutil"]["gen_txs"]=[]' $NODE/config/genesis.json > $NODE/config/tmp_genesis.json && mv $NODE/config/tmp_genesis.json $NODE/config/genesis.json

# Sign genesis transaction
echo ">>> Sign genesis transaction"
for i in {0..3}; do
  echo "Sign genesis transaction for ${KEYS[i]}"
  ethermintd gentx ${KEYS[i]} 1000000000000000000000bsq --moniker node$i  --keyring-backend $KEYRING --chain-id $CHAINID --home $NODE
  mv $NODE/config/node_key.json $NODE/config/node${i}_key.json
  mv $NODE/config/priv_validator_key.json $NODE/config/priv_validator${i}_key.json
done

# Collect genesis tx
echo ">>> Collect genesis tx"
ethermintd collect-gentxs --home $NODE

# Run this to ensure everything worked and that the genesis file is setup correctly
ethermintd validate-genesis --home $NODE

# disable produce empty block and enable prometheus metrics
if [[ "$OSTYPE" == "darwin"* ]]; then
    sed -i '' 's/create_empty_blocks = true/create_empty_blocks = false/g' $NODE/config/config.toml
    sed -i '' 's/prometheus = false/prometheus = true/' $NODE/config/config.toml
    sed -i '' 's/prometheus-retention-time = 0/prometheus-retention-time  = 1000000000000/g' $NODE/config/app.toml
    sed -i '' 's/enabled = false/enabled = true/g' $NODE/config/app.toml
else
    sed -i 's/create_empty_blocks = true/create_empty_blocks = false/g' $NODE/config/config.toml
    sed -i 's/prometheus = false/prometheus = true/' $NODE/config/config.toml
    sed -i 's/prometheus-retention-time  = "0"/prometheus-retention-time  = "1000000000000"/g' $NODE/config/app.toml
    sed -i 's/enabled = false/enabled = true/g' $NODE/config/app.toml
fi

if [[ $1 == "pending" ]]; then
    echo "pending mode is on, please wait for the first block committed."
    if [[ "$OSTYPE" == "darwin"* ]]; then
        sed -i '' 's/create_empty_blocks_interval = "0s"/create_empty_blocks_interval = "30s"/g' $NODE/config/config.toml
        sed -i '' 's/timeout_propose = "3s"/timeout_propose = "30s"/g' $NODE/config/config.toml
        sed -i '' 's/timeout_propose_delta = "500ms"/timeout_propose_delta = "5s"/g' $NODE/config/config.toml
        sed -i '' 's/timeout_prevote = "1s"/timeout_prevote = "10s"/g' $NODE/config/config.toml
        sed -i '' 's/timeout_prevote_delta = "500ms"/timeout_prevote_delta = "5s"/g' $NODE/config/config.toml
        sed -i '' 's/timeout_precommit = "1s"/timeout_precommit = "10s"/g' $NODE/config/config.toml
        sed -i '' 's/timeout_precommit_delta = "500ms"/timeout_precommit_delta = "5s"/g' $NODE/config/config.toml
        sed -i '' 's/timeout_commit = "5s"/timeout_commit = "150s"/g' $NODE/config/config.toml
        sed -i '' 's/timeout_broadcast_tx_commit = "10s"/timeout_broadcast_tx_commit = "150s"/g' $NODE/config/config.toml
    else
        sed -i 's/create_empty_blocks_interval = "0s"/create_empty_blocks_interval = "30s"/g' $NODE/config/config.toml
        sed -i 's/timeout_propose = "3s"/timeout_propose = "30s"/g' $NODE/config/config.toml
        sed -i 's/timeout_propose_delta = "500ms"/timeout_propose_delta = "5s"/g' $NODE/config/config.toml
        sed -i 's/timeout_prevote = "1s"/timeout_prevote = "10s"/g' $NODE/config/config.toml
        sed -i 's/timeout_prevote_delta = "500ms"/timeout_prevote_delta = "5s"/g' $NODE/config/config.toml
        sed -i 's/timeout_precommit = "1s"/timeout_precommit = "10s"/g' $NODE/config/config.toml
        sed -i 's/timeout_precommit_delta = "500ms"/timeout_precommit_delta = "5s"/g' $NODE/config/config.toml
        sed -i 's/timeout_commit = "5s"/timeout_commit = "150s"/g' $NODE/config/config.toml
        sed -i 's/timeout_broadcast_tx_commit = "10s"/timeout_broadcast_tx_commit = "150s"/g' $NODE/config/config.toml
    fi
fi

# Copy the config files to the nodes' directories
echo ">>> Copy the config files to the nodes' directories"
for i in {0..3}; do
    echo "Copying node$i's config to the temp directory"
    cp $NODE/config/genesis.json $HOME_DIR/node$i/ethermintd/config/genesis.json
    cp $NODE/config/node${i}_key.json $HOME_DIR/node$i/ethermintd/config/node_key.json
    cp $NODE/config/priv_validator${i}_key.json $HOME_DIR/node$i/ethermintd/config/priv_validator_key.json
done