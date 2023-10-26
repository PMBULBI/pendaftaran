package pendaftaran

import (
	"context"
	pmbulbi "github.com/PMBULBI/types/schemas"
	"github.com/golang-module/carbon/v2"
)

func Pendaftaran(ctx context.Context, Mariaenv string, val pmbulbi.Pendaftaran) (err error) {
	conn := CreateMariaGormConnection(Mariaenv)

	mypass := GenerateRandomPassword(10)
	passwordencrypted := Encrypter(mypass)
	data := pmbulbi.Pendaftaran{
		NamaMhs:         val.NamaMhs,
		AsalSekolah:     val.AsalSekolah,
		EmailMhs:        val.EmailMhs,
		HpMhs:           val.HpMhs,
		ProvinsiSekolah: val.ProvinsiSekolah,
		KotaSekolah:     val.KotaSekolah,
		Password:        passwordencrypted,
		UsernameAdmin:   val.UsernameAdmin,
		TglDaftarMhs:    carbon.Now(),
	}
	err = InsertDataPendaftar(conn, ctx, data)
	if err != nil {
		return err
	}
	return
}
