package domain

import "time"

type Paragraph struct {
	ID        int       `db:"id"`
	UUID      string    `db:"uuid"`
	Content   string    `db:"content"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
