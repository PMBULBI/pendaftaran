package pendaftaran

import (
	"context"
	"database/sql"
	pmbulbi "github.com/PMBULBI/types/schemas"
	"github.com/golang-module/carbon/v2"
)

func Pendaftaran(ctx context.Context, Mariaenv, secret string, val pmbulbi.Pendaftaran) (err error) {
	conn := CreateMariaGormConnection(Mariaenv)

	mypass := GenerateRandomPassword(10)
	passwordencrypted := Encrypt(mypass, secret)

	data := pmbulbi.Pendaftaran{
		NamaMhs:     val.NamaMhs,
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
	err = InsertDataPendaftar(conn, ctx, data)
	if err != nil {
		return err
	}
	return
}
