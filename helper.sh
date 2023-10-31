set -x

init() {
    apt install pipx
    pipx install toml-cli
}

info() {
    exec >"$FUNCNAME.log" 2>&1
    # cloc .
    find . -not \( -path */node_modules -prune \) -iname '*.md'
    find . -not \( -path */node_modules -prune \) -iname '*.sh'
    find . -not \( -path */node_modules -prune \) -iname '*make*'
    find . -not \( -path */node_modules -prune \) -iname '*docker*'
}

probe() {
    exec >"$FUNCNAME.log" 2>&1
    # ethermintd --help
    # ethermintd version
    ethermintd keys add --help
    ethermintd init --help
}

probeImage() {
    IMAGE=ghcr.io/b2network/b2-node:20231031-162216-eb3cc87
    docker run \
        --rm \
        -v /root/.ethermintd:/root/.ethermintd \
        -v $PWD:/host \
        -it $IMAGE \
        sh
}

debugImage() {
    ethermintd start \
        --metrics \
        --pruning=nothing \
        --rpc.unsafe \
        --keyring-backend test \
        --log_level info \
        --json-rpc.api eth,txpool,personal,net,debug,web3 \
        --api.enable
}

startOneNode(){
    docker-compose up -d node1
}

$@
