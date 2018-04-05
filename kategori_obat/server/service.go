package server

import "context"

type Status int32

const (
	//ServiceID is dispatch service ID
	ServiceID        = "KategoriObat.Rumahsakit.id"
	OnAdd     Status = 1
)

type KategoriObat struct {
	KodeKategoriObat string
	NamaKategoriObat string
	Deskripsi        string
	KodeObat         string
	CreatedBy        string
	CreatedOn        string
	UpdateBy         string
	UpdateOn         string
	Status           int32
}
type KategoriObats []KategoriObat

/*type Location struct {
	customerID   int64
	label        []int32
	locationType []int32
	name         []string
	street       string
	village      string
	district     string
	city         string
	province     string
	latitude     float64
	longitude    float64
}*/

type ReadWriter interface {
	AddKategoriObat(KategoriObat) error
	ReadKategoriObat() (KategoriObats, error)
	UpdateKategoriObat(KategoriObat) error
	ReadKategoriObatByNama(string) (KategoriObat, error)
}

type KategoriObatService interface {
	AddKategoriObatService(context.Context, KategoriObat) error
	ReadKategoriObatService(context.Context) (KategoriObats, error)
	UpdateKategoriObatService(context.Context, KategoriObat) error
	ReadKategoriObatByNamaService(context.Context, string) (KategoriObat, error)
}
