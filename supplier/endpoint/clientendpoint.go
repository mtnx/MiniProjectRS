package endpoint

import (
	"context"
	"fmt"

	sv "rumahsakit/supplier/server"
)

func (ke SupplierEndpoint) AddSupplierService(ctx context.Context, supplier sv.Supplier) error {
	_, err := ke.AddSupplierEndpoint(ctx, supplier)
	return err
}

func (ke SupplierEndpoint) ReadSupplierByNamaService(ctx context.Context, nama_supplier string) (sv.Supplier, error) {
	req := sv.Supplier{NamaSupplier: nama_supplier}
	fmt.Println(req)
	resp, err := ke.ReadSupplierByNamaEndpoint(ctx, req)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	cus := resp.(sv.Supplier)
	return cus, err
}

func (ke SupplierEndpoint) ReadSupplierService(ctx context.Context) (sv.Suppliers, error) {
	resp, err := ke.ReadSupplierEndpoint(ctx, nil)
	fmt.Println("ke resp", resp)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	return resp.(sv.Suppliers), err
}

func (ke SupplierEndpoint) UpdateSupplierService(ctx context.Context, kar sv.Supplier) error {
	_, err := ke.UpdateSupplierEndpoint(ctx, kar)
	if err != nil {
		fmt.Println("error pada endpoint:", err)
	}
	return err
}
