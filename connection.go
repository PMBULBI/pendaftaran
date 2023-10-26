package pendaftaran

import (
	"database/sql"
	"github.com/aiteung/atdb"
	"os"
)

func CreateMariaConnection(MariaString, dbname string) *sql.DB {
	MariaInfo := atdb.DBInfo{
		DBString: os.Getenv(MariaString),
		DBName:   dbname,
	}
	db := atdb.MariaConnect(MariaInfo)
	return db
}
