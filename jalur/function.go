package jalur

import (
	"context"
	pmbulbi "github.com/PMBULBI/types/schemas"
	"gorm.io/gorm"
)

func GetAllJalur(ctx context.Context, Mariaconn *gorm.DB) (data []pmbulbi.JalurPendaftaran, err error) {

	val, err := GetMAllJalur(Mariaconn, ctx)

	if err != nil {
		return nil, err
	}

	for _, v := range val {
		data = append(data, pmbulbi.JalurPendaftaran{
			IDJalur:         v.IDJalur,
			Jalur:           v.Jalur,
			NamaJalur:       v.NamaJalur,
			KeteranganJalur: v.KeteranganJalur,
			Status:          v.Status,
		})
	}

	return data, nil
}

func GetOneJalur(ctx context.Context, Mariaconn *gorm.DB, id string) (data pmbulbi.JalurPendaftaran, err error) {

	jalur, err := GetMOneJalur(Mariaconn, ctx, id)
	if err != nil {
		return pmbulbi.JalurPendaftaran{}, err
	}

	return jalur, nil
}

func InsertJalur(ctx context.Context, Mariaconn *gorm.DB, val pmbulbi.JalurPendaftaran) (err error) {
	err = InsJalur(Mariaconn, ctx, val)
	if err != nil {
		return err
	}
	return
}

func UpdateJalur(ctx context.Context, Mariaconn *gorm.DB, id string, val pmbulbi.JalurPendaftaran) (err error) {
	err = UpdJalur(Mariaconn, ctx, id, val)
	if err != nil {
		return err
	}
	return
}

func DeleteJalur(ctx context.Context, Mariaconn *gorm.DB, id string) (err error) {
	err = DelJalur(Mariaconn, ctx, id)
	if err != nil {
		return err
	}
	return
}
