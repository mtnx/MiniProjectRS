package main

import (
	"context"
	"fmt"
	cli "rumahsakit/kategori_obat/endpoint"
	//svc "rumahsakit/kategori_obat/server"
	opt "rumahsakit/util/grpc"
	util "rumahsakit/util/microservice"
	"time"

	tr "github.com/opentracing/opentracing-go"
)

func main() {
	logger := util.Logger()
	tracer := tr.GlobalTracer()
	option := opt.ClientOption{Retry: 3, RetryTimeout: 500 * time.Millisecond, Timeout: 30 * time.Second}

	client, err := cli.NewGRPCKategoriObatClient([]string{"127.0.0.1:2181"}, nil, option, tracer, logger)
	if err != nil {
		logger.Log("error", err)
	}

	//Add KategoriObat
	//client.AddKategoriObatService(context.Background(), svc.KategoriObat{KodeKategoriObat: "KOB007", NamaKategoriObat: "Antangin", Deskripsi: "Bukan Obat Kuat", KodeObat: "OBT003", CreatedBy: "Hyu", CreatedOn: "2017-12-31 14:50:21", UpdateBy: "Tio", UpdateOn: "2020-12-31 14:50:21", Status: 1})

	//Get KategoriObat By Nama
	cusNama, _ := client.ReadKategoriObatByNamaService(context.Background(), "Obat Keras")
	fmt.Println("kategori_obat based on nama_kategori_obat:", cusNama)

	//List KategoriObat
	//cuss, _ := client.ReadKategoriObatService(context.Background())
	//fmt.Println("all kategori_obats:", cuss)

	//Update KategoriObat
	//client.UpdateKategoriObatService(context.Background(), svc.KategoriObat{KodeKategoriObat: "KOB007", NamaKategoriObat: "Antangin", Deskripsi: "Bukan Obat Kuat", KodeObat: "OBT003", CreatedBy: "Hyu", CreatedOn: "2017-12-31 14:50:21", UpdateBy: "Tio", UpdateOn: "2020-12-31 14:50:21", Status: 1})

}
