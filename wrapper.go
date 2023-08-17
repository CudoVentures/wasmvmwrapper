package api

import (
	wasmvmapi "github.com/CosmWasm/wasmvm/api"
	"github.com/CosmWasm/wasmvm/types"
)

const MOCK_CONTRACT_ADDR = wasmvmapi.MOCK_CONTRACT_ADDR

func NewMockGasMeter(limit uint64) wasmvmapi.MockGasMeter {
	return wasmvmapi.NewMockGasMeter(limit)
}

func NewLookup(meter wasmvmapi.MockGasMeter) *wasmvmapi.Lookup {
	return wasmvmapi.NewLookup(meter)
}

func NewMockAPI() *wasmvmapi.GoAPI {
	return wasmvmapi.NewMockAPI()
}

func DefaultQuerier(contractAddr string, coins types.Coins) types.Querier {
	return wasmvmapi.DefaultQuerier(contractAddr, coins)
}

func MockEnv() types.Env {
	return wasmvmapi.MockEnv()
}

func MockInfo(sender string, funds []types.Coin) types.MessageInfo {
	return wasmvmapi.MockInfo(sender, funds)
}
