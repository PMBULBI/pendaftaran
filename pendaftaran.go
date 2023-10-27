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

func GetAllProvinsi(ctx context.Context, Mariaenv string) (dataProvinsi []pmbulbi.ProvinsiResponse, err error) {
	conn := CreateMariaGormConnection(Mariaenv)

	val, err := GetProvinsi(conn, ctx)

	if err != nil {
		return nil, err
	}

	for _, v := range val {
		dataProvinsi = append(dataProvinsi, pmbulbi.ProvinsiResponse{
			IDProvinsi:   v.IDProvinsi,
			NamaProvinsi: v.NamaProvinsi,
		})
	}

	return dataProvinsi, nil
}

func GetAllKota(ctx context.Context, Mariaenv string) (dataKota []pmbulbi.KotaResponse, err error) {
	conn := CreateMariaGormConnection(Mariaenv)

	val, err := GetKota(conn, ctx)

	if err != nil {
		return nil, err
	}

	for _, v := range val {
		dataKota = append(dataKota, pmbulbi.KotaResponse{
			IDKota:     v.IDKota,
			IDProvinsi: v.IDProvinsi,
			NamaKota:   v.NamaKota,
		})
	}

	return dataKota, nil
}
