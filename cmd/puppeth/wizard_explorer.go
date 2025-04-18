// Copyright 2017 The go-ethereum Authors & The SparkAI authors
// This file is part of wodchain. Forked from the go-ethereum.
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
	"encoding/json"
	"fmt"
	"time"

	"github.com/wodTeam/Wod_Chain/log"
)

// deployExplorer creates a new block explorer based on some user input.
func (w *wizard) deployExplorer() {
	// Do some sanity check before the user wastes time on input
	if w.conf.Genesis == nil {
		log.Error("No genesis block configured")
		return
	}
	if w.conf.ethstats == "" {
		log.Error("No ethstats server configured")
		return
	}
	// Select the server to interact with
	server := w.selectServer()
	if server == "" {
		return
	}
	client := w.servers[server]

	// Retrieve any active node configurations from the server
	infos, err := checkExplorer(client, w.network)
	if err != nil {
		infos = &explorerInfos{
			node: &nodeInfos{port: 30303},
			port: 80,
			host: client.server,
		}
	}
	existed := err == nil

	infos.node.genesis, _ = json.MarshalIndent(w.conf.Genesis, "", "  ")
	infos.node.network = w.conf.Genesis.Config.ChainID.Int64()

	// Figure out which port to listen on
	fmt.Println()
	fmt.Printf("Which port should the explorer listen on? (default = %d)\n", infos.port)
	infos.port = w.readDefaultInt(infos.port)

	// Figure which virtual-host to deploy ethstats on
	if infos.host, err = w.ensureVirtualHost(client, infos.port, infos.host); err != nil {
		log.Error("Failed to decide on explorer host", "err", err)
		return
	}
	// Figure out where the user wants to store the persistent data
	fmt.Println()
	if infos.node.datadir == "" {
		fmt.Printf("Where should node data be stored on the remote machine?\n")
		infos.node.datadir = w.readString()
	} else {
		fmt.Printf("Where should node data be stored on the remote machine? (default = %s)\n", infos.node.datadir)
		infos.node.datadir = w.readDefaultString(infos.node.datadir)
	}
	// Figure out where the user wants to store the persistent data for backend database
	fmt.Println()
	if infos.dbdir == "" {
		fmt.Printf("Where should postgres data be stored on the remote machine?\n")
		infos.dbdir = w.readString()
	} else {
		fmt.Printf("Where should postgres data be stored on the remote machine? (default = %s)\n", infos.dbdir)
		infos.dbdir = w.readDefaultString(infos.dbdir)
	}
	// Figure out which port to listen on
	fmt.Println()
	fmt.Printf("Which TCP/UDP port should the archive node listen on? (default = %d)\n", infos.node.port)
	infos.node.port = w.readDefaultInt(infos.node.port)

	// Set a proper name to report on the stats page
	fmt.Println()
	if infos.node.ethstats == "" {
		fmt.Printf("What should the explorer be called on the stats page?\n")
		infos.node.ethstats = w.readString() + ":" + w.conf.ethstats
	} else {
		fmt.Printf("What should the explorer be called on the stats page? (default = %s)\n", infos.node.ethstats)
		infos.node.ethstats = w.readDefaultString(infos.node.ethstats) + ":" + w.conf.ethstats
	}
	// Try to deploy the explorer on the host
	nocache := false
	if existed {
		fmt.Println()
		fmt.Printf("Should the explorer be built from scratch (y/n)? (default = no)\n")
		nocache = w.readDefaultYesNo(false)
	}
	if out, err := deployExplorer(client, w.network, w.conf.bootnodes, infos, nocache, w.conf.Genesis.Config.Clique != nil); err != nil {
		log.Error("Failed to deploy explorer container", "err", err)
		if len(out) > 0 {
			fmt.Printf("%s\n", out)
		}
		return
	}
	// All ok, run a network scan to pick any changes up
	log.Info("Waiting for node to finish booting")
	time.Sleep(3 * time.Second)

	w.networkStats()
}
