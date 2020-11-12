package database

import (
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/pkg/errors"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	gormV2 "gorm.io/gorm"
)

// NewConn returns database conn.
func NewConn(dns string) (*gorm.DB, error) {
	dialect := "mysql"
	if strings.HasSuffix(dns, ".db") {
		dialect = "sqlite3"
	}

	conn, err := gorm.Open(dialect, dns)
	if err != nil {
		return nil, errors.Wrap(err, "database: open connection")
	}

	return conn, nil
}

// NewConnV2 returns database conn by gormV2.
func NewConnV2(dns string, config *gormV2.Config) (*gormV2.DB, error) {
	dialect := mysql.Open(dns)
	if strings.HasSuffix(dns, ".db") {
		dialect = sqlite.Open(dns)
	}

	conn, err := gormV2.Open(dialect, config)
	if err != nil {
		return nil, errors.Wrap(err, "database: open connection")
	}

	return conn, nil
}
