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
