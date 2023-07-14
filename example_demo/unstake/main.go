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

	// 可以填写抵押资源借给对方的账号（回收）
	from := eos.AccountName(myAccountName)
	receiver := eos.AccountName(myAccountName)

	unStakeNetQuantity, err := eos.NewAMAXAssetFromString("0.01000000")
	if err != nil {
		panic(fmt.Errorf("invalid unStakeNetQuantity: %w", err))
	}
	unStakeCpuQuantity, err := eos.NewAMAXAssetFromString("0.01000000")
	if err != nil {
		panic(fmt.Errorf("invalid unStakeCpuQuantity: %w", err))
	}

	txOpts := &eos.TxOptions{}
	if err := txOpts.FillFromChain(context.Background(), client); err != nil {
		panic(fmt.Errorf("filling tx opts: %w", err))
	}

	tx := eos.NewTransaction([]*eos.Action{actions.NewUnstake(from, receiver, "amax", unStakeNetQuantity, unStakeCpuQuantity)}, txOpts)
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
		// Internal Service Error: eosio_assert_message assertion failure: assertion failure with message:
		// cannot undelegate bandwidth until the chain is activated (at least 15% of all tokens participate in voting): pending console output:
		panic(fmt.Errorf("push transaction: %w", err))
	}

	fmt.Printf("Transaction [%s] submitted to the network succesfully.\n", hex.EncodeToString(response.Processed.ID))
}
