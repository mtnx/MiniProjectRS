package endpoint

import (
	"context"
	"fmt"

	sv "rumahsakit/obat/server"
)

func (ke ObatEndpoint) AddObatService(ctx context.Context, obat sv.Obat) error {
	_, err := ke.AddObatEndpoint(ctx, obat)
	return err
}

func (ke ObatEndpoint) ReadObatByNamaService(ctx context.Context, nama_obat string) (sv.Obat, error) {
	req := sv.Obat{NamaObat: nama_obat}
	fmt.Println(req)
	resp, err := ke.ReadObatByNamaEndpoint(ctx, req)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	cus := resp.(sv.Obat)
	return cus, err
}

func (ke ObatEndpoint) ReadObatService(ctx context.Context) (sv.Obats, error) {
	resp, err := ke.ReadObatEndpoint(ctx, nil)
	fmt.Println("ke resp", resp)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	return resp.(sv.Obats), err
}

func (ke ObatEndpoint) UpdateObatService(ctx context.Context, kar sv.Obat) error {
	_, err := ke.UpdateObatEndpoint(ctx, kar)
	if err != nil {
		fmt.Println("error pada endpoint:", err)
	}
	return err
}
