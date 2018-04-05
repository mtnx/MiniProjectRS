package endpoint

import (
	"context"
	"fmt"

	sv "rumahsakit/kategori_obat/server"
)

func (ke KategoriObatEndpoint) AddKategoriObatService(ctx context.Context, kategori_obat sv.KategoriObat) error {
	_, err := ke.AddKategoriObatEndpoint(ctx, kategori_obat)
	return err
}

func (ke KategoriObatEndpoint) ReadKategoriObatByNamaService(ctx context.Context, nama_kategori_obat string) (sv.KategoriObat, error) {
	req := sv.KategoriObat{NamaKategoriObat: nama_kategori_obat}
	fmt.Println(req)
	resp, err := ke.ReadKategoriObatByNamaEndpoint(ctx, req)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	cus := resp.(sv.KategoriObat)
	return cus, err
}

func (ke KategoriObatEndpoint) ReadKategoriObatService(ctx context.Context) (sv.KategoriObats, error) {
	resp, err := ke.ReadKategoriObatEndpoint(ctx, nil)
	fmt.Println("ke resp", resp)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	return resp.(sv.KategoriObats), err
}

func (ke KategoriObatEndpoint) UpdateKategoriObatService(ctx context.Context, kar sv.KategoriObat) error {
	_, err := ke.UpdateKategoriObatEndpoint(ctx, kar)
	if err != nil {
		fmt.Println("error pada endpoint:", err)
	}
	return err
}
