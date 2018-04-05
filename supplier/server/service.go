package server

import "context"

type Status int32

const (
	//ServiceID is dispatch service ID
	ServiceID        = "Supplier.Rumahsakit.id"
	OnAdd     Status = 1
)

type Supplier struct {
	KodeSupplier  string
	NamaSupplier  string
	JenisSupplier string
	CreatedBy     string
	CreatedOn     string
	UpdateBy      string
	UpdateOn      string
	Status        int32
}
type Suppliers []Supplier

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
	AddSupplier(Supplier) error
	ReadSupplier() (Suppliers, error)
	UpdateSupplier(Supplier) error
	ReadSupplierByNama(string) (Supplier, error)
}

type SupplierService interface {
	AddSupplierService(context.Context, Supplier) error
	ReadSupplierService(context.Context) (Suppliers, error)
	UpdateSupplierService(context.Context, Supplier) error
	ReadSupplierByNamaService(context.Context, string) (Supplier, error)
}
