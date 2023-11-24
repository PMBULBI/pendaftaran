package pendaftaran

import (
	"context"
	"database/sql"
	"errors"
	pmbulbi "github.com/PMBULBI/types/schemas"
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

func Pendaftaran(ctx context.Context, Mariaconn *gorm.DB, secret string, val pmbulbi.Pendaftaran) (err error) {

	_, err = CheckUserExists(Mariaconn, ctx, val.EmailMhs, val.HpMhs)
	if err == nil {
		return errors.New("Email dan Nomor HP sudah terdaftar")
	}

	mypass := GenerateRandomPassword(10)
	passwordencrypted := Encrypt(mypass, secret)

	data := pmbulbi.Pendaftaran{
		NamaMhs:     val.NamaMhs,
		Tahun:       val.Tahun,
		AsalSekolah: val.AsalSekolah,
		EmailMhs:    val.EmailMhs,
		HpMhs:       val.HpMhs,
		ProvinsiSekolah: sql.NullString{
			String: val.ProvinsiSekolah.String,
			Valid:  val.ProvinsiSekolah.String != ""},
		KotaSekolah: sql.NullString{
			String: val.KotaSekolah.String,
			Valid:  val.KotaSekolah.String != ""},
		Password: passwordencrypted,
		UsernameAdmin: sql.NullString{
			String: val.UsernameAdmin.String,
			Valid:  val.UsernameAdmin.String != ""},
		TglDaftarMhs: carbon.Now(),
	}
	err = InsertDataPendaftar(Mariaconn, ctx, data)
	if err != nil {
		return err
	}
	return
}

func PendaftaranDecrypt(ctx context.Context, Mariaconn *gorm.DB, secret string, val pmbulbi.Pendaftaran) (dest pmbulbi.Pendaftaran, err error) {

	_, err = CheckUserExists(Mariaconn, ctx, val.EmailMhs, val.HpMhs)
	if err == nil {
		return pmbulbi.Pendaftaran{}, errors.New("Email dan Nomor HP sudah terdaftar")
	}

	mypass := GenerateRandomPassword(10)
	passwordencrypted := Encrypt(mypass, secret)

	data := pmbulbi.Pendaftaran{
		NamaMhs:     val.NamaMhs,
		Tahun:       val.Tahun,
		AsalSekolah: val.AsalSekolah,
		EmailMhs:    val.EmailMhs,
		HpMhs:       val.HpMhs,
		ProvinsiSekolah: sql.NullString{
			String: val.ProvinsiSekolah.String,
			Valid:  val.ProvinsiSekolah.String != ""},
		KotaSekolah: sql.NullString{
			String: val.KotaSekolah.String,
			Valid:  val.KotaSekolah.String != ""},
		Password: passwordencrypted,
		UsernameAdmin: sql.NullString{
			String: val.UsernameAdmin.String,
			Valid:  val.UsernameAdmin.String != ""},
		TglDaftarMhs: carbon.Now(),
	}
	err = InsertDataPendaftar(Mariaconn, ctx, data)
	dest = pmbulbi.Pendaftaran{
		NamaMhs:     val.NamaMhs,
		Tahun:       val.Tahun,
		AsalSekolah: val.AsalSekolah,
		EmailMhs:    val.EmailMhs,
		HpMhs:       val.HpMhs,
		ProvinsiSekolah: sql.NullString{
			String: val.ProvinsiSekolah.String,
			Valid:  val.ProvinsiSekolah.String != ""},
		KotaSekolah: sql.NullString{
			String: val.KotaSekolah.String,
			Valid:  val.KotaSekolah.String != ""},
		Password: mypass,
		UsernameAdmin: sql.NullString{
			String: val.UsernameAdmin.String,
			Valid:  val.UsernameAdmin.String != ""},
		TglDaftarMhs: carbon.Now(),
	}
	if err != nil {
		return pmbulbi.Pendaftaran{}, err
	}
	return
}

func VerifyPassword(user pmbulbi.Pendaftaran, secret string, val pmbulbi.RequestLogin) error {
	decryptedPassword := Decrypt(user.Password, secret)
	if user.EmailMhs != val.Email || decryptedPassword != val.Password {
		return errors.New("Email atau password salah")
	}

	return nil
}

func Login(ctx context.Context, Mariaconn *gorm.DB, secret string, val pmbulbi.RequestLogin) (err error) {

	user, err := CheckUserExists(Mariaconn, ctx, val.Email, val.Password)
	if err != nil {
		return errors.New("User tidak terdaftar")
	}

	err = VerifyPassword(user, secret, val)
	if err != nil {
		return err
	}

	return nil
}
