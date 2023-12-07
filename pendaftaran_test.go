package pendaftaran

import (
	"context"
	"database/sql"
	"fmt"
	pmbulbi "github.com/PMBULBI/types/schemas"
	"github.com/golang-module/carbon/v2"
	"testing"
)

func TestPmbdb_TableMigrator(t *testing.T) {
	Migrate, err := TableMigratorPostgres("pstgresstring",
		pmbulbi.Admin{},
		pmbulbi.Agama{},
		pmbulbi.AdminLevel{},
		pmbulbi.WilayahKota{},
		pmbulbi.Biaya{},
		//pmbulbi.BerkasPendaftaran{},
		pmbulbi.BiodataDataBerkas{},
		pmbulbi.BiodataDataSekolah{},
		pmbulbi.BiodataJalur{},
		pmbulbi.BiodataDataDiri{},
		pmbulbi.BiodataDataOrtu{},
		pmbulbi.LogVA{},
		pmbulbi.BiodataMaster{},
		pmbulbi.BiodataProdi{},
		pmbulbi.DaftarSekolah{},
		pmbulbi.JalurPendaftaran{},
		pmbulbi.Pekerjaan{},
		pmbulbi.ProgramStudi{},
		pmbulbi.Fakultas{},
		pmbulbi.WilayahKecamatan{},
		pmbulbi.WilayahKota{},
		pmbulbi.WilayahKelurahan{},
		pmbulbi.WilayahProvinsi{},
		pmbulbi.JenisSekolah{},
		pmbulbi.JurusanSekolah{},
		pmbulbi.TahunAkademik{},
	)
	fmt.Println(err)
	fmt.Println(Migrate)
}

func TestInsertDataPendaftar(t *testing.T) {
	conn := CreateMariaGormConnection("MARIA")
	data := pmbulbi.Pendaftaran{
		NamaMhs:     "Test",
		AsalSekolah: "Bandung High School",
		EmailMhs:    "Testing@gmail.com",
		HpMhs:       "62851666465",
		ProvinsiSekolah: sql.NullString{
			String: "Jawa Barat",
			Valid:  true,
		},
		KotaSekolah: sql.NullString{
			String: "Bandung",
			Valid:  true,
		},
		UsernameAdmin: sql.NullString{
			String: "gaben",
			Valid:  true,
		},
		TglDaftarMhs: carbon.Now(),
	}
	ins := InsertDataPendaftar(conn, context.Background(), data)
	fmt.Println(ins)
}

func TestEncrypt(t *testing.T) {
	strnger := fmt.Sprintf("%s", "rofiganteng")
	key := "rofigantengbanget"
	encrypted := Encrypt(strnger, key)
	decrypted := Decrypt(encrypted, key)

	fmt.Println("enc", encrypted)
	fmt.Println("dec", decrypted)
}

func TestPendaftaran(t *testing.T) {
	insdata := pmbulbi.Pendaftaran{
		NamaMhs:     "Test",
		Tahun:       "2024-2025",
		AsalSekolah: "Bandung High School",
		EmailMhs:    "popopo@gmail.com",
		HpMhs:       "6285156363",
		ProvinsiSekolah: sql.NullString{
			String: "Jawa Barat",
			Valid:  true,
		},
		KotaSekolah: sql.NullString{
			String: "Bandung",
			Valid:  true,
		},
		UsernameAdmin: sql.NullString{
			String: "",
			Valid:  false,
		},
		TglDaftarMhs: carbon.Now(),
	}
	koneksyen := CreateMariaGormConnection("Mariaenv")
	data, err := PendaftaranDecrypt(context.Background(), koneksyen, "secret", insdata)
	fmt.Printf("%+v\n", data)
	fmt.Println(err)
}

//func TestGetAllProvinsi(t *testing.T) {
//	data, err := GetAllProvinsi(context.Background(), "MARIA")
//
//	if err != nil {
//		t.Errorf("Error in GetAllProvinsi: %v", err)
//	}
//
//	fmt.Println(data)
//}
//
//func TestGetAllKota(t *testing.T) {
//	data, err := GetAllKota(context.Background(), "MARIA")
//
//	if err != nil {
//		t.Errorf("Error in GetAllProvinsi: %v", err)
//	}
//
//	fmt.Println(data)
//}
//
//func TestGetKotaPage(t *testing.T) {
//	conn := CreateMariaGormConnection("irc:rollyganteng@tcp(10.14.200.17:3307)/pmb_ulbi?parseTime=true")
//	data, err := GetSekolahLimits(conn, context.Background(), 5, "SMAN 3")
//	fmt.Println(err)
//	fmt.Printf("%+v\n", data)
//	fmt.Println(len(data))
//}
//
//func TestGetDaftarSekolahlimit5(t *testing.T) {
//	mariaenv := "irc:rollyganteng@tcp(10.14.200.17:3307)/pmb_ulbi?parseTime=true"
//	data, err := GetDaftarSekolahlimit5(context.Background(), mariaenv, 10, "SMAN 4")
//	fmt.Printf("%+v\n", err)
//	fmt.Printf("%+v\n", data)
//}
