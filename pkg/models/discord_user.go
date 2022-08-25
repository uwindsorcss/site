package models

type discord_user struct {
	Id          int  `db:"id" json:"id"`
	Discord_Uid int  `db:"discord_uid" json:"discord_uid"`
	Verified    bool `db:"verified" json:"verified"`
	User_Id     int  `db:"user_id" json:"user_id"`
}
