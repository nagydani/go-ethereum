// Copyright 2016 The go-ethereum Authors
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

package params

<<<<<<< HEAD
import "math/big"

type GasTable struct {
	ExtcodeSize *big.Int
	ExtcodeCopy *big.Int
	Balance     *big.Int
	SLoad       *big.Int
	Calls       *big.Int
	Suicide     *big.Int
=======
type GasTable struct {
	ExtcodeSize uint64
	ExtcodeCopy uint64
	Balance     uint64
	SLoad       uint64
	Calls       uint64
	Suicide     uint64

	ExpByte uint64
>>>>>>> 1d06e41f04d75c31334c455063e9ec7b4136bf23

	// CreateBySuicide occurs when the
	// refunded account is one that does
	// not exist. This logic is similar
	// to call. May be left nil. Nil means
	// not charged.
<<<<<<< HEAD
	CreateBySuicide *big.Int
=======
	CreateBySuicide uint64
>>>>>>> 1d06e41f04d75c31334c455063e9ec7b4136bf23
}

var (
	// GasTableHomestead contain the gas prices for
	// the homestead phase.
	GasTableHomestead = GasTable{
<<<<<<< HEAD
		ExtcodeSize: big.NewInt(20),
		ExtcodeCopy: big.NewInt(20),
		Balance:     big.NewInt(20),
		SLoad:       big.NewInt(50),
		Calls:       big.NewInt(40),
		Suicide:     big.NewInt(0),

		// explicitly set to nil to indicate
		// this rule does not apply to homestead.
		CreateBySuicide: nil,
=======
		ExtcodeSize: 20,
		ExtcodeCopy: 20,
		Balance:     20,
		SLoad:       50,
		Calls:       40,
		Suicide:     0,
		ExpByte:     10,
>>>>>>> 1d06e41f04d75c31334c455063e9ec7b4136bf23
	}

	// GasTableHomestead contain the gas re-prices for
	// the homestead phase.
<<<<<<< HEAD
	GasTableHomesteadGasRepriceFork = GasTable{
		ExtcodeSize: big.NewInt(700),
		ExtcodeCopy: big.NewInt(700),
		Balance:     big.NewInt(400),
		SLoad:       big.NewInt(200),
		Calls:       big.NewInt(700),
		Suicide:     big.NewInt(5000),

		CreateBySuicide: big.NewInt(25000),
=======
	GasTableEIP150 = GasTable{
		ExtcodeSize: 700,
		ExtcodeCopy: 700,
		Balance:     400,
		SLoad:       200,
		Calls:       700,
		Suicide:     5000,
		ExpByte:     10,

		CreateBySuicide: 25000,
	}

	GasTableEIP158 = GasTable{
		ExtcodeSize: 700,
		ExtcodeCopy: 700,
		Balance:     400,
		SLoad:       200,
		Calls:       700,
		Suicide:     5000,
		ExpByte:     50,

		CreateBySuicide: 25000,
>>>>>>> 1d06e41f04d75c31334c455063e9ec7b4136bf23
	}
)
