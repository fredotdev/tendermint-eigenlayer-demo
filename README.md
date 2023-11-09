export RPC_URL="https://goerli.infura.io/v3/"  
export CONTRACT_ADDRESS="0xaBF8a8e75011712FD0FEc3a15a2397F9d862763C"  
export BN254_PRIV_KEY="1000"
tendermint init --home .  
go build  
./example  


address can be found in broadcast register avs contract  



export ETH_PRIV_KEY="dbda1821b80551c9d65939329250298aa3472ba22feea921c0cf5d620ea67b97"  
go build scripts/tx_verify_signature.go  
./tx_verify_signature  

export COMPUTE_DATA="48656c6c6f2c20576f726c6421" 
go build scripts/tx_compute.go 
./tx_compute  


curl http://localhost:8080/sign  


