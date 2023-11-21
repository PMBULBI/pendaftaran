package fakultas

import (
	"context"
	pmbulbi "github.com/PMBULBI/types/schemas"
	"gorm.io/gorm"
)

func GetAllFakultas(ctx context.Context, Mariaconn *gorm.DB) (data []pmbulbi.Fakultas, err error) {

	val, err := GetMAllFakultas(Mariaconn, ctx)

	if err != nil {
		return nil, err
	}

	for _, v := range val {
		data = append(data, pmbulbi.Fakultas{
			IDFakultas:   v.IDFakultas,
			NamaFakultas: v.NamaFakultas,
		})
	}

	return data, nil
}

func GetOneFakultas(ctx context.Context, Mariaconn *gorm.DB, id string) (data pmbulbi.Fakultas, err error) {

	fakultas, err := GetMOneFakultas(Mariaconn, ctx, id)
	if err != nil {
		return pmbulbi.Fakultas{}, err
	}

	return fakultas, nil
}

func InsertFakultas(ctx context.Context, Mariaconn *gorm.DB, val pmbulbi.Fakultas) (err error) {

	data := pmbulbi.Fakultas{
		IDFakultas:   val.IDFakultas,
		NamaFakultas: val.NamaFakultas,
	}
	err = InsFakultas(Mariaconn, ctx, data)
	if err != nil {
		return err
	}
	return
}
