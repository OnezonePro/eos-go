package main

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/armoniax/eos-go"
	"github.com/armoniax/eos-go/actions"
)

var (
	api           = "https://test-chain.ambt.art"
	myAccountName = "aplgxgvkzopd"
	myPrivateKey  = "5JoeXokHJXKrwPnnjxvVMdxhitBaZbNC1d7TEyc253UeVPKbDmP"
	myPublicKey   = "AM7oJhsiu5wfPVoLZ5PDVu41p4r2zPswTjdj2yGLDuPzRAap1NXV"
	myPermission  = "active"
)

func main() {
	client := eos.New(api)
	ctx := context.Background()

	keyBag := &eos.KeyBag{}
	err := keyBag.ImportAmaxPrivateKey(ctx, myPrivateKey)
	if err != nil {
		panic(fmt.Errorf("import private key: %w", err))
	}
	client.SetSigner(keyBag)

	// 给自己买
	payer := eos.AccountName(myAccountName)
	receiver := eos.AccountName(myAccountName)

	quantity, err := eos.NewAMAXAssetFromString("0.10000000 AMAX")
	if err != nil {
		panic(fmt.Errorf("invalid stakeNetQuantity: %w", err))
	}

	txOpts := &eos.TxOptions{}
	if err := txOpts.FillFromChain(context.Background(), client); err != nil {
		panic(fmt.Errorf("filling tx opts: %w", err))
	}

	tx := eos.NewTransaction([]*eos.Action{actions.NewBuyRam(payer, receiver, "amax", quantity)}, txOpts)
	signedTx, packedTx, err := client.SignTransaction(context.Background(), tx, txOpts.ChainID, eos.CompressionNone)
	if err != nil {
		panic(fmt.Errorf("sign transaction: %w", err))
	}

	content, err := json.MarshalIndent(signedTx, "", "  ")
	if err != nil {
		panic(fmt.Errorf("json marshalling transaction: %w", err))
	}

	fmt.Printf("signedTx: %s\n\n", string(content))

	response, err := client.PushTransaction(context.Background(), packedTx)
	fmt.Printf("res: %v\n", response)
	if err != nil {
		panic(fmt.Errorf("push transaction: %w", err))
	}

	fmt.Printf("Transaction [%s] submitted to the network succesfully.\n", hex.EncodeToString(response.Processed.ID))
}
