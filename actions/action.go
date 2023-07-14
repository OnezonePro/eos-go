package actions

import (
	"github.com/armoniax/eos-go"
)

func NewStake(from, receiver, account eos.AccountName, stakeNetQuantity, stakeCpuQuantity eos.Asset, transfer bool) *eos.Action {
	return &eos.Action{
		Account: account,
		Name:    eos.ActN("delegatebw"),
		Authorization: []eos.PermissionLevel{
			{Actor: from, Permission: eos.PN("active")},
		},
		ActionData: eos.NewActionData(Stake{
			From:             from,
			Receiver:         receiver,
			StakeNetQuantity: stakeNetQuantity,
			StakeCpuQuantity: stakeCpuQuantity,
			Transfer:         transfer,
		}),
	}
}

// Stake represents the `Stake` struct on `eosio.token` contract.
type Stake struct {
	From             eos.AccountName `json:"from"`
	Receiver         eos.AccountName `json:"receiver"`
	StakeNetQuantity eos.Asset       `json:"stake_net_quantity"`
	StakeCpuQuantity eos.Asset       `json:"stake_cpu_quantity"`
	Transfer         bool            `json:"transfer"`
}

func NewBuyRam(payer, receiver, account eos.AccountName, quantity eos.Asset) *eos.Action {
	return &eos.Action{
		Account: account,
		Name:    eos.ActN("buyram"),
		Authorization: []eos.PermissionLevel{
			{Actor: payer, Permission: eos.PN("active")},
		},
		ActionData: eos.NewActionData(BuyRam{
			Payer:    payer,
			Receiver: receiver,
			Quant:    quantity,
		}),
	}
}

// BuyRam represents the `BuyRam` struct on `eosio.token` contract.
type BuyRam struct {
	Payer    eos.AccountName `json:"payer"`
	Receiver eos.AccountName `json:"receiver"`
	Quant    eos.Asset       `json:"quant"`
}

func NewSellRam(payer, account eos.AccountName, bytes int64) *eos.Action {
	return &eos.Action{
		Account: account,
		Name:    eos.ActN("sellram"),
		Authorization: []eos.PermissionLevel{
			{Actor: payer, Permission: eos.PN("active")},
		},
		ActionData: eos.NewActionData(SellRam{
			Account: payer,
			Bytes:   bytes,
		}),
	}
}

// SellRam represents the `SellRam` struct on `eosio.token` contract.
type SellRam struct {
	Account eos.AccountName `json:"account"`
	Bytes   int64           `json:"bytes"`
}

func NewUnstake(from, receiver, account eos.AccountName, unstakeNetQuantity, unstakeCpuQuantity eos.Asset) *eos.Action {
	return &eos.Action{
		Account: account,
		Name:    eos.ActN("undelegatebw"),
		Authorization: []eos.PermissionLevel{
			{Actor: from, Permission: eos.PN("active")},
		},
		ActionData: eos.NewActionData(Unstake{
			From:               from,
			Receiver:           receiver,
			UnstakeNetQuantity: unstakeNetQuantity,
			UnstakeCpuQuantity: unstakeCpuQuantity,
		}),
	}
}

// Unstake represents the `Unstake` struct on `eosio.token` contract.
type Unstake struct {
	From               eos.AccountName `json:"from"` //  可以填写抵押资源借给对方的账号（回收）
	Receiver           eos.AccountName `json:"receiver"`
	UnstakeNetQuantity eos.Asset       `json:"unstake_net_quantity"`
	UnstakeCpuQuantity eos.Asset       `json:"unstake_cpu_quantity"`
}

func NewAmaxTransfer(from, to, account eos.AccountName, quantity eos.Asset, memo string) *eos.Action {
	action := &eos.Action{
		Account: account,
		Name:    eos.ActN("transfer"),
		Authorization: []eos.PermissionLevel{
			{Actor: from, Permission: eos.PN("active")},
		},
		ActionData: eos.NewActionData(Transfer{
			From:     from,
			To:       to,
			Quantity: quantity,
			Memo:     memo,
		}),
	}

	return action
}

// Transfer represents the `transfer` struct on `eosio.token` contract.
type Transfer struct {
	From     eos.AccountName `json:"from"`
	To       eos.AccountName `json:"to"`
	Quantity eos.Asset       `json:"quantity"`
	Memo     string          `json:"memo"`
}

// NewUpdateAuth creates an action from the `eosio.system` contract
// called `updateauth`.
//
// usingPermission needs to be `owner` if you want to modify the
// `owner` authorization, otherwise `active` will do for the rest.
func NewUpdateAuth(account eos.AccountName, permission, parent eos.PermissionName, authority eos.Authority, usingPermission eos.PermissionName) *eos.Action {
	a := &eos.Action{
		Account: eos.AN("amax"),
		Name:    eos.ActN("updateauth"),
		Authorization: []eos.PermissionLevel{
			{
				Actor:      account,
				Permission: usingPermission,
			},
		},
		ActionData: eos.NewActionData(UpdateAuth{
			Account:    account,
			Permission: permission,
			Parent:     parent,
			Auth:       authority,
		}),
	}

	return a
}

// UpdateAuth represents the hard-coded `updateauth` action.
//
// If you change the `active` permission, `owner` is the required parent.
//
// If you change the `owner` permission, there should be no parent.
type UpdateAuth struct {
	Account    eos.AccountName    `json:"account"`
	Permission eos.PermissionName `json:"permission"`
	Parent     eos.PermissionName `json:"parent"`
	Auth       eos.Authority      `json:"auth"`
}
