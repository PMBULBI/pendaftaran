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
	Migrate, err := TableMigrator("MARIA`", pmbulbi.WilayahKota{})
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
			String: "",
			Valid:  false,
		},
		TglDaftarMhs: carbon.Now(),
	}
	data := Pendaftaran(context.Background(), "MARIA", "rofiganteng", insdata)
	fmt.Println(data)
}

func TestGetAllProvinsi(t *testing.T) {
	data, err := GetAllProvinsi(context.Background(), "MARIA")

	if err != nil {
		t.Errorf("Error in GetAllProvinsi: %v", err)
	}

	fmt.Println(data)
}

func TestGetAllKota(t *testing.T) {
	data, err := GetAllKota(context.Background(), "MARIA")

	if err != nil {
		t.Errorf("Error in GetAllProvinsi: %v", err)
	}

	fmt.Println(data)
}
