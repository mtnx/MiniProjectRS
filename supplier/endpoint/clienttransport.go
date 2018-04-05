package endpoint

import (
	"context"
	"time"

	svc "rumahsakit/supplier/server"

	pb "rumahsakit/supplier/grpc"

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
	grpcName = "grpc.SupplierService"
)

func NewGRPCSupplierClient(nodes []string, creds credentials.TransportCredentials, option util.ClientOption,
	tracer stdopentracing.Tracer, logger log.Logger) (svc.SupplierService, error) {

	instancer, err := disc.ServiceDiscovery(nodes, svc.ServiceID, logger)
	if err != nil {
		return nil, err
	}

	retryMax := option.Retry
	retryTimeout := option.RetryTimeout
	timeout := option.Timeout

	var addSupplierEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientAddSupplierEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		addSupplierEp = retry
	}

	var readSupplierByNamaEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadSupplierByNamaEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readSupplierByNamaEp = retry
	}

	var readSupplierEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadSupplierEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readSupplierEp = retry
	}

	var updateSupplierEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientUpdateSupplier, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		updateSupplierEp = retry
	}
	return SupplierEndpoint{AddSupplierEndpoint: addSupplierEp, ReadSupplierByNamaEndpoint: readSupplierByNamaEp,
		ReadSupplierEndpoint: readSupplierEp, UpdateSupplierEndpoint: updateSupplierEp}, nil
}

func encodeAddSupplierRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Supplier)
	return &pb.AddSupplierReq{

		KodeSupplier:  req.KodeSupplier,
		NamaSupplier:  req.NamaSupplier,
		JenisSupplier: req.JenisSupplier,
		CreatedBy:     req.CreatedBy,
		CreatedOn:     req.CreatedOn,
		UpdateBy:      req.UpdateBy,
		UpdateOn:      req.UpdateOn,
		Status:        req.Status,
	}, nil
}

func encodeReadSupplierByNamaRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Supplier)
	return &pb.ReadSupplierByNamaReq{NamaSupplier: req.NamaSupplier}, nil
}

func encodeReadSupplierRequest(_ context.Context, request interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func encodeUpdateSupplierRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Supplier)
	return &pb.UpdateSupplierReq{

		KodeSupplier:  req.KodeSupplier,
		NamaSupplier:  req.NamaSupplier,
		JenisSupplier: req.JenisSupplier,

		CreatedBy: req.CreatedBy,
		CreatedOn: req.CreatedOn,
		UpdateBy:  req.UpdateBy,
		UpdateOn:  req.UpdateOn,
		Status:    req.Status,
	}, nil
}

func decodeSupplierResponse(_ context.Context, response interface{}) (interface{}, error) {
	return nil, nil
}

func decodeReadSupplierByNamaRespones(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadSupplierByNamaResp)
	return svc.Supplier{

		KodeSupplier:  resp.KodeSupplier,
		NamaSupplier:  resp.NamaSupplier,
		JenisSupplier: resp.JenisSupplier,

		CreatedBy: resp.CreatedBy,
		CreatedOn: resp.CreatedOn,
		UpdateBy:  resp.UpdateBy,
		UpdateOn:  resp.UpdateOn,
		Status:    resp.Status,
	}, nil
}

func decodeReadSupplierResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadSupplierResp)
	var rsp svc.Suppliers

	for _, v := range resp.AllSupplier {
		itm := svc.Supplier{

			KodeSupplier:  v.KodeSupplier,
			NamaSupplier:  v.NamaSupplier,
			JenisSupplier: v.JenisSupplier,

			CreatedBy: v.CreatedBy,
			CreatedOn: v.CreatedOn,
			UpdateBy:  v.UpdateBy,
			UpdateOn:  v.UpdateOn,
			Status:    v.Status,
		}
		rsp = append(rsp, itm)
	}
	return rsp, nil
}

func makeClientAddSupplierEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn,
		grpcName,
		"AddSupplier",
		encodeAddSupplierRequest,
		decodeSupplierResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "AddSupplier")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "AddSupplier",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadSupplierByNamaEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadSupplierByNama",
		encodeReadSupplierByNamaRequest,
		decodeReadSupplierByNamaRespones,
		pb.ReadSupplierByNamaResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadSupplierByNama")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadSupplierByNama",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadSupplierEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadSupplier",
		encodeReadSupplierRequest,
		decodeReadSupplierResponse,
		pb.ReadSupplierResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadSupplier")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadSupplier",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientUpdateSupplier(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {
	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"UpdateSupplier",
		encodeUpdateSupplierRequest,
		decodeSupplierResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "UpdateSupplier")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "UpdateSupplier",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}
