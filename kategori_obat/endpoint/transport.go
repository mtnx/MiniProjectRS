package endpoint

import (
	"context"

	scv "rumahsakit/kategori_obat/server"

	pb "rumahsakit/kategori_obat/grpc"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	stdopentracing "github.com/opentracing/opentracing-go"
	oldcontext "golang.org/x/net/context"
)

type grpcKategoriObatServer struct {
	addKategoriObat        grpctransport.Handler
	readKategoriObatByNama grpctransport.Handler
	readKategoriObat       grpctransport.Handler
	updateKategoriObat     grpctransport.Handler
}

func NewGRPCKategoriObatServer(endpoints KategoriObatEndpoint, tracer stdopentracing.Tracer,
	logger log.Logger) pb.KategoriObatServiceServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
	}
	return &grpcKategoriObatServer{
		addKategoriObat: grpctransport.NewServer(endpoints.AddKategoriObatEndpoint,
			decodeAddKategoriObatRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "AddKategoriObat", logger)))...),
		readKategoriObatByNama: grpctransport.NewServer(endpoints.ReadKategoriObatByNamaEndpoint,
			decodeReadKategoriObatByNamaRequest,
			encodeReadKategoriObatByNamaResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadKategoriObatByNama", logger)))...),
		readKategoriObat: grpctransport.NewServer(endpoints.ReadKategoriObatEndpoint,
			decodeReadKategoriObatRequest,
			encodeReadKategoriObatResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadKategoriObat", logger)))...),
		updateKategoriObat: grpctransport.NewServer(endpoints.UpdateKategoriObatEndpoint,
			decodeUpdateKategoriObatRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "UpdateKategoriObat", logger)))...),
	}
}

func decodeAddKategoriObatRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.AddKategoriObatReq)
	return scv.KategoriObat{KodeKategoriObat: req.GetKodeKategoriObat(), NamaKategoriObat: req.GetNamaKategoriObat(), Deskripsi: req.GetDeskripsi(), KodeObat: req.GetKodeObat(),
		CreatedBy: req.GetCreatedBy(), CreatedOn: req.GetCreatedOn(), UpdateBy: req.GetUpdateBy(), UpdateOn: req.GetUpdateOn(), Status: req.GetStatus()}, nil
}

func encodeEmptyResponse(_ context.Context, response interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func (s *grpcKategoriObatServer) AddKategoriObat(ctx oldcontext.Context, kategori_obat *pb.AddKategoriObatReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.addKategoriObat.ServeGRPC(ctx, kategori_obat)
	if err != nil {
		return nil, err
	}
	return resp.(*google_protobuf.Empty), nil
}

func decodeReadKategoriObatByNamaRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ReadKategoriObatByNamaReq)
	return scv.KategoriObat{NamaKategoriObat: req.NamaKategoriObat}, nil
}

func decodeReadKategoriObatRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, nil
}

func encodeReadKategoriObatByNamaResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.KategoriObat)
	return &pb.ReadKategoriObatByNamaResp{KodeKategoriObat: resp.KodeKategoriObat, NamaKategoriObat: resp.NamaKategoriObat, Deskripsi: resp.Deskripsi,
		KodeObat: resp.KodeObat, CreatedBy: resp.CreatedBy, CreatedOn: resp.CreatedOn, UpdateBy: resp.UpdateBy, UpdateOn: resp.UpdateOn, Status: resp.Status}, nil
}

func encodeReadKategoriObatResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.KategoriObats)

	rsp := &pb.ReadKategoriObatResp{}

	for _, v := range resp {
		itm := &pb.ReadKategoriObatByNamaResp{

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
		rsp.AllKategoriObat = append(rsp.AllKategoriObat, itm)
	}
	return rsp, nil
}

func (s *grpcKategoriObatServer) ReadKategoriObatByNama(ctx oldcontext.Context, nama_kategori_obat *pb.ReadKategoriObatByNamaReq) (*pb.ReadKategoriObatByNamaResp, error) {
	_, resp, err := s.readKategoriObatByNama.ServeGRPC(ctx, nama_kategori_obat)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadKategoriObatByNamaResp), nil
}

func (s *grpcKategoriObatServer) ReadKategoriObat(ctx oldcontext.Context, e *google_protobuf.Empty) (*pb.ReadKategoriObatResp, error) {
	_, resp, err := s.readKategoriObat.ServeGRPC(ctx, e)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadKategoriObatResp), nil
}

func decodeUpdateKategoriObatRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.UpdateKategoriObatReq)
	return scv.KategoriObat{KodeKategoriObat: req.KodeKategoriObat, NamaKategoriObat: req.NamaKategoriObat, Deskripsi: req.Deskripsi,
		KodeObat: req.KodeObat, CreatedBy: req.CreatedBy, CreatedOn: req.CreatedOn, UpdateBy: req.UpdateBy, UpdateOn: req.UpdateOn, Status: req.Status}, nil
}

func (s *grpcKategoriObatServer) UpdateKategoriObat(ctx oldcontext.Context, cus *pb.UpdateKategoriObatReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.updateKategoriObat.ServeGRPC(ctx, cus)
	if err != nil {
		return &google_protobuf.Empty{}, err
	}
	return resp.(*google_protobuf.Empty), nil
}
