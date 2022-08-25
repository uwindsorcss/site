package models

type event struct {
	Id                   int    `db:"id" json:"id"`
	Title                string `db:"title" json:"title"`
	Description          string `db:"description" json:"description"`
	Location             string `db:"location" json:"location"`
	Capacity             int    `db:"capacity" json:"capacity"`
	Start_Date           string `db:"start_date" json:"start_date"`
	End_Date             string `db:"end_date" json:"end_date"`
	Registration_Enabled bool   `db:"registration_enabled" json:"registration_enabled"`
	Discord_Message_Id   string `db:"discord_message_id" json:"discord_message_id"`
	Discord_Enabled      bool   `db:"discord_enabled" json:"discord_enabled"`
}
