package main

import (
	"context"
	"fmt"
	cli "rumahsakit/obat/endpoint"
	//svc "rumahsakit/obat/server"
	opt "rumahsakit/util/grpc"
	util "rumahsakit/util/microservice"
	"time"

	tr "github.com/opentracing/opentracing-go"
)

func main() {
	logger := util.Logger()
	tracer := tr.GlobalTracer()
	option := opt.ClientOption{Retry: 3, RetryTimeout: 500 * time.Millisecond, Timeout: 30 * time.Second}

	client, err := cli.NewGRPCObatClient([]string{"127.0.0.1:2181"}, nil, option, tracer, logger)
	if err != nil {
		logger.Log("error", err)
	}

	//Add Obat
	//client.AddObatService(context.Background(), svc.Obat{KodeObat: "OBT007", NamaObat: "OkiJelly", TanggalKadaluwarsa: "2019-10-30", Harga: 3750, CreatedBy: "Hyu", CreatedOn: "2019-10-30 09:45:21", UpdateBy: "Hyu", UpdateOn: "2020-10-21 23:00:21", Status: 1})

	//Get Obat By Nama
	//cusNama, _ := client.ReadObatByNamaService(context.Background(), "Aktived")
	//fmt.Println("obat based on nama_obat:", cusNama)

	//List Obat
	cuss, _ := client.ReadObatService(context.Background())
	fmt.Println("all obats:", cuss)

	//Update Obat
	//client.UpdateObatService(context.Background(), svc.Obat{NamaObat: "Marisa", TanggalKadaluwarsa: "2019-10-30", Harga: 3750, CreatedBy: "Hyu", CreatedOn: "2019-10-30 09:45:21", UpdateBy: "Hyu", UpdateOn: "2020-10-21 23:00:21", Status: 0, KodeObat: "KOB045"})

}
