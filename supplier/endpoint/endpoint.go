package endpoint

import (
	"context"

	svc "rumahsakit/supplier/server"

	kit "github.com/go-kit/kit/endpoint"
)

type SupplierEndpoint struct {
	AddSupplierEndpoint        kit.Endpoint
	ReadSupplierByNamaEndpoint kit.Endpoint
	ReadSupplierEndpoint       kit.Endpoint
	UpdateSupplierEndpoint     kit.Endpoint
}

func NewSupplierEndpoint(service svc.SupplierService) SupplierEndpoint {
	addSupplierEp := makeAddSupplierEndpoint(service)
	readSupplierByNamaEp := makeReadSupplierByNamaEndpoint(service)
	readSupplierEp := makeReadSupplierEndpoint(service)
	updateSupplierEp := makeUpdateSupplierEndpoint(service)
	return SupplierEndpoint{AddSupplierEndpoint: addSupplierEp,
		ReadSupplierByNamaEndpoint: readSupplierByNamaEp,
		ReadSupplierEndpoint:       readSupplierEp,
		UpdateSupplierEndpoint:     updateSupplierEp,
	}
}

func makeAddSupplierEndpoint(service svc.SupplierService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Supplier)
		err := service.AddSupplierService(ctx, req)
		return nil, err
	}
}

func makeReadSupplierByNamaEndpoint(service svc.SupplierService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Supplier)
		result, err := service.ReadSupplierByNamaService(ctx, req.NamaSupplier)
		return result, err
	}
}

func makeReadSupplierEndpoint(service svc.SupplierService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		result, err := service.ReadSupplierService(ctx)
		return result, err
	}
}

func makeUpdateSupplierEndpoint(service svc.SupplierService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Supplier)
		err := service.UpdateSupplierService(ctx, req)
		return nil, err
	}
}
