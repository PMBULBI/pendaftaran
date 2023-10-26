package pendaftaran

import (
	"context"
	"fmt"
	pmbulbi "github.com/PMBULBI/types/schemas"
	"github.com/golang-module/carbon/v2"
	"testing"
)

func TestPmbdb_TableMigrator(t *testing.T) {
	Migrate, err := TableMigrator("irc:rollyganteng@tcp(10.14.200.17:3307)/pmb_ulbi?parseTime=true", pmbulbi.DaftarSekolah{})
	fmt.Println(err)
	fmt.Println(Migrate)
}

func TestInsertDataPendaftar(t *testing.T) {
	conn := CreateMariaGormConnection("irc:rollyganteng@tcp(10.14.200.17:3307)/pmb_ulbi?parseTime=true")
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
	hasil := "TktBSVlaUDRJUQ=="
	encrypted := Encrypter(strnger)
	decrypted := Decrypter(hasil)

	fmt.Println("enc", encrypted)
	fmt.Println("dec", decrypted)
}

func TestPendaftaran(t *testing.T) {
	insdata := pmbulbi.Pendaftaran{
		NamaMhs:         "Testeeeee",
		AsalSekolah:     "Bandung High School",
		EmailMhs:        "Testing@gmail.com",
		HpMhs:           "62851666465",
		ProvinsiSekolah: "Jawa barat",
		KotaSekolah:     "Bandung",
		UsernameAdmin:   "rofi",
		TglDaftarMhs:    carbon.Now(),
	}
	data := Pendaftaran(context.Background(), "irc:rollyganteng@tcp(10.14.200.17:3307)/pmb_ulbi?parseTime=true", insdata)
	fmt.Println(data)
}
