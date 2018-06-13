package contract

import (
	"io/ioutil"
	"math/big"
	"minievm/common"
	"minievm/core"
	"minievm/params"
	"minievm/core/vm"
	st "minievm/core/state"
	. "github.com/gucumber/gucumber"
)

func init() {
	var evm *vm.EVM
	var addr common.Address
	state := st.New()

	Given(`^Create the evm$`, func() {
		addr = common.BytesToAddress([]byte("test1"))
		state.AddBalance(addr, big.NewInt(int64(1000000)))
		state.SetNonce(addr, uint64(20))
		context := vm.Context{Transfer: core.Transfer, CanTransfer: core.CanTransfer, BlockNumber: big.NewInt(4370001)}
		evm = vm.NewEVM(context, state, params.MainnetChainConfig, vm.Config{EnableJit: false, ForceJit: false})
	})

	And(`^Deploy the contract "(.+)?\.sol"$`, func(c string) {
		code, err := ioutil.ReadFile("sols/" + c + ".bin")
		if err != nil {
			T.Errorf("%s", err)
			return
		}
		_, e := createContract(evm, addr, common.Hex2Bytes(string(code)), big.NewInt(0))
		if e != nil {
			T.Errorf("create contract %s failed with %v", c, e)
		}
	})
}

func createContract(evm *vm.EVM, addr common.Address, code []byte, value *big.Int) (common.Address, error) {
	_, newAddr, _, err := evm.Create(vm.AccountRef(addr), code, uint64(100000000000), value)
	return newAddr, err
}

func callFunction(evm *vm.EVM, addr common.Address, to common.Address, input []byte, value *big.Int) error {
	_, _, err := evm.Call(vm.AccountRef(addr), to, input, uint64(100000000000), value)
	return err
}
