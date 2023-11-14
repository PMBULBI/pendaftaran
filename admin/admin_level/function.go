package admin_level

import (
	"context"
	"github.com/PMBULBI/pendaftaran"
	pmbulbi "github.com/PMBULBI/types/schemas"
)

func GetAllLevelAdm(ctx context.Context, Mariaenv string) (dataLvlAdm []pmbulbi.AdminLevel, err error) {
	conn := pendaftaran.CreateMariaGormConnection(Mariaenv)

	val, err := GetAllLevelAdmin(conn, ctx)

	if err != nil {
		return nil, err
	}

	for _, v := range val {
		dataLvlAdm = append(dataLvlAdm, pmbulbi.AdminLevel{
			IDLevel:   v.IDLevel,
			NamaLevel: v.NamaLevel,
		})
	}

	return dataLvlAdm, nil
}

func GetOneLevelAdm(ctx context.Context, Mariaenv, id string) (data pmbulbi.AdminLevel, err error) {
	conn := pendaftaran.CreateMariaGormConnection(Mariaenv)

	lvladm, err := GetOneLevelAdmin(conn, ctx, id)
	if err != nil {
		return pmbulbi.AdminLevel{}, err
	}

	return lvladm, nil
}
