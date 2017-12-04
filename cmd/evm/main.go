// Copyright 2014 The go-ethereum Authors
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

// evm executes EVM code snippets.
package main

import (
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/cmd/utils"
<<<<<<< HEAD
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/logger/glog"
	"github.com/ethereum/go-ethereum/params"
=======
>>>>>>> 1d06e41f04d75c31334c455063e9ec7b4136bf23
	"gopkg.in/urfave/cli.v1"
)

var gitCommit = "" // Git SHA1 commit hash of the release (set via linker flags)

var (
	app = utils.NewApp(gitCommit, "the evm command line interface")

	DebugFlag = cli.BoolFlag{
		Name:  "debug",
		Usage: "output full trace logs",
	}
	MemProfileFlag = cli.StringFlag{
		Name:  "memprofile",
		Usage: "creates a memory profile at the given path",
	}
	CPUProfileFlag = cli.StringFlag{
		Name:  "cpuprofile",
		Usage: "creates a CPU profile at the given path",
	}
	StatDumpFlag = cli.BoolFlag{
		Name:  "statdump",
		Usage: "displays stack and heap memory information",
	}
	CodeFlag = cli.StringFlag{
		Name:  "code",
		Usage: "EVM code",
	}
	CodeFileFlag = cli.StringFlag{
		Name:  "codefile",
		Usage: "File containing EVM code. If '-' is specified, code is read from stdin ",
	}
	GasFlag = cli.Uint64Flag{
		Name:  "gas",
		Usage: "gas limit for the evm",
		Value: 10000000000,
	}
	PriceFlag = utils.BigFlag{
		Name:  "price",
		Usage: "price set for the evm",
		Value: new(big.Int),
	}
	ValueFlag = utils.BigFlag{
		Name:  "value",
		Usage: "value set for the evm",
		Value: new(big.Int),
	}
	DumpFlag = cli.BoolFlag{
		Name:  "dump",
		Usage: "dumps the state after the run",
	}
	InputFlag = cli.StringFlag{
		Name:  "input",
		Usage: "input for the EVM",
	}
	VerbosityFlag = cli.IntFlag{
		Name:  "verbosity",
		Usage: "sets the verbosity level",
	}
	CreateFlag = cli.BoolFlag{
		Name:  "create",
		Usage: "indicates the action should be create rather than call",
	}
<<<<<<< HEAD
=======
	DisableGasMeteringFlag = cli.BoolFlag{
		Name:  "nogasmetering",
		Usage: "disable gas metering",
	}
	GenesisFlag = cli.StringFlag{
		Name:  "prestate",
		Usage: "JSON file with prestate (genesis) config",
	}
	MachineFlag = cli.BoolFlag{
		Name:  "json",
		Usage: "output trace logs in machine readable format (json)",
	}
	SenderFlag = cli.StringFlag{
		Name:  "sender",
		Usage: "The transaction origin",
	}
	ReceiverFlag = cli.StringFlag{
		Name:  "receiver",
		Usage: "The transaction receiver (execution context)",
	}
	DisableMemoryFlag = cli.BoolFlag{
		Name:  "nomemory",
		Usage: "disable memory output",
	}
	DisableStackFlag = cli.BoolFlag{
		Name:  "nostack",
		Usage: "disable stack output",
	}
>>>>>>> 1d06e41f04d75c31334c455063e9ec7b4136bf23
)

func init() {
	app.Flags = []cli.Flag{
		CreateFlag,
		DebugFlag,
		VerbosityFlag,
		CodeFlag,
		CodeFileFlag,
		GasFlag,
		PriceFlag,
		ValueFlag,
		DumpFlag,
		InputFlag,
		DisableGasMeteringFlag,
		MemProfileFlag,
		CPUProfileFlag,
		StatDumpFlag,
		GenesisFlag,
		MachineFlag,
		SenderFlag,
		ReceiverFlag,
		DisableMemoryFlag,
		DisableStackFlag,
	}
	app.Commands = []cli.Command{
		compileCommand,
		disasmCommand,
		runCommand,
		stateTestCommand,
	}
<<<<<<< HEAD
	app.Action = run
}

func run(ctx *cli.Context) error {
	glog.SetToStderr(true)
	glog.SetV(ctx.GlobalInt(VerbosityFlag.Name))

	db, _ := ethdb.NewMemDatabase()
	statedb, _ := state.New(common.Hash{}, db)
	sender := statedb.CreateAccount(common.StringToAddress("sender"))

	vmenv := NewEnv(statedb, common.StringToAddress("evmuser"), common.Big(ctx.GlobalString(ValueFlag.Name)), vm.Config{
		Debug:     ctx.GlobalBool(DebugFlag.Name),
		ForceJit:  ctx.GlobalBool(ForceJitFlag.Name),
		EnableJit: !ctx.GlobalBool(DisableJitFlag.Name),
	})

	tstart := time.Now()

	var (
		ret []byte
		err error
	)

	if ctx.GlobalBool(CreateFlag.Name) {
		input := append(common.Hex2Bytes(ctx.GlobalString(CodeFlag.Name)), common.Hex2Bytes(ctx.GlobalString(InputFlag.Name))...)
		ret, _, err = vmenv.Create(
			sender,
			input,
			common.Big(ctx.GlobalString(GasFlag.Name)),
			common.Big(ctx.GlobalString(PriceFlag.Name)),
			common.Big(ctx.GlobalString(ValueFlag.Name)),
		)
	} else {
		receiver := statedb.CreateAccount(common.StringToAddress("receiver"))

		code := common.Hex2Bytes(ctx.GlobalString(CodeFlag.Name))
		receiver.SetCode(crypto.Keccak256Hash(code), code)
		ret, err = vmenv.Call(
			sender,
			receiver.Address(),
			common.Hex2Bytes(ctx.GlobalString(InputFlag.Name)),
			common.Big(ctx.GlobalString(GasFlag.Name)),
			common.Big(ctx.GlobalString(PriceFlag.Name)),
			common.Big(ctx.GlobalString(ValueFlag.Name)),
		)
	}
	vmdone := time.Since(tstart)

	if ctx.GlobalBool(DumpFlag.Name) {
		statedb.Commit()
		fmt.Println(string(statedb.Dump()))
	}
	vm.StdErrFormat(vmenv.StructLogs())

	if ctx.GlobalBool(SysStatFlag.Name) {
		var mem runtime.MemStats
		runtime.ReadMemStats(&mem)
		fmt.Printf("vm took %v\n", vmdone)
		fmt.Printf(`alloc:      %d
tot alloc:  %d
no. malloc: %d
heap alloc: %d
heap objs:  %d
num gc:     %d
`, mem.Alloc, mem.TotalAlloc, mem.Mallocs, mem.HeapAlloc, mem.HeapObjects, mem.NumGC)
	}

	fmt.Printf("OUT: 0x%x", ret)
	if err != nil {
		fmt.Printf(" error: %v", err)
	}
	fmt.Println()
	return nil
=======
>>>>>>> 1d06e41f04d75c31334c455063e9ec7b4136bf23
}

func main() {
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
<<<<<<< HEAD

type VMEnv struct {
	state *state.StateDB
	block *types.Block

	transactor *common.Address
	value      *big.Int

	depth int
	Gas   *big.Int
	time  *big.Int
	logs  []vm.StructLog

	evm *vm.EVM
}

func NewEnv(state *state.StateDB, transactor common.Address, value *big.Int, cfg vm.Config) *VMEnv {
	env := &VMEnv{
		state:      state,
		transactor: &transactor,
		value:      value,
		time:       big.NewInt(time.Now().Unix()),
	}
	cfg.Logger.Collector = env

	env.evm = vm.New(env, cfg)
	return env
}

// ruleSet implements vm.RuleSet and will always default to the homestead rule set.
type ruleSet struct{}

func (ruleSet) IsHomestead(*big.Int) bool { return true }
func (ruleSet) GasTable(*big.Int) params.GasTable {
	return params.GasTableHomesteadGasRepriceFork
}

func (self *VMEnv) RuleSet() vm.RuleSet       { return ruleSet{} }
func (self *VMEnv) Vm() vm.Vm                 { return self.evm }
func (self *VMEnv) Db() vm.Database           { return self.state }
func (self *VMEnv) SnapshotDatabase() int     { return self.state.Snapshot() }
func (self *VMEnv) RevertToSnapshot(snap int) { self.state.RevertToSnapshot(snap) }
func (self *VMEnv) Origin() common.Address    { return *self.transactor }
func (self *VMEnv) BlockNumber() *big.Int     { return common.Big0 }
func (self *VMEnv) Coinbase() common.Address  { return *self.transactor }
func (self *VMEnv) Time() *big.Int            { return self.time }
func (self *VMEnv) Difficulty() *big.Int      { return common.Big1 }
func (self *VMEnv) BlockHash() []byte         { return make([]byte, 32) }
func (self *VMEnv) Value() *big.Int           { return self.value }
func (self *VMEnv) GasLimit() *big.Int        { return big.NewInt(1000000000) }
func (self *VMEnv) VmType() vm.Type           { return vm.StdVmTy }
func (self *VMEnv) Depth() int                { return 0 }
func (self *VMEnv) SetDepth(i int)            { self.depth = i }
func (self *VMEnv) GetHash(n uint64) common.Hash {
	if self.block.Number().Cmp(big.NewInt(int64(n))) == 0 {
		return self.block.Hash()
	}
	return common.Hash{}
}
func (self *VMEnv) AddStructLog(log vm.StructLog) {
	self.logs = append(self.logs, log)
}
func (self *VMEnv) StructLogs() []vm.StructLog {
	return self.logs
}
func (self *VMEnv) AddLog(log *vm.Log) {
	self.state.AddLog(log)
}
func (self *VMEnv) CanTransfer(from common.Address, balance *big.Int) bool {
	return self.state.GetBalance(from).Cmp(balance) >= 0
}
func (self *VMEnv) Transfer(from, to vm.Account, amount *big.Int) {
	core.Transfer(from, to, amount)
}

func (self *VMEnv) Call(caller vm.ContractRef, addr common.Address, data []byte, gas, price, value *big.Int) ([]byte, error) {
	self.Gas = gas
	return core.Call(self, caller, addr, data, gas, price, value)
}

func (self *VMEnv) CallCode(caller vm.ContractRef, addr common.Address, data []byte, gas, price, value *big.Int) ([]byte, error) {
	return core.CallCode(self, caller, addr, data, gas, price, value)
}

func (self *VMEnv) DelegateCall(caller vm.ContractRef, addr common.Address, data []byte, gas, price *big.Int) ([]byte, error) {
	return core.DelegateCall(self, caller, addr, data, gas, price)
}

func (self *VMEnv) Create(caller vm.ContractRef, data []byte, gas, price, value *big.Int) ([]byte, common.Address, error) {
	return core.Create(self, caller, data, gas, price, value)
}
=======
>>>>>>> 1d06e41f04d75c31334c455063e9ec7b4136bf23
