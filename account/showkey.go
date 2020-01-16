re: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
/	Usage:  "Show private key of a keystore file",
	Action: accountShowkey,
	Flags: []cli.Flag{
		utils.PasswordFileFlag,
	},
	ArgsUsage: "<keyFile>",
	Description: `
Show private key of a keystore file

For non-interactive use the passphrase can be specified with the --password flag:

    efsn account showkey [options] <keyFile>
`,
}

func accountShowkey(ctx *cli.Context) error {
	keyfile := ctx.Args().First()
	if len(keyfile) == 0 {
		utils.Fatalf("keyfile must be given as argument")
	}
	keyJSON, err := ioutil.ReadFile(keyfile)
	if err != nil {
		utils.Fatalf("Could not read wallet file: %v", err)
	}

	// Decrypt key with passphrase.
	passphrase := getPassPhrase("", false, 0, utils.MakePasswordList(ctx))
	key, err := keystore.DecryptKey(keyJSON, passphrase)
	if err != nil {
		utils.Fatalf("Error decrypting key: %v", err)
	}
	privateKey := hex.EncodeToString(crypto.FromECDSA(key.PrivateKey))

	fmt.Printf("Address: {%s}, PrivateKey: {%s}\n", key.Address.String(), privateKey)
	return nil
}
