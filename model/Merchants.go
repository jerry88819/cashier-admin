package model

import (
	"time"

	"github.com/uptrace/bun"
)

type Merchants struct {
	bun.BaseModel 					`bun:"table:merchants"`
	ID				uint 			`bun:",pk"`
	Name			string
	Username		string
	Password		string
	CreatedAt		*time.Time
	UpdatedAt		*time.Time
	LastLoginAt		*time.Time
	PublicKey		string
	TotpSecret		string
} //  Merchants()

func GetOneUserInfo( username string ) ( *Merchants, error ) {
	var user Merchants
	err := GetDB().NewSelect().Model(&user).Where("username = ?", username ).Order("id DESC").Scan(ctx)
	return &user, err
} // GetAllUserInfo()

func ChangeUserPassword( temp *Merchants ) ( err error ) {
	_, err = 	GetDB().NewUpdate().
				Model(temp).
				Column("password").
				Where("username = ?", temp.Username ).
				Exec( ctx ) 
	return err
} // ChangeUserPassword()
