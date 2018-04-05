package endpoint

import (
	"context"

	scv "rumahsakit/supplier/server"

	pb "rumahsakit/supplier/grpc"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	stdopentracing "github.com/opentracing/opentracing-go"
	oldcontext "golang.org/x/net/context"
)

type grpcSupplierServer struct {
	addSupplier        grpctransport.Handler
	readSupplierByNama grpctransport.Handler
	readSupplier       grpctransport.Handler
	updateSupplier     grpctransport.Handler
}

func NewGRPCSupplierServer(endpoints SupplierEndpoint, tracer stdopentracing.Tracer,
	logger log.Logger) pb.SupplierServiceServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
	}
	return &grpcSupplierServer{
		addSupplier: grpctransport.NewServer(endpoints.AddSupplierEndpoint,
			decodeAddSupplierRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "AddSupplier", logger)))...),
		readSupplierByNama: grpctransport.NewServer(endpoints.ReadSupplierByNamaEndpoint,
			decodeReadSupplierByNamaRequest,
			encodeReadSupplierByNamaResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadSupplierByNama", logger)))...),
		readSupplier: grpctransport.NewServer(endpoints.ReadSupplierEndpoint,
			decodeReadSupplierRequest,
			encodeReadSupplierResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadSupplier", logger)))...),
		updateSupplier: grpctransport.NewServer(endpoints.UpdateSupplierEndpoint,
			decodeUpdateSupplierRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "UpdateSupplier", logger)))...),
	}
}

func decodeAddSupplierRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.AddSupplierReq)
	return scv.Supplier{KodeSupplier: req.GetKodeSupplier(), NamaSupplier: req.GetNamaSupplier(), JenisSupplier: req.GetJenisSupplier(),
		CreatedBy: req.GetCreatedBy(), CreatedOn: req.GetCreatedOn(), UpdateBy: req.GetUpdateBy(), UpdateOn: req.GetUpdateOn(), Status: req.GetStatus()}, nil
}

func encodeEmptyResponse(_ context.Context, response interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func (s *grpcSupplierServer) AddSupplier(ctx oldcontext.Context, supplier *pb.AddSupplierReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.addSupplier.ServeGRPC(ctx, supplier)
	if err != nil {
		return nil, err
	}
	return resp.(*google_protobuf.Empty), nil
}

func decodeReadSupplierByNamaRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ReadSupplierByNamaReq)
	return scv.Supplier{NamaSupplier: req.NamaSupplier}, nil
}

func decodeReadSupplierRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, nil
}

func encodeReadSupplierByNamaResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Supplier)
	return &pb.ReadSupplierByNamaResp{KodeSupplier: resp.KodeSupplier, NamaSupplier: resp.NamaSupplier, JenisSupplier: resp.JenisSupplier,
		CreatedBy: resp.CreatedBy, CreatedOn: resp.CreatedOn, UpdateBy: resp.UpdateBy, UpdateOn: resp.UpdateOn, Status: resp.Status}, nil
}

func encodeReadSupplierResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Suppliers)

	rsp := &pb.ReadSupplierResp{}

	for _, v := range resp {
		itm := &pb.ReadSupplierByNamaResp{

			KodeSupplier:  v.KodeSupplier,
			NamaSupplier:  v.NamaSupplier,
			JenisSupplier: v.JenisSupplier,

			CreatedBy: v.CreatedBy,
			CreatedOn: v.CreatedOn,
			UpdateBy:  v.UpdateBy,
			UpdateOn:  v.UpdateOn,
			Status:    v.Status,
		}
		rsp.AllSupplier = append(rsp.AllSupplier, itm)
	}
	return rsp, nil
}

func (s *grpcSupplierServer) ReadSupplierByNama(ctx oldcontext.Context, nama_supplier *pb.ReadSupplierByNamaReq) (*pb.ReadSupplierByNamaResp, error) {
	_, resp, err := s.readSupplierByNama.ServeGRPC(ctx, nama_supplier)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadSupplierByNamaResp), nil
}

func (s *grpcSupplierServer) ReadSupplier(ctx oldcontext.Context, e *google_protobuf.Empty) (*pb.ReadSupplierResp, error) {
	_, resp, err := s.readSupplier.ServeGRPC(ctx, e)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadSupplierResp), nil
}

func decodeUpdateSupplierRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.UpdateSupplierReq)
	return scv.Supplier{KodeSupplier: req.KodeSupplier, NamaSupplier: req.NamaSupplier, JenisSupplier: req.JenisSupplier,
		CreatedBy: req.CreatedBy, CreatedOn: req.CreatedOn, UpdateBy: req.UpdateBy, UpdateOn: req.UpdateOn, Status: req.Status}, nil
}

func (s *grpcSupplierServer) UpdateSupplier(ctx oldcontext.Context, cus *pb.UpdateSupplierReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.updateSupplier.ServeGRPC(ctx, cus)
	if err != nil {
		return &google_protobuf.Empty{}, err
	}
	return resp.(*google_protobuf.Empty), nil
}
