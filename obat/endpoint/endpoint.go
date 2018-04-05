package endpoint

import (
	"context"

	svc "rumahsakit/obat/server"

	kit "github.com/go-kit/kit/endpoint"
)

type ObatEndpoint struct {
	AddObatEndpoint        kit.Endpoint
	ReadObatByNamaEndpoint kit.Endpoint
	ReadObatEndpoint       kit.Endpoint
	UpdateObatEndpoint     kit.Endpoint
}

func NewObatEndpoint(service svc.ObatService) ObatEndpoint {
	addObatEp := makeAddObatEndpoint(service)
	readObatByNamaEp := makeReadObatByNamaEndpoint(service)
	readObatEp := makeReadObatEndpoint(service)
	updateObatEp := makeUpdateObatEndpoint(service)
	return ObatEndpoint{AddObatEndpoint: addObatEp,
		ReadObatByNamaEndpoint: readObatByNamaEp,
		ReadObatEndpoint:       readObatEp,
		UpdateObatEndpoint:     updateObatEp,
	}
}

func makeAddObatEndpoint(service svc.ObatService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Obat)
		err := service.AddObatService(ctx, req)
		return nil, err
	}
}

func makeReadObatByNamaEndpoint(service svc.ObatService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Obat)
		result, err := service.ReadObatByNamaService(ctx, req.NamaObat)
		return result, err
	}
}

func makeReadObatEndpoint(service svc.ObatService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		result, err := service.ReadObatService(ctx)
		return result, err
	}
}

func makeUpdateObatEndpoint(service svc.ObatService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Obat)
		err := service.UpdateObatService(ctx, req)
		return nil, err
	}
}
