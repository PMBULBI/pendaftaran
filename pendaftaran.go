package pendaftaran

import (
	"context"
	pmbulbi "github.com/PMBULBI/types/schemas"
	"github.com/golang-module/carbon/v2"
)

func Pendaftaran(ctx context.Context, Mariaenv string, val pmbulbi.Pendaftaran) (res string, err error) {
	conn := CreateMariaGormConnection(Mariaenv)
	passwordencrypted := Encrypter(val.Password)
	data := pmbulbi.Pendaftaran{
		ID:              val.ID,
		NamaMhs:         val.NamaMhs,
		AsalSekolah:     val.AsalSekolah,
		EmailMhs:        val.EmailMhs,
		HpMhs:           val.HpMhs,
		ProvinsiSekolah: val.ProvinsiSekolah,
		KotaSekolah:     val.KotaSekolah,
		Password:        passwordencrypted,
		StatusMhs:       val.StatusMhs,
		UsernameAdmin:   val.UsernameAdmin,
		TglDaftarMhs:    carbon.Now(),
	}
	err = InsertDataPendaftar(conn, ctx, data)
	if err != nil {
		return "data gagal input ", err
	}
	return "Data berhasil Dinput", nil
}
