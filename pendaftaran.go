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

func GetAllProvinsi(ctx context.Context, Mariaconn *gorm.DB) (dataProvinsi []pmbulbi.ProvinsiResponse, err error) {

	val, err := GetProvinsi(Mariaconn, ctx)

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

func GetAllKota(ctx context.Context, Mariaconn *gorm.DB) (dataKota []pmbulbi.KotaResponse, err error) {

	val, err := GetKota(Mariaconn, ctx)

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

func GetDaftarSekolahlimit5(ctx context.Context, Mariaconn *gorm.DB, limit int, schoolname string) (data []pmbulbi.DaftarSekolah, err error) {

	school, err := GetSekolahLimits(Mariaconn, ctx, limit, schoolname)
	if err != nil {
		return nil, err
	}

	return school, nil
}

func GetProvinsiNama(ctx context.Context, Mariaconn *gorm.DB, provname string) (data []pmbulbi.WilayahProvinsi, err error) {

	provinsi, err := GetProvNama(Mariaconn, ctx, provname)
	if err != nil {
		return nil, err
	}

	return provinsi, nil
}

func GetKotaByProvId(ctx context.Context, Mariaconn *gorm.DB, id_prov string) (data []pmbulbi.WilayahKota, err error) {

	kota, err := GetKotaByIdProvinsi(Mariaconn, ctx, id_prov)
	if err != nil {
		return nil, err
	}

	return kota, nil
}

func GetKotaByProvIdKotaNama(ctx context.Context, Mariaconn *gorm.DB, id_prov string, nama_kota string) (data []pmbulbi.WilayahKota, err error) {

	kota, err := GetKotaByIdProvinsiNamaKota(Mariaconn, ctx, id_prov, nama_kota)
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
