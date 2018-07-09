// Copyright 2017 The go-ethereum Authors
// This file is part of go-ethereum.
//
// go-ethereum is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// go-ethereum is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with go-ethereum. If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"time"

	cli "gopkg.in/urfave/cli.v1"
	"minievm/common"
	"minievm/core"
	"minievm/core/state"
	"minievm/core/vm"
	"minievm/log"
	"minievm/params"
	"minievm/utils"
)

var runCommand = cli.Command{
	Action:      runCmd,
	Name:        "run",
	Usage:       "run arbitrary evm binary",
	ArgsUsage:   "<code>",
	Description: `The run command runs arbitrary EVM code.`,
}

func runCmd(ctx *cli.Context) error {
	glogger := log.NewGlogHandler(log.StreamHandler(os.Stderr, log.TerminalFormat(false)))
	log.Root().SetHandler(glogger)

	var (
		statedb     *state.StateDB
		account     = common.BytesToAddress([]byte("test"))
		blockNumber uint64
	)

	if ctx.GlobalString(AddressFlag.Name) != "" {
		account = common.HexToAddress(ctx.GlobalString(AddressFlag.Name))
	}

	statedb = state.New()
	statedb.CreateAccount(account)

	context := vm.Context{Transfer: core.Transfer, CanTransfer: core.CanTransfer, BlockNumber: new(big.Int).SetUint64(blockNumber)}
	evm := vm.NewEVM(context, statedb, params.MainnetChainConfig, vm.Config{EnableJit: false, ForceJit: false})

	var (
		code []byte
	)
	// The '--code' or '--codefile' flag overrides code in state
	if ctx.GlobalString(CodeFileFlag.Name) != "" {
		var hexcode []byte
		var err error
		// If - is specified, it means that code comes from stdin
		if ctx.GlobalString(CodeFileFlag.Name) == "-" {
			//Try reading from stdin
			if hexcode, err = ioutil.ReadAll(os.Stdin); err != nil {
				fmt.Printf("Could not load code from stdin: %v\n", err)
				os.Exit(1)
			}
		} else {
			// Codefile with hex assembly
			if hexcode, err = ioutil.ReadFile(ctx.GlobalString(CodeFileFlag.Name)); err != nil {
				fmt.Printf("Could not load code from file: %v\n", err)
				os.Exit(1)
			}
		}
		code = common.Hex2Bytes(string(bytes.TrimRight(hexcode, "\n")))

	} else if ctx.GlobalString(CodeFlag.Name) != "" {
		code = common.Hex2Bytes(ctx.GlobalString(CodeFlag.Name))
	}

	tstart := time.Now()
	initialGas := ctx.GlobalUint64(GasFlag.Name)
	value := utils.GlobalBig(ctx, ValueFlag.Name)
	input := ctx.GlobalString(InputFlag.Name)

	if len(code) > 0 {
		_, addr, _, _ := evm.Create(vm.AccountRef(account), code, initialGas, value)
		if len(input) > 0 {
			evm.Call(vm.AccountRef(account), addr, common.Hex2Bytes(input), initialGas, value)
		}
	}
	execTime := time.Since(tstart)

	if ctx.GlobalBool(StatDumpFlag.Name) {
		statedb.Print()
		fmt.Fprintf(os.Stderr, `evm execution time: %v`, execTime)
	}

	return nil
}
