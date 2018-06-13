package state

import (
	"fmt"
	"math/big"
	"minievm/common"
	"minievm/core/types"
	"strings"
)

type StateDB struct {
	StateMap map[common.Address]*State
}

type State struct {
	address   common.Address
	balance   *big.Int
	nounce    uint64
	codeHash  common.Hash
	code      []byte
	isSuicide bool
	storage   map[common.Hash]common.Hash
}

func New() *StateDB {
	return &StateDB{map[common.Address]*State{}}
}

func (st State) String() string {
	var lines []string

	lines = append(lines, fmt.Sprintf("--- State %x:", st.address))

	lines = append(lines, fmt.Sprintf("     balance: %d:", st.balance.Int64()))
	lines = append(lines, fmt.Sprintf("     nounce: %d", st.nounce))
	lines = append(lines, fmt.Sprintf("     codeHash: %x", st.codeHash))
	lines = append(lines, fmt.Sprintf("     code: %x", st.code))
	lines = append(lines, fmt.Sprintf("     isSuicide: %v", st.isSuicide))
	lines = append(lines, fmt.Sprintf("     storage:"))
	for k, v := range st.storage {
		lines = append(lines, fmt.Sprintf("       key: %x, val: %x", k, v))
	}

	return strings.Join(lines, "\n")
}

func (st *StateDB) Print() {
	fmt.Printf("\nstatedb:\n")
	for _, s := range st.StateMap {
		fmt.Printf("%s\n", s)
	}
}

func (st *StateDB) CreateAccount(addr common.Address) {
	st.StateMap[addr] = &State{addr, big.NewInt(0), 0, common.Hash{}, nil, false, map[common.Hash]common.Hash{}}
	return
}

func (st *StateDB) SubBalance(addr common.Address, value *big.Int) {
	if _, ok := st.StateMap[addr]; !ok {
		return
	}
	st.StateMap[addr].balance.Set(new(big.Int).Sub(st.StateMap[addr].balance, value))
	return
}
func (st *StateDB) AddBalance(addr common.Address, value *big.Int) {
	if _, ok := st.StateMap[addr]; !ok {
		st.StateMap[addr] = &State{addr, big.NewInt(0), 0, common.Hash{}, nil, false, map[common.Hash]common.Hash{}}
	}
	st.StateMap[addr].balance.Set(new(big.Int).Add(st.StateMap[addr].balance, value))
	return
}
func (st *StateDB) GetBalance(addr common.Address) *big.Int {
	if _, ok := st.StateMap[addr]; !ok {
		return nil
	}
	return st.StateMap[addr].balance
}

func (st *StateDB) GetNonce(addr common.Address) uint64 {
	if _, ok := st.StateMap[addr]; !ok {
		return 0
	}
	return st.StateMap[addr].nounce
}
func (st *StateDB) SetNonce(addr common.Address, value uint64) {
	if _, ok := st.StateMap[addr]; !ok {
		st.StateMap[addr] = &State{addr, big.NewInt(0), 0, common.Hash{}, nil, false, map[common.Hash]common.Hash{}}
	}
	st.StateMap[addr].nounce = value
	return
}

func (st *StateDB) GetCodeHash(addr common.Address) common.Hash {
	if _, ok := st.StateMap[addr]; !ok {
		return common.Hash{}
	}
	return st.StateMap[addr].codeHash
}
func (st *StateDB) GetCode(addr common.Address) []byte {
	if _, ok := st.StateMap[addr]; !ok {
		return nil
	}
	return st.StateMap[addr].code
}
func (st *StateDB) SetCode(addr common.Address, code []byte) {
	if _, ok := st.StateMap[addr]; !ok {
		st.StateMap[addr] = &State{addr, big.NewInt(0), 0, common.Hash{}, nil, false, map[common.Hash]common.Hash{}}
	}
	st.StateMap[addr].code = code
	return
}
func (st *StateDB) GetCodeSize(addr common.Address) int {
	if _, ok := st.StateMap[addr]; !ok {
		return 0
	}
	code := st.StateMap[addr].code
	return len(code)
}

func (st *StateDB) AddRefund(uint64) {
	return
}
func (st *StateDB) GetRefund() uint64 {
	return 0
}

func (st *StateDB) GetState(addr common.Address, key common.Hash) common.Hash {
	if _, ok := st.StateMap[addr]; !ok {
		st.StateMap[addr] = &State{addr, big.NewInt(0), 0, common.Hash{}, nil, false, map[common.Hash]common.Hash{}}
	}
	if _, ok := st.StateMap[addr].storage[key]; !ok {
		return common.Hash{}
	}
	return st.StateMap[addr].storage[key]
}
func (st *StateDB) SetState(addr common.Address, key common.Hash, value common.Hash) {
	if _, ok := st.StateMap[addr]; !ok {
		st.StateMap[addr] = &State{addr, big.NewInt(0), 0, common.Hash{}, nil, false, map[common.Hash]common.Hash{}}
	}
	st.StateMap[addr].storage[key] = value
	return
}

func (st *StateDB) Suicide(addr common.Address) bool {
	if _, ok := st.StateMap[addr]; !ok {
		st.StateMap[addr] = &State{addr, big.NewInt(0), 0, common.Hash{}, nil, false, map[common.Hash]common.Hash{}}
	}
	st.StateMap[addr].isSuicide = true
	return true
}
func (st *StateDB) HasSuicided(addr common.Address) bool {
	if _, ok := st.StateMap[addr]; !ok {
		st.StateMap[addr] = &State{addr, big.NewInt(0), 0, common.Hash{}, nil, false, map[common.Hash]common.Hash{}}
	}
	return st.StateMap[addr].isSuicide
}

func (st *StateDB) Exist(common.Address) bool {
	return true
}

func (st *StateDB) Empty(common.Address) bool {
	return true
}

func (st *StateDB) RevertToSnapshot(int) {
	return
}
func (st *StateDB) Snapshot() int {
	return 0
}

func (st *StateDB) AddLog(*types.Log) {
	return
}
func (st *StateDB) AddPreimage(common.Hash, []byte) {
	return
}

func (st *StateDB) ForEachStorage(common.Address, func(common.Hash, common.Hash) bool) {
	return
}
