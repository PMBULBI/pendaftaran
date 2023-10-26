package pendaftaran

import (
	"context"
	"encoding/base64"
	pmbulbi "github.com/PMBULBI/types/schemas"
	"gorm.io/gorm"
)

func TableMigrator(maria string, model interface{}) (res string, err error) {
	err = CreateMariaGormConnection(maria).AutoMigrate(&model)
	if err != nil {
		return "Gagal Migrate", err
	}
	return
}

func InsertDataPendaftar(conn *gorm.DB, ctx context.Context, val pmbulbi.Pendaftaran) (err error) {
	err = conn.
		WithContext(ctx).
		Create(&val).
		Error
	return
}

func Encrypter(password string) string {
	encrypted := base64.StdEncoding.EncodeToString([]byte(password))
	return encrypted
}

func Decrypter(passenc string) string {
	decrypted, err := base64.StdEncoding.DecodeString(passenc)
	if err != nil {
		return "Decrypt gagal"
	}
	return string(decrypted)
}
