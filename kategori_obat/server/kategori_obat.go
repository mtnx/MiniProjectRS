package server

import (
	"context"
)

type kategori_obat struct {
	writer ReadWriter
}

func NewKategoriObat(writer ReadWriter) KategoriObatService {
	return &kategori_obat{writer: writer}
}

//Methode pada interface MahasiswaService di service.go
func (c *kategori_obat) AddKategoriObatService(ctx context.Context, kategori_obat KategoriObat) error {
	//fmt.Println("mahasiswa")
	err := c.writer.AddKategoriObat(kategori_obat)
	if err != nil {
		return err
	}

	return nil
}

func (c *kategori_obat) ReadKategoriObatService(ctx context.Context) (KategoriObats, error) {
	kar, err := c.writer.ReadKategoriObat()
	//fmt.Println("mahasiswa", mhs)
	if err != nil {
		return kar, err
	}
	return kar, nil
}

func (c *kategori_obat) UpdateKategoriObatService(ctx context.Context, kar KategoriObat) error {
	err := c.writer.UpdateKategoriObat(kar)
	if err != nil {
		return err
	}
	return nil
}

func (c *kategori_obat) ReadKategoriObatByNamaService(ctx context.Context, nama_kategori_obat string) (KategoriObat, error) {
	kar, err := c.writer.ReadKategoriObatByNama(nama_kategori_obat)
	//fmt.Println("mahasiswa:", mhs)
	if err != nil {
		return kar, err
	}
	return kar, nil
}
