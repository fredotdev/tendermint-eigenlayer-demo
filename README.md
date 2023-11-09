# Tendermint Eigenlayer Demo
**Playground:**  
T1:
```
anvil --fork-url https://goerli.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161
```

T2:
```
export RPC_URL=http://localhost:8545
export PRIVATE_KEY=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80
./setup.sh
```

**Demo**  
T3:
```
export RPC_URL="http://localhost:8545"
export CONTRACT_ADDRESS="0xaBF8a8e75011712FD0FEc3a15a2397F9d862763C"
export BN254_PRIV_KEY="1000"
tendermint init --home .
go build
./example
```

**CONTRACT_ADDRESS:** can be found in playground dir ```broadcast/Operators.s.sol/5/registerOperatorsWithPlaygroundAVSFromConfigFile-latest.json```

**BN254_PRIV_KEY:** can be found in playground dir
```script/input/5/playground_avs_input.json```

**Demo scripts:**
```
export ETH_PRIV_KEY="dbda1821b80551c9d65939329250298aa3472ba22feea921c0cf5d620ea67b97"
go build scripts/tx_verify_signature.go
./tx_verify_signature
```
  
```
export COMPUTE_DATA="48656c6c6f2c20576f726c6421"
go build scripts/tx_compute.go
./tx_compute
```

**ETH_PRIV_KEY:** can be found in playground dir (no 0x prefix)
```script/input/5/playground_avs_input.json```

**COMPUTE_DATA:** compute input bytes (no 0x prefix)

**Express server hardcoded running on http://localhost:3000/compute**

**Custom curl:**
```
curl http://localhost:8080/sign
```
