package biodata

import (
	"context"
	pmbulbi "github.com/PMBULBI/types/schemas"
	"gorm.io/gorm"
)

func InsertBiodataJalur(ctx context.Context, Mariaconn *gorm.DB, val pmbulbi.BiodataJalur) (err error) {

	data := pmbulbi.BiodataJalur{
		IdHash:     val.IdHash,
		IDJalur:    val.IDJalur,
		TahunLulus: val.TahunLulus,
		KodeRef:    val.KodeRef,
	}
	err = InsertBioJalur(Mariaconn, ctx, data)
	if err != nil {
		return err
	}
	return
}
