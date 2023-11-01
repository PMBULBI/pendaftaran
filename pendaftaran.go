package pendaftaran

import (
	"context"
	"database/sql"
	"errors"
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

func GetOnePendaftar(ctx context.Context, Mariaenv, id string) (data pmbulbi.Pendaftaran, err error) {
	conn := CreateMariaGormConnection(Mariaenv)

	use, err := GetSatuPendaftar(conn, ctx, id)
	if err != nil {
		return pmbulbi.Pendaftaran{}, err
	}

	return use, nil
}

func GetDaftarSekolahlimit5(ctx context.Context, Mariaenv string, limit int, schoolname string) (data []pmbulbi.DaftarSekolah, err error) {
	conn := CreateMariaGormConnection(Mariaenv)

	school, err := GetSekolahLimits(conn, ctx, limit, schoolname)
	if err != nil {
		return nil, err
	}

	return school, nil
}

func VerifyPassword(user pmbulbi.Pendaftaran, secret string, val pmbulbi.RequestLogin) error {
	decryptedPassword := Decrypt(user.Password, secret)
	if user.EmailMhs != val.Email || decryptedPassword != val.Password {
		return errors.New("Email atau password salah")
	}

	return nil
}

func Login(ctx context.Context, Mariaenv, email, password string, secret string) (err error) {
	conn := CreateMariaGormConnection(Mariaenv)

	user, err := CheckUserExists(conn, ctx, email)
	if err != nil {
		return errors.New("User tidak terdaftar")
	}

	err = VerifyPassword(user, email, password, secret)
	if err != nil {
		return err
	}

	return nil
}
