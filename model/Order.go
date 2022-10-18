package model

import (
	"time"

	"github.com/uptrace/bun"
)

type Order struct {
	bun.BaseModel	 					`bun:"table:orders"`
	ID					uint 			`bun:",pk"`
	MchOrderNo			string
	CashierOrderNo		string
	GwOrderNo			string
	MchId				int
	ChannelId			int
	Status				string	
	Name				string
	Amount				int
	NotifyTo			string
	ReturnTo			string
	ClientIp			string
    CreatedAt			time.Time
	UpdatedAt			*time.Time
	PaidAt				*time.Time
	NotifyAt			*time.Time
} // Order()

func GetUserOrder( id int ) ( []Order, error ) {
	var user []Order
	err := GetDB().NewSelect().Model(&user).Where("mch_id = ?", id ).Order("created_at DESC").Scan(ctx)
	return user, err
} // GetAllUserInfo()

