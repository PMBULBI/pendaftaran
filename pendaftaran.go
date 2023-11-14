package pendaftaran

import (
	"context"
	"database/sql"
	"errors"
	"github.com/PMBULBI/pendaftaran/Admin"
	pmbulbi "github.com/PMBULBI/types/schemas"
	"github.com/golang-module/carbon/v2"
)

func Pendaftaran(ctx context.Context, Mariaenv, secret string, val pmbulbi.Pendaftaran) (err error) {
	conn := CreateMariaGormConnection(Mariaenv)

	_, err = CheckUserExists(conn, ctx, val.EmailMhs, val.HpMhs)
	if err == nil {
		return errors.New("Email dan Nomor HP sudah terdaftar")
	}

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

func GetAllPendaftaran(ctx context.Context, Mariaenv string) (dataProvinsi []pmbulbi.ResponsePendaftaran, err error) {
	conn := CreateMariaGormConnection(Mariaenv)

	val, err := Admin.GetAllPendaftar(conn, ctx)

	if err != nil {
		return nil, err
	}

	for _, v := range val {
		dataProvinsi = append(dataProvinsi, pmbulbi.ResponsePendaftaran{
			ID:              v.ID,
			NamaMhs:         v.NamaMhs,
			AsalSekolah:     v.AsalSekolah,
			EmailMhs:        v.EmailMhs,
			HpMhs:           v.HpMhs,
			ProvinsiSekolah: v.ProvinsiSekolah.String,
			KotaSekolah:     v.KotaSekolah.String,
			Password:        v.Password,
			StatusMhs:       v.StatusMhs,
			UsernameAdmin:   v.UsernameAdmin.String,
			TglDaftarMhs:    carbon.DateTime{v.TglDaftarMhs},
		})
	}

	return dataProvinsi, nil
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

func GetOnePendaftar(ctx context.Context, Mariaenv, id string) (data pmbulbi.ResponsePendaftaran, err error) {
	conn := CreateMariaGormConnection(Mariaenv)

	user, err := Admin.GetSatuPendaftar(conn, ctx, id)
	if err != nil {
		return pmbulbi.ResponsePendaftaran{}, err
	}

	data.ID = user.ID
	data.NamaMhs = user.NamaMhs
	data.AsalSekolah = user.AsalSekolah
	data.EmailMhs = user.EmailMhs
	data.HpMhs = user.HpMhs
	data.ProvinsiSekolah = user.ProvinsiSekolah.String
	data.KotaSekolah = user.KotaSekolah.String
	data.Password = user.Password
	data.StatusMhs = user.StatusMhs
	data.UsernameAdmin = user.UsernameAdmin.String
	data.TglDaftarMhs = carbon.DateTime{Carbon: user.TglDaftarMhs}

	return data, nil
}

func GetDaftarSekolahlimit5(ctx context.Context, Mariaenv string, limit int, schoolname string) (data []pmbulbi.DaftarSekolah, err error) {
	conn := CreateMariaGormConnection(Mariaenv)

	school, err := GetSekolahLimits(conn, ctx, limit, schoolname)
	if err != nil {
		return nil, err
	}

	return school, nil
}

func GetProvinsiNama(ctx context.Context, Mariaenv string, provname string) (data []pmbulbi.WilayahProvinsi, err error) {
	conn := CreateMariaGormConnection(Mariaenv)

	provinsi, err := GetProvNama(conn, ctx, provname)
	if err != nil {
		return nil, err
	}

	return provinsi, nil
}

func GetKotaByProvId(ctx context.Context, Mariaenv string, id_prov string) (data []pmbulbi.WilayahKota, err error) {
	conn := CreateMariaGormConnection(Mariaenv)

	kota, err := GetKotaByIdProvinsi(conn, ctx, id_prov)
	if err != nil {
		return nil, err
	}

	return kota, nil
}

func GetKotaByProvIdKotaNama(ctx context.Context, Mariaenv string, id_prov string, nama_kota string) (data []pmbulbi.WilayahKota, err error) {
	conn := CreateMariaGormConnection(Mariaenv)

	kota, err := GetKotaByIdProvinsiNamaKota(conn, ctx, id_prov, nama_kota)
	if err != nil {
		return nil, err
	}

	return kota, nil
}

func VerifyPassword(user pmbulbi.Pendaftaran, secret string, val pmbulbi.RequestLogin) error {
	decryptedPassword := Decrypt(user.Password, secret)
	if user.EmailMhs != val.Email || decryptedPassword != val.Password {
		return errors.New("Email atau password salah")
	}

	return nil
}

func Login(ctx context.Context, Mariaenv, secret string, val pmbulbi.RequestLogin) (err error) {
	conn := CreateMariaGormConnection(Mariaenv)

	user, err := CheckUserExists(conn, ctx, val.Email, val.Password)
	if err != nil {
		return errors.New("User tidak terdaftar")
	}

	err = VerifyPassword(user, secret, val)
	if err != nil {
		return err
	}

	return nil
}
