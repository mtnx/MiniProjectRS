package endpoint

import (
	"context"
	"time"

	svc "rumahsakit/kategori_obat/server"

	pb "rumahsakit/kategori_obat/grpc"

	util "rumahsakit/util/grpc"
	disc "rumahsakit/util/microservice"

	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/lb"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	stdopentracing "github.com/opentracing/opentracing-go"
	"github.com/sony/gobreaker"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	grpcName = "grpc.KategoriObatService"
)

func NewGRPCKategoriObatClient(nodes []string, creds credentials.TransportCredentials, option util.ClientOption,
	tracer stdopentracing.Tracer, logger log.Logger) (svc.KategoriObatService, error) {

	instancer, err := disc.ServiceDiscovery(nodes, svc.ServiceID, logger)
	if err != nil {
		return nil, err
	}

	retryMax := option.Retry
	retryTimeout := option.RetryTimeout
	timeout := option.Timeout

	var addKategoriObatEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientAddKategoriObatEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		addKategoriObatEp = retry
	}

	var readKategoriObatByNamaEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadKategoriObatByNamaEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readKategoriObatByNamaEp = retry
	}

	var readKategoriObatEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadKategoriObatEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readKategoriObatEp = retry
	}

	var updateKategoriObatEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientUpdateKategoriObat, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		updateKategoriObatEp = retry
	}
	return KategoriObatEndpoint{AddKategoriObatEndpoint: addKategoriObatEp, ReadKategoriObatByNamaEndpoint: readKategoriObatByNamaEp,
		ReadKategoriObatEndpoint: readKategoriObatEp, UpdateKategoriObatEndpoint: updateKategoriObatEp}, nil
}

func encodeAddKategoriObatRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.KategoriObat)
	return &pb.AddKategoriObatReq{

		KodeKategoriObat: req.KodeKategoriObat,
		NamaKategoriObat: req.NamaKategoriObat,
		Deskripsi:        req.Deskripsi,
		KodeObat:         req.KodeObat,
		CreatedBy:        req.CreatedBy,
		CreatedOn:        req.CreatedOn,
		UpdateBy:         req.UpdateBy,
		UpdateOn:         req.UpdateOn,
		Status:           req.Status,
	}, nil
}

func encodeReadKategoriObatByNamaRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.KategoriObat)
	return &pb.ReadKategoriObatByNamaReq{NamaKategoriObat: req.NamaKategoriObat}, nil
}

func encodeReadKategoriObatRequest(_ context.Context, request interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func encodeUpdateKategoriObatRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.KategoriObat)
	return &pb.UpdateKategoriObatReq{

		KodeKategoriObat: req.KodeKategoriObat,
		NamaKategoriObat: req.NamaKategoriObat,
		Deskripsi:        req.Deskripsi,
		KodeObat:         req.KodeObat,
		CreatedBy:        req.CreatedBy,
		CreatedOn:        req.CreatedOn,
		UpdateBy:         req.UpdateBy,
		UpdateOn:         req.UpdateOn,
		Status:           req.Status,
	}, nil
}

func decodeKategoriObatResponse(_ context.Context, response interface{}) (interface{}, error) {
	return nil, nil
}

func decodeReadKategoriObatByNamaRespones(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadKategoriObatByNamaResp)
	return svc.KategoriObat{

		KodeKategoriObat: resp.KodeKategoriObat,
		NamaKategoriObat: resp.NamaKategoriObat,
		Deskripsi:        resp.Deskripsi,
		KodeObat:         resp.KodeObat,
		CreatedBy:        resp.CreatedBy,
		CreatedOn:        resp.CreatedOn,
		UpdateBy:         resp.UpdateBy,
		UpdateOn:         resp.UpdateOn,
		Status:           resp.Status,
	}, nil
}

func decodeReadKategoriObatResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadKategoriObatResp)
	var rsp svc.KategoriObats

	for _, v := range resp.AllKategoriObat {
		itm := svc.KategoriObat{

			KodeKategoriObat: v.KodeKategoriObat,
			NamaKategoriObat: v.NamaKategoriObat,
			Deskripsi:        v.Deskripsi,
			KodeObat:         v.KodeObat,
			CreatedBy:        v.CreatedBy,
			CreatedOn:        v.CreatedOn,
			UpdateBy:         v.UpdateBy,
			UpdateOn:         v.UpdateOn,
			Status:           v.Status,
		}
		rsp = append(rsp, itm)
	}
	return rsp, nil
}

func makeClientAddKategoriObatEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn,
		grpcName,
		"AddKategoriObat",
		encodeAddKategoriObatRequest,
		decodeKategoriObatResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "AddKategoriObat")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "AddKategoriObat",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadKategoriObatByNamaEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadKategoriObatByNama",
		encodeReadKategoriObatByNamaRequest,
		decodeReadKategoriObatByNamaRespones,
		pb.ReadKategoriObatByNamaResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadKategoriObatByNama")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadKategoriObatByNama",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadKategoriObatEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadKategoriObat",
		encodeReadKategoriObatRequest,
		decodeReadKategoriObatResponse,
		pb.ReadKategoriObatResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadKategoriObat")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadKategoriObat",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientUpdateKategoriObat(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {
	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"UpdateKategoriObat",
		encodeUpdateKategoriObatRequest,
		decodeKategoriObatResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "UpdateKategoriObat")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "UpdateKategoriObat",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}
