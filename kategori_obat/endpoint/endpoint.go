package endpoint

import (
	"context"

	svc "rumahsakit/kategori_obat/server"

	kit "github.com/go-kit/kit/endpoint"
)

type KategoriObatEndpoint struct {
	AddKategoriObatEndpoint        kit.Endpoint
	ReadKategoriObatByNamaEndpoint kit.Endpoint
	ReadKategoriObatEndpoint       kit.Endpoint
	UpdateKategoriObatEndpoint     kit.Endpoint
}

func NewKategoriObatEndpoint(service svc.KategoriObatService) KategoriObatEndpoint {
	addKategoriObatEp := makeAddKategoriObatEndpoint(service)
	readKategoriObatByNamaEp := makeReadKategoriObatByNamaEndpoint(service)
	readKategoriObatEp := makeReadKategoriObatEndpoint(service)
	updateKategoriObatEp := makeUpdateKategoriObatEndpoint(service)
	return KategoriObatEndpoint{AddKategoriObatEndpoint: addKategoriObatEp,
		ReadKategoriObatByNamaEndpoint: readKategoriObatByNamaEp,
		ReadKategoriObatEndpoint:       readKategoriObatEp,
		UpdateKategoriObatEndpoint:     updateKategoriObatEp,
	}
}

func makeAddKategoriObatEndpoint(service svc.KategoriObatService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.KategoriObat)
		err := service.AddKategoriObatService(ctx, req)
		return nil, err
	}
}

func makeReadKategoriObatByNamaEndpoint(service svc.KategoriObatService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.KategoriObat)
		result, err := service.ReadKategoriObatByNamaService(ctx, req.NamaKategoriObat)
		return result, err
	}
}

func makeReadKategoriObatEndpoint(service svc.KategoriObatService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		result, err := service.ReadKategoriObatService(ctx)
		return result, err
	}
}

func makeUpdateKategoriObatEndpoint(service svc.KategoriObatService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.KategoriObat)
		err := service.UpdateKategoriObatService(ctx, req)
		return nil, err
	}
}
