package daftar_sekolah

import (
	"context"
	pmbulbi "github.com/PMBULBI/types/schemas"
)

func (r *Repository) GetSekolahLimits(ctx context.Context, lim int, name string) (dest []pmbulbi.DaftarSekolah, err error) {
	err = r.db.
		WithContext(ctx).
		Limit(lim).
		Where("nama_sekolah ILIKE ?", "%"+name+"%").
		Find(&dest).
		Error
	return
}
