// Copyright 2015 The go-ethereum Authors & The SparkAI authors
//Copyright 2023 The SparkAI authors
// This file is part of The SparkAI library. Forked from the  go-ethereum project
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
	"bufio"
	"bytes"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/wodTeam/Wod_Chain/core"
	"github.com/wodTeam/Wod_Chain/core/rawdb"
	"github.com/wodTeam/Wod_Chain/core/types"
	"github.com/wodTeam/Wod_Chain/core/vm"
	"github.com/wodTeam/Wod_Chain/wod/tracers/logger"
)

func TestState(t *testing.T) {
	t.Parallel()

	st := new(testMatcher)
	// Long tests:
	st.slow(`^stAttackTest/ContractCreationSpam`)
	st.slow(`^stBadOpcode/badOpcodes`)
	st.slow(`^stPreCompiledContracts/modexp`)
	st.slow(`^stQuadraticComplexityTest/`)
	st.slow(`^stStaticCall/static_Call50000`)
	st.slow(`^stStaticCall/static_Return50000`)
	st.slow(`^stSystemOperationsTest/CallRecursiveBomb`)
	st.slow(`^stTransactionTest/Opcodes_TransactionInit`)

	// Very time consuming
	st.skipLoad(`^stTimeConsuming/`)
	st.skipLoad(`.*vmPerformance/loop.*`)

	// Uses 1GB RAM per tested fork
	st.skipLoad(`^stStaticCall/static_Call1MB`)

	// Broken tests:
	// Expected failures:
	//st.fails(`^stRevertTest/RevertPrecompiledTouch(_storage)?\.json/Byzantium/0`, "bug in test")
	//st.fails(`^stRevertTest/RevertPrecompiledTouch(_storage)?\.json/Byzantium/3`, "bug in test")
	//st.fails(`^stRevertTest/RevertPrecompiledTouch(_storage)?\.json/Constantinople/0`, "bug in test")
	//st.fails(`^stRevertTest/RevertPrecompiledTouch(_storage)?\.json/Constantinople/3`, "bug in test")
	//st.fails(`^stRevertTest/RevertPrecompiledTouch(_storage)?\.json/ConstantinopleFix/0`, "bug in test")
	//st.fails(`^stRevertTest/RevertPrecompiledTouch(_storage)?\.json/ConstantinopleFix/3`, "bug in test")

	// For Istanbul, older tests were moved into LegacyTests
	for _, dir := range []string{
		stateTestDir,
		legacyStateTestDir,
		benchmarksDir,
	} {
		st.walk(t, dir, func(t *testing.T, name string, test *StateTest) {
			for _, subtest := range test.Subtests() {
				subtest := subtest
				key := fmt.Sprintf("%s/%d", subtest.Fork, subtest.Index)

				t.Run(key+"/trie", func(t *testing.T) {
					withTrace(t, test.gasLimit(subtest), func(vmconfig vm.Config) error {
						_, _, err := test.Run(subtest, vmconfig, false)
						if err != nil && len(test.json.Post[subtest.Fork][subtest.Index].ExpectException) > 0 {
							// Ignore expected errors (TODO MariusVanDerWijden check error string)
							return nil
						}
						return st.checkFailure(t, err)
					})
				})
				t.Run(key+"/snap", func(t *testing.T) {
					withTrace(t, test.gasLimit(subtest), func(vmconfig vm.Config) error {
						snaps, statedb, err := test.Run(subtest, vmconfig, true)
						if snaps != nil && statedb != nil {
							if _, err := snaps.Journal(statedb.IntermediateRoot(false)); err != nil {
								return err
							}
						}
						if err != nil && len(test.json.Post[subtest.Fork][subtest.Index].ExpectException) > 0 {
							// Ignore expected errors (TODO MariusVanDerWijden check error string)
							return nil
						}
						return st.checkFailure(t, err)
					})
				})
			}
		})
	}
}

// Transactions with gasLimit above this value will not get a VM trace on failure.
const traceErrorLimit = 400000

func withTrace(t *testing.T, gasLimit uint64, test func(vm.Config) error) {
	// Use config from command line arguments.
	config := vm.Config{}
	err := test(config)
	if err == nil {
		return
	}

	// Test failed, re-run with tracing enabled.
	t.Error(err)
	if gasLimit > traceErrorLimit {
		t.Log("gas limit too high for EVM trace")
		return
	}
	buf := new(bytes.Buffer)
	w := bufio.NewWriter(buf)
	tracer := logger.NewJSONLogger(&logger.Config{}, w)
	config.Debug, config.Tracer = true, tracer
	err2 := test(config)
	if !reflect.DeepEqual(err, err2) {
		t.Errorf("different error for second run: %v", err2)
	}
	w.Flush()
	if buf.Len() == 0 {
		t.Log("no EVM operation logs generated")
	} else {
		t.Log("EVM operation log:\n" + buf.String())
	}
	// t.Logf("EVM output: 0x%x", tracer.Output())
	// t.Logf("EVM error: %v", tracer.Error())
}

func BenchmarkEVM(b *testing.B) {
	// Walk the directory.
	dir := benchmarksDir
	dirinfo, err := os.Stat(dir)
	if os.IsNotExist(err) || !dirinfo.IsDir() {
		fmt.Fprintf(os.Stderr, "can't find test files in %s, did you clone the evm-benchmarks submodule?\n", dir)
		b.Skip("missing test files")
	}
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if ext := filepath.Ext(path); ext == ".json" {
			name := filepath.ToSlash(strings.TrimPrefix(strings.TrimSuffix(path, ext), dir+string(filepath.Separator)))
			b.Run(name, func(b *testing.B) { runBenchmarkFile(b, path) })
		}
		return nil
	})
	if err != nil {
		b.Fatal(err)
	}
}

func runBenchmarkFile(b *testing.B, path string) {
	m := make(map[string]StateTest)
	if err := readJSONFile(path, &m); err != nil {
		b.Fatal(err)
		return
	}
	if len(m) != 1 {
		b.Fatal("expected single benchmark in a file")
		return
	}
	for _, t := range m {
		t := t
		runBenchmark(b, &t)
	}
}

func runBenchmark(b *testing.B, t *StateTest) {
	for _, subtest := range t.Subtests() {
		subtest := subtest
		key := fmt.Sprintf("%s/%d", subtest.Fork, subtest.Index)

		b.Run(key, func(b *testing.B) {
			vmconfig := vm.Config{}

			config, eips, err := GetChainConfig(subtest.Fork)
			if err != nil {
				b.Error(err)
				return
			}
			vmconfig.ExtraEips = eips
			block := t.genesis(config).ToBlock()
			_, statedb := MakePreState(rawdb.NewMemoryDatabase(), t.json.Pre, false)

			var baseFee *big.Int
			if config.IsLondon(new(big.Int)) {
				baseFee = t.json.Env.BaseFee
				if baseFee == nil {
					// Retesteth uses `0x10` for genesis baseFee. Therefore, it defaults to
					// parent - 2 : 0xa as the basefee for 'this' context.
					baseFee = big.NewInt(0x0a)
				}
			}
			post := t.json.Post[subtest.Fork][subtest.Index]
			msg, err := t.json.Tx.toMessage(post, baseFee)
			if err != nil {
				b.Error(err)
				return
			}

			// Try to recover tx with current signer
			if len(post.TxBytes) != 0 {
				var ttx types.Transaction
				err := ttx.UnmarshalBinary(post.TxBytes)
				if err != nil {
					b.Error(err)
					return
				}

				if _, err := types.Sender(types.LatestSigner(config), &ttx); err != nil {
					b.Error(err)
					return
				}
			}

			// Prepare the EVM.
			txContext := core.NewEVMTxContext(msg)
			context := core.NewEVMBlockContext(block.Header(), nil, &t.json.Env.Coinbase)
			context.GetHash = vmTestBlockHash
			context.BaseFee = baseFee
			evm := vm.NewEVM(context, txContext, statedb, config, vmconfig)

			// Create "contract" for sender to cache code analysis.
			sender := vm.NewContract(vm.AccountRef(msg.From()), vm.AccountRef(msg.From()),
				nil, 0)

			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				// Execute the message.
				snapshot := statedb.Snapshot()
				_, _, err = evm.Call(sender, *msg.To(), msg.Data(), msg.Gas(), msg.Value())
				if err != nil {
					b.Error(err)
					return
				}
				statedb.RevertToSnapshot(snapshot)
			}
		})
	}
}
