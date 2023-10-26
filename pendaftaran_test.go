package pendaftaran

import (
	"context"
	"fmt"
	pmbulbi "github.com/PMBULBI/types/schemas"
	"github.com/golang-module/carbon/v2"
	"testing"
)

func TestPmbdb_TableMigrator(t *testing.T) {
	Migrate, err := TableMigrator("MARIASTRING", pmbulbi.Pendaftaran{})
	fmt.Println(err)
	fmt.Println(Migrate)
}

func TestInsertDataPendaftar(t *testing.T) {
	conn := CreateMariaGormConnection("MARIASTRING")
	data := pmbulbi.Pendaftaran{
		NamaMhs:         "Test",
		AsalSekolah:     "Bandung High School",
		EmailMhs:        "Testing@gmail.com",
		HpMhs:           "62851666465",
		ProvinsiSekolah: "Jawa barat",
		KotaSekolah:     "Bandung",
		UsernameAdmin:   "rofi",
		TglDaftarMhs:    carbon.Now(),
	}
	ins := InsertDataPendaftar(conn, context.Background(), data)
	fmt.Println(ins)
}

func TestEncrypt(t *testing.T) {
	strnger := "rofiganteng"
	encrypted := Encrypter(strnger)
	decrypted := Decrypter(encrypted)

	fmt.Println("enc", encrypted)
	fmt.Println("dec", decrypted)
}

func TestPendaftaran(t *testing.T) {
	insdata := pmbulbi.Pendaftaran{
		NamaMhs:         "Test",
		AsalSekolah:     "Bandung High School",
		EmailMhs:        "Testing@gmail.com",
		HpMhs:           "62851666465",
		ProvinsiSekolah: "Jawa barat",
		KotaSekolah:     "Bandung",
		UsernameAdmin:   "rofi",
		TglDaftarMhs:    carbon.Now(),
	}
	data := Pendaftaran(context.Background(), "MARIASTRING", insdata)
	fmt.Println(data)
}
