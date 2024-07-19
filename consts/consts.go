// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package consts

import (
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp"
	"github.com/ava-labs/hypersdk/chain"
	"github.com/ava-labs/hypersdk/codec"
	"github.com/ava-labs/hypersdk/consts"
)

const (
	// Choose a human-readable part for your hyperchain
	HRP = "example"

	// Choose a name for your hyperchain
	Name = "examplechain"

	// Choose a token symbol
	Symbol = "EXC"
)

var ID ids.ID

func init() {
	b := make([]byte, consts.IDLen)
	copy(b, []byte(Name))
	vmID, err := ids.ToID(b)
	if err != nil {
		panic(err)
	}
	ID = vmID
}

// Instantiate registry here so it can be imported by any package. We set these
// values in [controller/registry].
var (
	ActionRegistry *codec.TypeParser[chain.Action, *warp.Message, bool] = codec.NewTypeParser[chain.Action, *warp.Message, bool]()

	AuthRegistry *codec.TypeParser[chain.Auth, *warp.Message, bool] = codec.NewTypeParser[chain.Auth, *warp.Message, bool]()
)
