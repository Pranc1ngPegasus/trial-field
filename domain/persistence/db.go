//go:generate go run github.com/golang/mock/mockgen -source=$GOFILE -destination=mock/$GOFILE -package=mock

package persistence

import "database/sql"

type DB interface {
	Conn() *sql.DB
}
