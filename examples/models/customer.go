package models

import (
	"database/sql"
	"time"
)

type Customer struct {
	Paying        bool
	TaxIdentifier string
	CountryCode   string
	PayingDate    time.Time
	FileId        sql.NullInt64
	FileName      string
}
