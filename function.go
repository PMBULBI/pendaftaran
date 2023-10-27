package pendaftaran

import (
	"context"
	pmbulbi "github.com/PMBULBI/types/schemas"
	"github.com/golang-module/dongle"
	"gorm.io/gorm"
	"math/rand"
	"time"
)

func Encrypt(plaintext string, key string) (res string) {
	res = dongle.Encrypt.FromString(plaintext).ByRc4(key).ToHexString()
	return

}

func Decrypt(enc string, key string) (res string) {
	res = dongle.Decrypt.FromHexString(enc).ByRc4(key).ToString()
	return
}

const letterBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateRandomPassword(length int) string {
	rand.NewSource(time.Now().UnixNano())

	passwordBytes := make([]byte, length)

	for i := 0; i < length; i++ {
		passwordBytes[i] = letterBytes[rand.Intn(len(letterBytes))]
	}

	return string(passwordBytes)
}

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

func GetProvinsi(conn *gorm.DB, ctx context.Context) (val []pmbulbi.WilayahProvinsi, err error) {
	err = conn.
		WithContext(ctx).
		Find(&val).
		Error
	return
}
