package main

import (
	"context"
	"fmt"
	"time"

	cli "rumahsakit/supplier/endpoint"
	//svc "rumahsakit/supplier/server"
	opt "rumahsakit/util/grpc"
	util "rumahsakit/util/microservice"

	tr "github.com/opentracing/opentracing-go"
)

func main() {
	logger := util.Logger()
	tracer := tr.GlobalTracer()
	option := opt.ClientOption{Retry: 3, RetryTimeout: 500 * time.Millisecond, Timeout: 30 * time.Second}

	client, err := cli.NewGRPCSupplierClient([]string{"127.0.0.1:2181"}, nil, option, tracer, logger)
	if err != nil {
		logger.Log("error", err)
	}

	//Add Supplier
	//client.AddSupplierService(context.Background(), svc.Supplier{KodeSupplier: "KOB007", NamaSupplier: "Antangin", JenisSupplier: "2019-10-30", CreatedBy: "Hyu", CreatedOn: "2017-12-31 14:50:21", UpdateBy: "Tio", UpdateOn: "2020-12-31 14:50:21", Status: 1})

	//Get Supplier By Nama
	//cusNama, _ := client.ReadSupplierByNamaService(context.Background(), "Aktived")
	//fmt.Println("supplier based on nama_supplier:", cusNama)

	//List Supplier
	cuss, _ := client.ReadSupplierService(context.Background())
	fmt.Println("all suppliers:", cuss)

	//Update Supplier
	//client.UpdateSupplierService(context.Background(), svc.Supplier{IdAgama: 1, Alamat: "Ciledug", Jeniskelamin: "L", Status: 0, UpdateBy: "Admin2", NamaSupplier: "Ari Prakoso"})

}
