// Copyright 2022 The go-ethereum Authors & The SparkAI authors
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

package flags

import "github.com/urfave/cli/v2"

const (
	EthCategory        = "ETHEREUM"
	LightCategory      = "LIGHT CLIENT"
	DevCategory        = "DEVELOPER CHAIN"
	EthashCategory     = "ETHASH"
	TxPoolCategory     = "TRANSACTION POOL"
	PerfCategory       = "PERFORMANCE TUNING"
	AccountCategory    = "ACCOUNT"
	APICategory        = "API AND CONSOLE"
	NetworkingCategory = "NETWORKING"
	MinerCategory      = "MINER"
	GasPriceCategory   = "GAS PRICE ORACLE"
	VMCategory         = "VIRTUAL MACHINE"
	LoggingCategory    = "LOGGING AND DEBUGGING"
	MetricsCategory    = "METRICS AND STATS"
	MiscCategory       = "MISC"
	DeprecatedCategory = "ALIASED (deprecated)"
)

func init() {
	cli.HelpFlag.(*cli.BoolFlag).Category = MiscCategory
	cli.VersionFlag.(*cli.BoolFlag).Category = MiscCategory
}
