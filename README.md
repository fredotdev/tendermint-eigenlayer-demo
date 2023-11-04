export RPC_URL="https://goerli.infura.io/v3/"
export CONTRAC_ADDRESS="0xaBF8a8e75011712FD0FEc3a15a2397F9d862763C"

tendermint init --home .
go build
./example


address can be found in broadcast register avs contract