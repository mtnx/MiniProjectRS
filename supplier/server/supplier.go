package server

import (
	"context"
)

type supplier struct {
	writer ReadWriter
}

func NewSupplier(writer ReadWriter) SupplierService {
	return &supplier{writer: writer}
}

//Methode pada interface MahasiswaService di service.go
func (c *supplier) AddSupplierService(ctx context.Context, supplier Supplier) error {
	//fmt.Println("mahasiswa")
	err := c.writer.AddSupplier(supplier)
	if err != nil {
		return err
	}

	return nil
}

func (c *supplier) ReadSupplierService(ctx context.Context) (Suppliers, error) {
	kar, err := c.writer.ReadSupplier()
	//fmt.Println("mahasiswa", mhs)
	if err != nil {
		return kar, err
	}
	return kar, nil
}

func (c *supplier) UpdateSupplierService(ctx context.Context, kar Supplier) error {
	err := c.writer.UpdateSupplier(kar)
	if err != nil {
		return err
	}
	return nil
}

func (c *supplier) ReadSupplierByNamaService(ctx context.Context, nama_supplier string) (Supplier, error) {
	kar, err := c.writer.ReadSupplierByNama(nama_supplier)
	//fmt.Println("mahasiswa:", mhs)
	if err != nil {
		return kar, err
	}
	return kar, nil
}
