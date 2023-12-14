package jalur_berkas

import (
	"context"
	pmbulbi "github.com/PMBULBI/types/schemas"
)

func (r *Repository) Fetch(ctx context.Context) (val []pmbulbi.JalurBerkas, err error) {
	err = r.db.
		WithContext(ctx).
		Find(&val).
		Error
	return
}

func (r *Repository) GetBerkasByJalur(ctx context.Context, id_jalur int) (dest []pmbulbi.JalurBerkasJoin, err error) {
	err = r.db.
		WithContext(ctx).
		Table("jalur_berkas jaber").
		Joins("JOIN jalur_pendaftaran japen ON japen.id_jalur = jaber.id_jalur").
		Where("jaber.id_jalur = ?", id_jalur).
		Select("jaber.*, japen.*").
		Find(&dest).
		Error
	return
}
