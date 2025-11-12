package models

import "time"

type User struct {
	ID        string    `db:"id" json:"id"`
	Email     string    `db:"email" json:"email"`
	Password  string    `db:"password" json:"password"`
	Name      string    `db:"name" json:"name"`
	Username  string    `db:"username" json:"username"`
	LastSeen  time.Time `db:"last_seen" json:"last_seen"`
	AvatarURL string    `db:"avatar_url" json:"avatar_url"`
	Bio       string    `db:"bio" json:"bio"`
}
