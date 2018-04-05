package endpoint

import (
	"context"
	"time"

	svc "rumahsakit/obat/server"

	pb "rumahsakit/obat/grpc"

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
	grpcName = "grpc.ObatService"
)

func NewGRPCObatClient(nodes []string, creds credentials.TransportCredentials, option util.ClientOption,
	tracer stdopentracing.Tracer, logger log.Logger) (svc.ObatService, error) {

	instancer, err := disc.ServiceDiscovery(nodes, svc.ServiceID, logger)
	if err != nil {
		return nil, err
	}

	retryMax := option.Retry
	retryTimeout := option.RetryTimeout
	timeout := option.Timeout

	var addObatEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientAddObatEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		addObatEp = retry
	}

	var readObatByNamaEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadObatByNamaEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readObatByNamaEp = retry
	}

	var readObatEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadObatEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readObatEp = retry
	}

	var updateObatEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientUpdateObat, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		updateObatEp = retry
	}
	return ObatEndpoint{AddObatEndpoint: addObatEp, ReadObatByNamaEndpoint: readObatByNamaEp,
		ReadObatEndpoint: readObatEp, UpdateObatEndpoint: updateObatEp}, nil
}

func encodeAddObatRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Obat)
	return &pb.AddObatReq{

		KodeObat:           req.KodeObat,
		NamaObat:           req.NamaObat,
		TanggalKadaluwarsa: req.TanggalKadaluwarsa,
		Harga:              req.Harga,
		CreatedBy:          req.CreatedBy,
		CreatedOn:          req.CreatedOn,
		UpdateBy:           req.UpdateBy,
		UpdateOn:           req.UpdateOn,
		Status:             req.Status,
	}, nil
}

func encodeReadObatByNamaRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Obat)
	return &pb.ReadObatByNamaReq{NamaObat: req.NamaObat}, nil
}

func encodeReadObatRequest(_ context.Context, request interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func encodeUpdateObatRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Obat)
	return &pb.UpdateObatReq{

		KodeObat:           req.KodeObat,
		NamaObat:           req.NamaObat,
		TanggalKadaluwarsa: req.TanggalKadaluwarsa,
		Harga:              req.Harga,
		CreatedBy:          req.CreatedBy,
		CreatedOn:          req.CreatedOn,
		UpdateBy:           req.UpdateBy,
		UpdateOn:           req.UpdateOn,
		Status:             req.Status,
	}, nil
}

func decodeObatResponse(_ context.Context, response interface{}) (interface{}, error) {
	return nil, nil
}

func decodeReadObatByNamaRespones(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadObatByNamaResp)
	return svc.Obat{

		KodeObat:           resp.KodeObat,
		NamaObat:           resp.NamaObat,
		TanggalKadaluwarsa: resp.TanggalKadaluwarsa,
		Harga:              resp.Harga,
		CreatedBy:          resp.CreatedBy,
		CreatedOn:          resp.CreatedOn,
		UpdateBy:           resp.UpdateBy,
		UpdateOn:           resp.UpdateOn,
		Status:             resp.Status,
	}, nil
}

func decodeReadObatResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadObatResp)
	var rsp svc.Obats

	for _, v := range resp.AllObat {
		itm := svc.Obat{

			KodeObat:           v.KodeObat,
			NamaObat:           v.NamaObat,
			TanggalKadaluwarsa: v.TanggalKadaluwarsa,
			Harga:              v.Harga,
			CreatedBy:          v.CreatedBy,
			CreatedOn:          v.CreatedOn,
			UpdateBy:           v.UpdateBy,
			UpdateOn:           v.UpdateOn,
			Status:             v.Status,
		}
		rsp = append(rsp, itm)
	}
	return rsp, nil
}

func makeClientAddObatEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn,
		grpcName,
		"AddObat",
		encodeAddObatRequest,
		decodeObatResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "AddObat")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "AddObat",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadObatByNamaEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadObatByNama",
		encodeReadObatByNamaRequest,
		decodeReadObatByNamaRespones,
		pb.ReadObatByNamaResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadObatByNama")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadObatByNama",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadObatEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadObat",
		encodeReadObatRequest,
		decodeReadObatResponse,
		pb.ReadObatResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadObat")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadObat",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientUpdateObat(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {
	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"UpdateObat",
		encodeUpdateObatRequest,
		decodeObatResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "UpdateObat")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "UpdateObat",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}
