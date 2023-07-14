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
	myPermission  = "active"

	toAccountName = "aplj2nfyy5zz"
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
	client.Debug = true

	from := eos.AccountName(myAccountName)
	to := eos.AccountName(toAccountName)

	// symbol := eos.Symbol{
	// 	Precision: 8,
	// 	Symbol:    "AMAX",
	// }
	// amounts := fmt.Sprintf("%v AMAX", 1.2)
	// quantity, err := eos.NewFixedSymbolAssetFromString(symbol, amounts)

	quantity, err := eos.NewAMAXAssetFromString("0.00010000 AMAX")
	memo := "test transfer push call debug"

	fmt.Printf("quantity: %#v\n", quantity)

	if err != nil {
		panic(fmt.Errorf("invalid quantity: %w", err))
	}

	txOpts := &eos.TxOptions{}
	if err := txOpts.FillFromChain(ctx, client); err != nil {
		panic(fmt.Errorf("filling tx opts: %w", err))
	}

	// 其他action自己定义，类似 NewTransfer
	// 也可以直接使用 SignPushActions 方法，sign & push
	tx := eos.NewTransaction([]*eos.Action{actions.NewAmaxTransfer(from, to, "amax.token", quantity, memo)}, txOpts)

	signedTrx, packedTrx, err := client.SignTransaction(ctx, tx, txOpts.ChainID, eos.CompressionNone)
	if err != nil {
		panic(fmt.Errorf("sign transaction: %w", err))
	}

	fmt.Printf("sign-mesg degist: %v\n", signedTrx.Actions[0].Digest())

	content, err := json.MarshalIndent(signedTrx, "", "  ")
	if err != nil {
		panic(fmt.Errorf("json marshalling transaction: %w", err))
	}

	fmt.Printf("signedTrx: %s\n\n", string(content))
	fmt.Printf("packedTrx: %s\n\n", packedTrx.PackedTransaction.String())

	response, err := client.PushTransaction(ctx, packedTrx)
	if err != nil {
		panic(fmt.Errorf("push transaction: %w", err))
	}

	fmt.Printf("amax-push transaction, response====: %v\n", response)

	fmt.Printf("Transaction [%s] submitted to the network succesfully.\n", hex.EncodeToString(response.Processed.ID))
}
