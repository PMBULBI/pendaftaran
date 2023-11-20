package jalur

import (
	"context"
	pmbulbi "github.com/PMBULBI/types/schemas"
	"gorm.io/gorm"
)

func GetAllJalurPendaftaran(ctx context.Context, Mariaconn *gorm.DB) (data []pmbulbi.JalurPendaftaran, err error) {

	val, err := GetAllJalur(Mariaconn, ctx)

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
