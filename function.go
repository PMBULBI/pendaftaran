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

func GetKotaLimits(conn *gorm.DB, ctx context.Context, page, perPage int) (val []pmbulbi.WilayahKota, err error) {
	offset := (page - 1) * perPage
	err = conn.
		WithContext(ctx).
		Limit(perPage).
		Offset(offset).
		Find(&val).
		Error
	return
}

func GetSekolahLimits(conn *gorm.DB, ctx context.Context, lim int, name string) (dest []pmbulbi.DaftarSekolah, err error) {
	err = conn.WithContext(ctx).
		Limit(lim).
		Where("nama_sekolah LIKE ?", "%"+name+"%").
		Find(&dest).
		Error
	return
}

func GetKotaByIdProvinsi(conn *gorm.DB, ctx context.Context, id_prov string) (dest []pmbulbi.WilayahKota, err error) {
	err = conn.WithContext(ctx).
		Where("id_provinsi LIKE ?", "%"+id_prov+"%").
		Find(&dest).
		Error
	return
}

func GetKotaByIdProvinsiNamaKota(conn *gorm.DB, ctx context.Context, id_prov string, kota string) (dest []pmbulbi.WilayahKota, err error) {
	err = conn.WithContext(ctx).
		Where("id_provinsi LIKE ? AND nama_kota LIKE ?", "%"+id_prov+"%", "%"+kota+"%").
		Find(&dest).
		Error
	return
}

func CheckUserExists(conn *gorm.DB, ctx context.Context, email string, nohp string) (user pmbulbi.Pendaftaran, err error) {
	err = conn.WithContext(ctx).
		Where("email_mhs = ? OR hp_mhs = ?", email, nohp).
		First(&user).
		Error
	return user, err
}
