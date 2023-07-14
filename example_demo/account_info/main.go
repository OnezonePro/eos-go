package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/armoniax/eos-go"
)

const (
	api         = "https://test-chain.ambt.art"
	publicKey   = "AM7oJhsiu5wfPVoLZ5PDVu41p4r2zPswTjdj2yGLDuPzRAap1NXV"
	accountName = "aplgxgvkzopd"
)

func main() {
	client := eos.New(api)
	ctx := context.Background()

	// Get AMAX Node Info
	info, err := client.GetInfo(ctx)
	if err != nil {
		panic(fmt.Errorf("get info: %s", err.Error()))
	}

	bytes, err := json.Marshal(info)
	if err != nil {
		panic(fmt.Errorf("json marshal response: %s", err.Error()))
	}
	fmt.Printf("node info: %v\n", string(bytes))

	// Get Account Info
	account, err := client.GetAccount(ctx, accountName)
	if err != nil {
		panic(fmt.Errorf("GetAccount: %s", err.Error()))
	}

	fmt.Printf("publicKey: %s\n\n", account.Permissions[0].RequiredAuth.Keys[0].PublicKey.String())

	fmt.Printf("可用余额: %d, symbol: %s, precision: %v\n", account.CoreLiquidBalance.Amount, account.CoreLiquidBalance.Symbol.Symbol, account.CoreLiquidBalance.Symbol.Precision)

	fmt.Printf("存储空间总量: %v KB, 已使用: %v KB\n", account.RAMQuota/1024, account.RAMUsage/1024)

	fmt.Printf("cpu 总量: %d µs, 可用: %d µs, 已使用: %d µs\n", account.CPULimit.Max, account.CPULimit.Available, account.CPULimit.Used)

	fmt.Printf("网络总量: %v KB, 可用: %v KB, 已使用: %v KB\n", account.NetLimit.Max/1024, account.NetLimit.Available/1024, account.NetLimit.Used/1024)

	fmt.Printf("cpu 总质押: %d, symbol: %s, 其中自己质押: %d, 别人质押: %d\n", account.TotalResources.CPUWeight.Amount,
		account.TotalResources.CPUWeight.Symbol.Symbol, account.SelfDelegatedBandwidth.CPUWeight.Amount, account.TotalResources.CPUWeight.Amount-account.SelfDelegatedBandwidth.CPUWeight.Amount)

	fmt.Printf("网络总质押: %d AMAX, 其中自己质押: %d AMAX, 别人质押: %d AMAX\n", account.TotalResources.NetWeight.Amount,
		account.SelfDelegatedBandwidth.NetWeight.Amount, account.TotalResources.NetWeight.Amount-account.SelfDelegatedBandwidth.NetWeight.Amount)

	if account.RefundRequest != nil {
		fmt.Printf("赎回中: %d AMAX\n", account.RefundRequest.CPUAmount.Amount+account.RefundRequest.NetAmount.Amount)
	}

	var cpuPrice = account.TotalResources.CPUWeight.Amount / (account.CPULimit.Max / 1000) / 3
	fmt.Printf("CPU 价格: %v AMAX/ms/Day\n", cpuPrice)

	var netPrice = account.TotalResources.NetWeight.Amount / (account.NetLimit.Max / 1024) / 3
	fmt.Printf("Net 价格: %v AMAX/KiB/Day\n", netPrice)

	tokenBalance, err := client.GetCurrencyBalance(ctx, accountName, "", "amax.token")
	if err != nil {
		panic(fmt.Errorf("get currency balance: %s", err.Error()))
	}
	fmt.Printf("token balance: %v\n", tokenBalance)
}
