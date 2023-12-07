package pendaftaran

import (
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func CreateMariaGormConnection(Mariastring string) *gorm.DB {
	DB, err := gorm.Open(mysql.Open(Mariastring), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	DB.Statement.RaiseErrorOnNotFound = true

	if err != nil {
		panic(err)
	}
	return DB
}

func CreatePostgresConnection(postgresstring string) *gorm.DB {
	DB, err := gorm.Open(postgres.Open(postgresstring), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	DB.Statement.RaiseErrorOnNotFound = true

	if err != nil {
		panic(err)
	}
	return DB
}
