package models

type Registration struct {
	Id         int  `db:"id" json:"id"`
	User_Id    int  `db:"user_id" json:"user_id"`
	Event_Id   int  `db:"event_id" json:"event_id"`
	Waitlisted bool `db:"waitlisted" json:"waitlisted"`
}
