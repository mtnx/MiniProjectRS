package endpoint

import (
	"context"

	scv "rumahsakit/obat/server"

	pb "rumahsakit/obat/grpc"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	stdopentracing "github.com/opentracing/opentracing-go"
	oldcontext "golang.org/x/net/context"
)

type grpcObatServer struct {
	addObat        grpctransport.Handler
	readObatByNama grpctransport.Handler
	readObat       grpctransport.Handler
	updateObat     grpctransport.Handler
}

func NewGRPCObatServer(endpoints ObatEndpoint, tracer stdopentracing.Tracer,
	logger log.Logger) pb.ObatServiceServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
	}
	return &grpcObatServer{
		addObat: grpctransport.NewServer(endpoints.AddObatEndpoint,
			decodeAddObatRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "AddObat", logger)))...),
		readObatByNama: grpctransport.NewServer(endpoints.ReadObatByNamaEndpoint,
			decodeReadObatByNamaRequest,
			encodeReadObatByNamaResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadObatByNama", logger)))...),
		readObat: grpctransport.NewServer(endpoints.ReadObatEndpoint,
			decodeReadObatRequest,
			encodeReadObatResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadObat", logger)))...),
		updateObat: grpctransport.NewServer(endpoints.UpdateObatEndpoint,
			decodeUpdateObatRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "UpdateObat", logger)))...),
	}
}

func decodeAddObatRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.AddObatReq)
	return scv.Obat{KodeObat: req.GetKodeObat(), NamaObat: req.GetNamaObat(), TanggalKadaluwarsa: req.GetTanggalKadaluwarsa(), Harga: req.GetHarga(),
		CreatedBy: req.GetCreatedBy(), CreatedOn: req.GetCreatedOn(), UpdateBy: req.GetUpdateBy(), UpdateOn: req.GetUpdateOn(), Status: req.GetStatus()}, nil
}

func encodeEmptyResponse(_ context.Context, response interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func (s *grpcObatServer) AddObat(ctx oldcontext.Context, obat *pb.AddObatReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.addObat.ServeGRPC(ctx, obat)
	if err != nil {
		return nil, err
	}
	return resp.(*google_protobuf.Empty), nil
}

func decodeReadObatByNamaRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ReadObatByNamaReq)
	return scv.Obat{NamaObat: req.NamaObat}, nil
}

func decodeReadObatRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, nil
}

func encodeReadObatByNamaResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Obat)
	return &pb.ReadObatByNamaResp{KodeObat: resp.KodeObat, NamaObat: resp.NamaObat, TanggalKadaluwarsa: resp.TanggalKadaluwarsa,
		Harga: resp.Harga, CreatedBy: resp.CreatedBy, CreatedOn: resp.CreatedOn, UpdateBy: resp.UpdateBy, UpdateOn: resp.UpdateOn, Status: resp.Status}, nil
}

func encodeReadObatResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Obats)

	rsp := &pb.ReadObatResp{}

	for _, v := range resp {
		itm := &pb.ReadObatByNamaResp{

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
		rsp.AllObat = append(rsp.AllObat, itm)
	}
	return rsp, nil
}

func (s *grpcObatServer) ReadObatByNama(ctx oldcontext.Context, nama_obat *pb.ReadObatByNamaReq) (*pb.ReadObatByNamaResp, error) {
	_, resp, err := s.readObatByNama.ServeGRPC(ctx, nama_obat)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadObatByNamaResp), nil
}

func (s *grpcObatServer) ReadObat(ctx oldcontext.Context, e *google_protobuf.Empty) (*pb.ReadObatResp, error) {
	_, resp, err := s.readObat.ServeGRPC(ctx, e)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadObatResp), nil
}

func decodeUpdateObatRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.UpdateObatReq)
	return scv.Obat{KodeObat: req.KodeObat, NamaObat: req.NamaObat, TanggalKadaluwarsa: req.TanggalKadaluwarsa,
		Harga: req.Harga, CreatedBy: req.CreatedBy, CreatedOn: req.CreatedOn, UpdateBy: req.UpdateBy, UpdateOn: req.UpdateOn, Status: req.Status}, nil
}

func (s *grpcObatServer) UpdateObat(ctx oldcontext.Context, cus *pb.UpdateObatReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.updateObat.ServeGRPC(ctx, cus)
	if err != nil {
		return &google_protobuf.Empty{}, err
	}
	return resp.(*google_protobuf.Empty), nil
}
