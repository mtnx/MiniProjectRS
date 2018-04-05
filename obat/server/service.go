package server

import "context"

type Status int32

const (
	//ServiceID is dispatch service ID
	ServiceID = "Obat.Rumahsakit.id"
	//OnAdd     Status = 1
)

type Obat struct {
	KodeObat           string
	NamaObat           string
	TanggalKadaluwarsa string
	Harga              int64
	CreatedBy          string
	CreatedOn          string
	UpdateBy           string
	UpdateOn           string
	Status             int32
}
type Obats []Obat

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
	AddObat(Obat) error
	ReadObat() (Obats, error)
	UpdateObat(Obat) error
	ReadObatByNama(string) (Obat, error)
}

type ObatService interface {
	AddObatService(context.Context, Obat) error
	ReadObatService(context.Context) (Obats, error)
	UpdateObatService(context.Context, Obat) error
	ReadObatByNamaService(context.Context, string) (Obat, error)
}
