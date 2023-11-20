package admin_user

import (
	"context"
	"github.com/PMBULBI/pendaftaran"
	pmbulbi "github.com/PMBULBI/types/schemas"
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
	"math/rand"
	"time"
)

func GetAllAdm(ctx context.Context, Mariaconn *gorm.DB) (dataLvlAdm []pmbulbi.Admin, err error) {
	val, err := GetAllAdmin(Mariaconn, ctx)

	if err != nil {
		return nil, err
	}

	for _, v := range val {
		dataLvlAdm = append(dataLvlAdm, pmbulbi.Admin{
			IDAdmin:       v.IDAdmin,
			NamaAdmin:     v.NamaAdmin,
			UsernameAdmin: v.UsernameAdmin,
			Email:         v.Email,
			NoHp:          v.NoHp,
			Password:      v.Password,
			TglBuatAkun:   v.TglBuatAkun,
			IsAktif:       v.IsAktif,
			Level:         v.Level,
			KodeRef:       v.KodeRef,
		})
	}

	return dataLvlAdm, nil
}

func GetOneAdm(ctx context.Context, Mariaconn *gorm.DB, id string) (data pmbulbi.Admin, err error) {

	adm, err := GetOneAdmin(Mariaconn, ctx, id)
	if err != nil {
		return pmbulbi.Admin{}, err
	}

	return adm, nil
}

const letterBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateRandomString(length int) string {
	rand.NewSource(time.Now().UnixNano())

	stringBytes := make([]byte, length)

	for i := 0; i < length; i++ {
		stringBytes[i] = letterBytes[rand.Intn(len(letterBytes))]
	}

	return string(stringBytes)
}

func InsertAdm(ctx context.Context, Mariaconn *gorm.DB, secret string, val pmbulbi.Admin) (err error) {
	kodeku := GenerateRandomString(8)
	mypass := val.Password
	passwordencrypted := pendaftaran.Encrypt(mypass, secret)

	data := pmbulbi.Admin{
		IDAdmin:       val.IDAdmin,
		NamaAdmin:     val.NamaAdmin,
		UsernameAdmin: val.UsernameAdmin,
		Email:         val.Email,
		NoHp:          val.NoHp,
		Password:      passwordencrypted,
		TglBuatAkun:   carbon.Now(),
		IsAktif:       val.IsAktif,
		Level:         val.Level,
		KodeRef:       kodeku,
	}
	err = InsertAdmin(Mariaconn, ctx, data)
	if err != nil {
		return err
	}
	return
}

func UpdateAdm(ctx context.Context, Mariaconn *gorm.DB, id, secret string, val pmbulbi.Admin) (err error) {
	mypass := val.Password
	passwordencrypted := pendaftaran.Encrypt(mypass, secret)

	data := pmbulbi.Admin{
		IDAdmin:       val.IDAdmin,
		NamaAdmin:     val.NamaAdmin,
		UsernameAdmin: val.UsernameAdmin,
		Email:         val.Email,
		NoHp:          val.NoHp,
		Password:      passwordencrypted,
		TglBuatAkun:   carbon.Now(),
		IsAktif:       val.IsAktif,
		Level:         val.Level,
	}

	err = UpdateAdmin(Mariaconn, ctx, id, data)
	if err != nil {
		return err
	}
	return
}

func DeleteAdm(ctx context.Context, Mariaconn *gorm.DB, id string) (err error) {
	err = DeleteAdmin(Mariaconn, ctx, id)
	if err != nil {
		return err
	}
	return
}
