// Copyright 2017 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package tests

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
<<<<<<< HEAD
)

func BenchmarkStateCall1024(b *testing.B) {
	fn := filepath.Join(stateTestDir, "stCallCreateCallCodeTest.json")
	if err := BenchVmTest(fn, bconf{"Call1024BalanceTooLow", true, os.Getenv("JITVM") == "true"}, b); err != nil {
		b.Error(err)
	}
}

func TestStateSystemOperations(t *testing.T) {
	ruleSet := RuleSet{
		HomesteadBlock: big.NewInt(1150000),
	}

	fn := filepath.Join(stateTestDir, "stSystemOperationsTest.json")
	if err := RunStateTest(ruleSet, fn, StateSkipTests); err != nil {
		t.Error(err)
	}
}

func TestStateExample(t *testing.T) {
	ruleSet := RuleSet{
		HomesteadBlock: big.NewInt(1150000),
	}

	fn := filepath.Join(stateTestDir, "stExample.json")
	if err := RunStateTest(ruleSet, fn, StateSkipTests); err != nil {
		t.Error(err)
	}
}

func TestStatePreCompiledContracts(t *testing.T) {
	ruleSet := RuleSet{
		HomesteadBlock: big.NewInt(1150000),
	}

	fn := filepath.Join(stateTestDir, "stPreCompiledContracts.json")
	if err := RunStateTest(ruleSet, fn, StateSkipTests); err != nil {
		t.Error(err)
	}
}

func TestStateRecursiveCreate(t *testing.T) {
	ruleSet := RuleSet{
		HomesteadBlock: big.NewInt(1150000),
	}

	fn := filepath.Join(stateTestDir, "stRecursiveCreate.json")
	if err := RunStateTest(ruleSet, fn, StateSkipTests); err != nil {
		t.Error(err)
	}
}

func TestStateSpecial(t *testing.T) {
	ruleSet := RuleSet{
		HomesteadBlock: big.NewInt(1150000),
	}

	fn := filepath.Join(stateTestDir, "stSpecialTest.json")
	if err := RunStateTest(ruleSet, fn, StateSkipTests); err != nil {
		t.Error(err)
	}
}

func TestStateRefund(t *testing.T) {
	ruleSet := RuleSet{
		HomesteadBlock: big.NewInt(1150000),
	}

	fn := filepath.Join(stateTestDir, "stRefundTest.json")
	if err := RunStateTest(ruleSet, fn, StateSkipTests); err != nil {
		t.Error(err)
	}
}

func TestStateBlockHash(t *testing.T) {
	ruleSet := RuleSet{
		HomesteadBlock: big.NewInt(1150000),
	}

	fn := filepath.Join(stateTestDir, "stBlockHashTest.json")
	if err := RunStateTest(ruleSet, fn, StateSkipTests); err != nil {
		t.Error(err)
	}
}

func TestStateInitCode(t *testing.T) {
	ruleSet := RuleSet{
		HomesteadBlock: big.NewInt(1150000),
	}

	fn := filepath.Join(stateTestDir, "stInitCodeTest.json")
	if err := RunStateTest(ruleSet, fn, StateSkipTests); err != nil {
		t.Error(err)
	}
}

func TestStateLog(t *testing.T) {
	ruleSet := RuleSet{
		HomesteadBlock: big.NewInt(1150000),
	}

	fn := filepath.Join(stateTestDir, "stLogTests.json")
	if err := RunStateTest(ruleSet, fn, StateSkipTests); err != nil {
		t.Error(err)
	}
}

func TestStateTransaction(t *testing.T) {
	ruleSet := RuleSet{
		HomesteadBlock: big.NewInt(1150000),
	}

	fn := filepath.Join(stateTestDir, "stTransactionTest.json")
	if err := RunStateTest(ruleSet, fn, StateSkipTests); err != nil {
		t.Error(err)
	}
}

func TestStateTransition(t *testing.T) {
	ruleSet := RuleSet{
		HomesteadBlock: big.NewInt(1150000),
	}

	fn := filepath.Join(stateTestDir, "stTransitionTest.json")
	if err := RunStateTest(ruleSet, fn, StateSkipTests); err != nil {
		t.Error(err)
	}
}

func TestCallCreateCallCode(t *testing.T) {
	ruleSet := RuleSet{
		HomesteadBlock: big.NewInt(1150000),
	}

	fn := filepath.Join(stateTestDir, "stCallCreateCallCodeTest.json")
	if err := RunStateTest(ruleSet, fn, StateSkipTests); err != nil {
		t.Error(err)
	}
}

func TestCallCodes(t *testing.T) {
	ruleSet := RuleSet{
		HomesteadBlock: big.NewInt(1150000),
	}

	fn := filepath.Join(stateTestDir, "stCallCodes.json")
	if err := RunStateTest(ruleSet, fn, StateSkipTests); err != nil {
		t.Error(err)
	}
}

func TestDelegateCall(t *testing.T) {
	ruleSet := RuleSet{
		HomesteadBlock: big.NewInt(1150000),
	}

	fn := filepath.Join(stateTestDir, "stDelegatecallTest.json")
	if err := RunStateTest(ruleSet, fn, StateSkipTests); err != nil {
		t.Error(err)
	}
}

func TestMemory(t *testing.T) {
	ruleSet := RuleSet{
		HomesteadBlock: big.NewInt(1150000),
	}

	fn := filepath.Join(stateTestDir, "stMemoryTest.json")
	if err := RunStateTest(ruleSet, fn, StateSkipTests); err != nil {
		t.Error(err)
	}
}

func TestMemoryStress(t *testing.T) {
	ruleSet := RuleSet{
		HomesteadBlock: big.NewInt(1150000),
	}

	if os.Getenv("TEST_VM_COMPLEX") == "" {
		t.Skip()
	}
	fn := filepath.Join(stateTestDir, "stMemoryStressTest.json")
	if err := RunStateTest(ruleSet, fn, StateSkipTests); err != nil {
		t.Error(err)
	}
}

func TestQuadraticComplexity(t *testing.T) {
	ruleSet := RuleSet{
		HomesteadBlock: big.NewInt(1150000),
	}

	if os.Getenv("TEST_VM_COMPLEX") == "" {
		t.Skip()
	}
	fn := filepath.Join(stateTestDir, "stQuadraticComplexityTest.json")
	if err := RunStateTest(ruleSet, fn, StateSkipTests); err != nil {
		t.Error(err)
	}
}

func TestSolidity(t *testing.T) {
	ruleSet := RuleSet{
		HomesteadBlock: big.NewInt(1150000),
	}

	fn := filepath.Join(stateTestDir, "stSolidityTest.json")
	if err := RunStateTest(ruleSet, fn, StateSkipTests); err != nil {
		t.Error(err)
	}
}

func TestWallet(t *testing.T) {
	ruleSet := RuleSet{
		HomesteadBlock: big.NewInt(1150000),
	}

	fn := filepath.Join(stateTestDir, "stWalletTest.json")
	if err := RunStateTest(ruleSet, fn, StateSkipTests); err != nil {
		t.Error(err)
	}
}

func TestStateTestsRandom(t *testing.T) {
	ruleSet := RuleSet{
		HomesteadBlock: big.NewInt(1150000),
	}

	fns, _ := filepath.Glob("./files/StateTests/RandomTests/*")
	for _, fn := range fns {
		if err := RunStateTest(ruleSet, fn, StateSkipTests); err != nil {
			t.Error(fn, err)
=======

	"github.com/ethereum/go-ethereum/core/vm"
)

func TestState(t *testing.T) {
	t.Parallel()

	st := new(testMatcher)
	// Long tests:
	st.skipShortMode(`^stQuadraticComplexityTest/`)
	// Broken tests:
	st.skipLoad(`^stTransactionTest/OverflowGasRequire\.json`) // gasLimit > 256 bits
	st.skipLoad(`^stTransactionTest/zeroSigTransa[^/]*\.json`) // EIP-86 is not supported yet
	// Expected failures:
	st.fails(`^stRevertTest/RevertPrecompiledTouch\.json/EIP158`, "bug in test")
	st.fails(`^stRevertTest/RevertPrefoundEmptyOOG\.json/EIP158`, "bug in test")
	st.fails(`^stRevertTest/RevertPrecompiledTouch\.json/Byzantium`, "bug in test")
	st.fails(`^stRevertTest/RevertPrefoundEmptyOOG\.json/Byzantium`, "bug in test")
	st.fails(`^stRandom/randomStatetest645\.json/EIP150/.*`, "known bug #15119")
	st.fails(`^stRandom/randomStatetest645\.json/Frontier/.*`, "known bug #15119")
	st.fails(`^stRandom/randomStatetest645\.json/Homestead/.*`, "known bug #15119")
	st.fails(`^stRandom/randomStatetest644\.json/EIP150/.*`, "known bug #15119")
	st.fails(`^stRandom/randomStatetest644\.json/Frontier/.*`, "known bug #15119")
	st.fails(`^stRandom/randomStatetest644\.json/Homestead/.*`, "known bug #15119")
	st.fails(`^stCreateTest/TransactionCollisionToEmpty\.json/EIP158/2`, "known bug ")
	st.fails(`^stCreateTest/TransactionCollisionToEmpty\.json/EIP158/3`, "known bug ")
	st.fails(`^stCreateTest/TransactionCollisionToEmpty\.json/Byzantium/2`, "known bug ")
	st.fails(`^stCreateTest/TransactionCollisionToEmpty\.json/Byzantium/3`, "known bug ")
	st.walk(t, stateTestDir, func(t *testing.T, name string, test *StateTest) {
		for _, subtest := range test.Subtests() {
			subtest := subtest
			key := fmt.Sprintf("%s/%d", subtest.Fork, subtest.Index)
			name := name + "/" + key
			t.Run(key, func(t *testing.T) {
				if subtest.Fork == "Constantinople" {
					t.Skip("constantinople not supported yet")
				}
				withTrace(t, test.gasLimit(subtest), func(vmconfig vm.Config) error {
					_, err := test.Run(subtest, vmconfig)
					return st.checkFailure(t, name, err)
				})
			})
>>>>>>> 1d06e41f04d75c31334c455063e9ec7b4136bf23
		}
	})
}

// Transactions with gasLimit above this value will not get a VM trace on failure.
//const traceErrorLimit = 400000
const traceErrorLimit = 0

func withTrace(t *testing.T, gasLimit uint64, test func(vm.Config) error) {
	err := test(vm.Config{})
	if err == nil {
		return
	}
	t.Error(err)
	if gasLimit > traceErrorLimit {
		t.Log("gas limit too high for EVM trace")
		return
	}
	tracer := vm.NewStructLogger(nil)
	err2 := test(vm.Config{Debug: true, Tracer: tracer})
	if !reflect.DeepEqual(err, err2) {
		t.Errorf("different error for second run: %v", err2)
	}
	buf := new(bytes.Buffer)
	vm.WriteTrace(buf, tracer.StructLogs())
	if buf.Len() == 0 {
		t.Log("no EVM operation logs generated")
	} else {
		t.Log("EVM operation log:\n" + buf.String())
	}
}

func TestHomesteadBounds(t *testing.T) {
	ruleSet := RuleSet{
		HomesteadBlock: new(big.Int),
	}

	fn := filepath.Join(stateTestDir, "Homestead", "stBoundsTest.json")
	if err := RunStateTest(ruleSet, fn, StateSkipTests); err != nil {
		t.Error(err)
	}
}

// EIP150 tests
func TestEIP150Specific(t *testing.T) {
	ruleSet := RuleSet{
		HomesteadBlock:           new(big.Int),
		HomesteadGasRepriceBlock: big.NewInt(2457000),
	}

	fn := filepath.Join(stateTestDir, "EIP150", "stEIPSpecificTest.json")
	if err := RunStateTest(ruleSet, fn, StateSkipTests); err != nil {
		t.Error(err)
	}
}

func TestEIP150SingleCodeGasPrice(t *testing.T) {
	ruleSet := RuleSet{
		HomesteadBlock:           new(big.Int),
		HomesteadGasRepriceBlock: big.NewInt(2457000),
	}

	fn := filepath.Join(stateTestDir, "EIP150", "stEIPSingleCodeGasPrices.json")
	if err := RunStateTest(ruleSet, fn, StateSkipTests); err != nil {
		t.Error(err)
	}
}

func TestEIP150MemExpandingCalls(t *testing.T) {
	ruleSet := RuleSet{
		HomesteadBlock:           new(big.Int),
		HomesteadGasRepriceBlock: big.NewInt(2457000),
	}

	fn := filepath.Join(stateTestDir, "EIP150", "stMemExpandingEIPCalls.json")
	if err := RunStateTest(ruleSet, fn, StateSkipTests); err != nil {
		t.Error(err)
	}
}

func TestEIP150HomesteadStateSystemOperations(t *testing.T) {
	ruleSet := RuleSet{
		HomesteadBlock:           new(big.Int),
		HomesteadGasRepriceBlock: big.NewInt(2457000),
	}

	fn := filepath.Join(stateTestDir, "EIP150", "Homestead", "stSystemOperationsTest.json")
	if err := RunStateTest(ruleSet, fn, StateSkipTests); err != nil {
		t.Error(err)
	}
}

func TestEIP150HomesteadStatePreCompiledContracts(t *testing.T) {
	ruleSet := RuleSet{
		HomesteadBlock:           new(big.Int),
		HomesteadGasRepriceBlock: big.NewInt(2457000),
	}

	fn := filepath.Join(stateTestDir, "EIP150", "Homestead", "stPreCompiledContracts.json")
	if err := RunStateTest(ruleSet, fn, StateSkipTests); err != nil {
		t.Error(err)
	}
}

func TestEIP150HomesteadStateRecursiveCreate(t *testing.T) {
	ruleSet := RuleSet{
		HomesteadBlock:           new(big.Int),
		HomesteadGasRepriceBlock: big.NewInt(2457000),
	}

	fn := filepath.Join(stateTestDir, "EIP150", "Homestead", "stSpecialTest.json")
	if err := RunStateTest(ruleSet, fn, StateSkipTests); err != nil {
		t.Error(err)
	}
}

func TestEIP150HomesteadStateRefund(t *testing.T) {
	ruleSet := RuleSet{
		HomesteadBlock:           new(big.Int),
		HomesteadGasRepriceBlock: big.NewInt(2457000),
	}

	fn := filepath.Join(stateTestDir, "EIP150", "Homestead", "stRefundTest.json")
	if err := RunStateTest(ruleSet, fn, StateSkipTests); err != nil {
		t.Error(err)
	}
}

func TestEIP150HomesteadStateInitCode(t *testing.T) {
	ruleSet := RuleSet{
		HomesteadBlock:           new(big.Int),
		HomesteadGasRepriceBlock: big.NewInt(2457000),
	}

	fn := filepath.Join(stateTestDir, "EIP150", "Homestead", "stInitCodeTest.json")
	if err := RunStateTest(ruleSet, fn, StateSkipTests); err != nil {
		t.Error(err)
	}
}

func TestEIP150HomesteadStateLog(t *testing.T) {
	ruleSet := RuleSet{
		HomesteadBlock:           new(big.Int),
		HomesteadGasRepriceBlock: big.NewInt(2457000),
	}

	fn := filepath.Join(stateTestDir, "EIP150", "Homestead", "stLogTests.json")
	if err := RunStateTest(ruleSet, fn, StateSkipTests); err != nil {
		t.Error(err)
	}
}

func TestEIP150HomesteadStateTransaction(t *testing.T) {
	ruleSet := RuleSet{
		HomesteadBlock:           new(big.Int),
		HomesteadGasRepriceBlock: big.NewInt(2457000),
	}

	fn := filepath.Join(stateTestDir, "EIP150", "Homestead", "stTransactionTest.json")
	if err := RunStateTest(ruleSet, fn, StateSkipTests); err != nil {
		t.Error(err)
	}
}

func TestEIP150HomesteadCallCreateCallCode(t *testing.T) {
	ruleSet := RuleSet{
		HomesteadBlock:           new(big.Int),
		HomesteadGasRepriceBlock: big.NewInt(2457000),
	}

	fn := filepath.Join(stateTestDir, "EIP150", "Homestead", "stCallCreateCallCodeTest.json")
	if err := RunStateTest(ruleSet, fn, StateSkipTests); err != nil {
		t.Error(err)
	}
}

func TestEIP150HomesteadCallCodes(t *testing.T) {
	ruleSet := RuleSet{
		HomesteadBlock:           new(big.Int),
		HomesteadGasRepriceBlock: big.NewInt(2457000),
	}

	fn := filepath.Join(stateTestDir, "EIP150", "Homestead", "stCallCodes.json")
	if err := RunStateTest(ruleSet, fn, StateSkipTests); err != nil {
		t.Error(err)
	}
}

func TestEIP150HomesteadMemory(t *testing.T) {
	ruleSet := RuleSet{
		HomesteadBlock:           new(big.Int),
		HomesteadGasRepriceBlock: big.NewInt(2457000),
	}

	fn := filepath.Join(stateTestDir, "EIP150", "Homestead", "stMemoryTest.json")
	if err := RunStateTest(ruleSet, fn, StateSkipTests); err != nil {
		t.Error(err)
	}
}

func TestEIP150HomesteadMemoryStress(t *testing.T) {
	ruleSet := RuleSet{
		HomesteadBlock:           new(big.Int),
		HomesteadGasRepriceBlock: big.NewInt(2457000),
	}

	if os.Getenv("TEST_VM_COMPLEX") == "" {
		t.Skip()
	}
	fn := filepath.Join(stateTestDir, "EIP150", "Homestead", "stMemoryStressTest.json")
	if err := RunStateTest(ruleSet, fn, StateSkipTests); err != nil {
		t.Error(err)
	}
}

func TestEIP150HomesteadQuadraticComplexity(t *testing.T) {
	ruleSet := RuleSet{
		HomesteadBlock:           new(big.Int),
		HomesteadGasRepriceBlock: big.NewInt(2457000),
	}

	if os.Getenv("TEST_VM_COMPLEX") == "" {
		t.Skip()
	}
	fn := filepath.Join(stateTestDir, "EIP150", "Homestead", "stQuadraticComplexityTest.json")
	if err := RunStateTest(ruleSet, fn, StateSkipTests); err != nil {
		t.Error(err)
	}
}

func TestEIP150HomesteadWallet(t *testing.T) {
	ruleSet := RuleSet{
		HomesteadBlock:           new(big.Int),
		HomesteadGasRepriceBlock: big.NewInt(2457000),
	}

	fn := filepath.Join(stateTestDir, "EIP150", "Homestead", "stWalletTest.json")
	if err := RunStateTest(ruleSet, fn, StateSkipTests); err != nil {
		t.Error(err)
	}
}

func TestEIP150HomesteadDelegateCodes(t *testing.T) {
	ruleSet := RuleSet{
		HomesteadBlock:           new(big.Int),
		HomesteadGasRepriceBlock: big.NewInt(2457000),
	}

	fn := filepath.Join(stateTestDir, "EIP150", "Homestead", "stCallDelegateCodes.json")
	if err := RunStateTest(ruleSet, fn, StateSkipTests); err != nil {
		t.Error(err)
	}
}

func TestEIP150HomesteadDelegateCodesCallCode(t *testing.T) {
	ruleSet := RuleSet{
		HomesteadBlock:           new(big.Int),
		HomesteadGasRepriceBlock: big.NewInt(2457000),
	}

	fn := filepath.Join(stateTestDir, "EIP150", "Homestead", "stCallDelegateCodesCallCode.json")
	if err := RunStateTest(ruleSet, fn, StateSkipTests); err != nil {
		t.Error(err)
	}
}

func TestEIP150HomesteadBounds(t *testing.T) {
	ruleSet := RuleSet{
		HomesteadBlock:           new(big.Int),
		HomesteadGasRepriceBlock: big.NewInt(2457000),
	}

	fn := filepath.Join(stateTestDir, "EIP150", "Homestead", "stBoundsTest.json")
	if err := RunStateTest(ruleSet, fn, StateSkipTests); err != nil {
		t.Error(err)
	}
}
