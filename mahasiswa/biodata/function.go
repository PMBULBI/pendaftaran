package biodata

import (
	"context"
	"database/sql"
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

func InsertBiodataProdi(ctx context.Context, Mariaconn *gorm.DB, val pmbulbi.BiodataProdi) (err error) {

	data := pmbulbi.BiodataProdi{
		IdHash:     val.IdHash,
		IDFakultas: val.IDFakultas,
		Prodi1:     val.Prodi1,
		Prodi2:     val.Prodi2,
		Konsentrasi: sql.NullInt16{
			Int16: val.Konsentrasi.Int16,
			Valid: val.Konsentrasi.Int16 != 0,
		},
	}
	err = InsertBioProdi(Mariaconn, ctx, data)
	if err != nil {
		return err
	}
	return
}
