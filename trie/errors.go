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

package trie

import (
	"fmt"

	"github.com/wodTeam/Wod_Chain/common"
)

// MissingNodeError is returned by the trie functions (TryGet, TryUpdate, TryDelete)
// in the case where a trie node is not present in the local database. It contains
// information necessary for retrieving the missing node.
type MissingNodeError struct {
	Owner    common.Hash // owner of the trie if it's 2-layered trie
	NodeHash common.Hash // hash of the missing node
	Path     []byte      // hex-encoded path to the missing node
	err      error       // concrete error for missing trie node
}

// Unwrap returns the concrete error for missing trie node which
// allows us for further analysis outside.
func (err *MissingNodeError) Unwrap() error {
	return err.err
}

func (err *MissingNodeError) Error() string {
	if err.Owner == (common.Hash{}) {
		return fmt.Sprintf("missing trie node %x (path %x) %v", err.NodeHash, err.Path, err.err)
	}
	return fmt.Sprintf("missing trie node %x (owner %x) (path %x) %v", err.NodeHash, err.Owner, err.Path, err.err)
}
