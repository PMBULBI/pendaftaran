package admin_user

import (
	"context"
	"github.com/PMBULBI/pendaftaran"
	pmbulbi "github.com/PMBULBI/types/schemas"
	"github.com/golang-module/carbon/v2"
	"math/rand"
	"time"
)

func GetAllAdm(ctx context.Context, Mariaenv string) (dataLvlAdm []pmbulbi.Admin, err error) {
	conn := pendaftaran.CreateMariaGormConnection(Mariaenv)

	val, err := GetAllAdmin(conn, ctx)

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

func GetOneAdm(ctx context.Context, Mariaenv, id string) (data pmbulbi.Admin, err error) {
	conn := pendaftaran.CreateMariaGormConnection(Mariaenv)

	adm, err := GetOneAdmin(conn, ctx, id)
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

func InsertAdm(ctx context.Context, Mariaenv, secret string, val pmbulbi.Admin) (err error) {
	conn := pendaftaran.CreateMariaGormConnection(Mariaenv)

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
	err = InsertAdmin(conn, ctx, data)
	if err != nil {
		return err
	}
	return
}
