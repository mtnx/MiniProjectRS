package server

import (
	"context"
)

type obat struct {
	writer ReadWriter
}

func NewObat(writer ReadWriter) ObatService {
	return &obat{writer: writer}
}

//Methode pada interface MahasiswaService di service.go
func (c *obat) AddObatService(ctx context.Context, obat Obat) error {
	//fmt.Println("mahasiswa")
	err := c.writer.AddObat(obat)
	if err != nil {
		return err
	}

	return nil
}

func (c *obat) ReadObatService(ctx context.Context) (Obats, error) {
	kar, err := c.writer.ReadObat()
	//fmt.Println("mahasiswa", mhs)
	if err != nil {
		return kar, err
	}
	return kar, nil
}

func (c *obat) UpdateObatService(ctx context.Context, kar Obat) error {
	err := c.writer.UpdateObat(kar)
	if err != nil {
		return err
	}
	return nil
}

func (c *obat) ReadObatByNamaService(ctx context.Context, nama_obat string) (Obat, error) {
	kar, err := c.writer.ReadObatByNama(nama_obat)
	//fmt.Println("mahasiswa:", mhs)
	if err != nil {
		return kar, err
	}
	return kar, nil
}
