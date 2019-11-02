// Copyright 2019 The fsn-go-sdk Authors
// This file is part of the fsn-go-sdk library.
//
// The fsn-go-sdk library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The fsn-go-sdk library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the fsn-go-sdk library. If not, see <http://www.gnu.org/licenses/>.

package offline

import (
	"github.com/FusionFoundation/fsn-go-sdk/efsn/cmd/utils"
	"github.com/FusionFoundation/fsn-go-sdk/efsn/common"
	clicommon "github.com/FusionFoundation/fsn-go-sdk/fsn-cli/common"
	"gopkg.in/urfave/cli.v1"
)

var CommandSendAsset = cli.Command{
	Name:      "sendasset",
	Usage:     "(offline) build send asset raw transaction",
	ArgsUsage: "<assetID> <to> <value>",
	Description: `
build send asset raw transaction`,
	Flags:  append([]cli.Flag{}, commonFlags...),
	Action: sendasset,
}

func sendasset(ctx *cli.Context) error {
	if len(ctx.Args()) != 3 {
		cli.ShowCommandHelpAndExit(ctx, "sendasset", 1)
	}

	assetID_ := ctx.Args().Get(0)
	to_ := ctx.Args().Get(1)
	value_ := ctx.Args().Get(2)

	assetID := clicommon.GetHashFromText("assetID", assetID_)
	to := clicommon.GetAddressFromText("to", to_)
	value := clicommon.GetBigIntFromText("asset", value_)
	commonOptions := getCommonOptions(ctx)

	param := common.SendAssetParam{
		AssetID: assetID,
		To:      to,
		Value:   value,
	}
	input, err := genTxInput(common.SendAssetFunc, &param)
	if err != nil {
		utils.Fatalf("generate tx input error: %v", err)
	}

	tx, err := toFSNTx(input, commonOptions)
	if err != nil {
		utils.Fatalf("create tx error: %v", err)
	}

	return printTx(tx, false)
}