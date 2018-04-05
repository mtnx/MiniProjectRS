package server

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	addKategoriObat          = `insert into kategori_obat(kode_kategori_obat, nama_kategori_obat, deskripsi, kode_obat, createdby, createdon,updateby ,updateon,status)values (?,?,?,?,?,?,?,?,?)`
	selectKategoriObat       = `select kode_kategori_obat, nama_kategori_obat, deskripsi, kode_obat, createdby, createdon,updateby ,updateon,status from kategori_obat Where status =1`
	updateKategoriObat       = `update kategori_obat set kode_kategori_obat=?, nama_kategori_obat=?, deskripsi=?, kode_obat=?, createdby=?, createdon=?,updateby=? ,updateon=?,status=? where nama_kategori_obat="Obat Bebas"`
	selectKategoriObatByNama = `select kode_kategori_obat, nama_kategori_obat, deskripsi, kode_obat, createdby, createdon,updateby ,updateon,status from kategori_obat where nama_kategori_obat=?`
)

//langkah 4
type dbReadWriter struct {
	db *sql.DB
}

func NewDBReadWriter(url string, schema string, user string, password string) ReadWriter {
	schemaURL := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, url, schema)
	db, err := sql.Open("mysql", schemaURL)
	if err != nil {
		panic(err)
	}
	return &dbReadWriter{db: db}
}

func (rw *dbReadWriter) AddKategoriObat(kategori_obat KategoriObat) error {
	fmt.Println("insert")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(addKategoriObat, kategori_obat.KodeKategoriObat, kategori_obat.NamaKategoriObat, kategori_obat.Deskripsi, kategori_obat.KodeObat, kategori_obat.CreatedBy, kategori_obat.CreatedOn, kategori_obat.UpdateBy, kategori_obat.UpdateOn, kategori_obat.Status, time.Now())
	//fmt.Println(err)
	if err != nil {
		return err

	}
	return tx.Commit()
}

func (rw *dbReadWriter) ReadKategoriObatByNama(nama_kategori_obat string) (KategoriObat, error) {
	fmt.Println("show by nama")
	kategori_obat := KategoriObat{NamaKategoriObat: nama_kategori_obat}
	err := rw.db.QueryRow(selectKategoriObatByNama, nama_kategori_obat).Scan(&kategori_obat.KodeKategoriObat, &kategori_obat.NamaKategoriObat, &kategori_obat.Deskripsi, &kategori_obat.KodeObat, &kategori_obat.CreatedBy, &kategori_obat.CreatedOn, &kategori_obat.UpdateBy, &kategori_obat.UpdateOn, &kategori_obat.Status)

	if err != nil {
		return KategoriObat{}, err
	}

	return kategori_obat, nil
}

func (rw *dbReadWriter) ReadKategoriObat() (KategoriObats, error) {
	fmt.Println("show all")
	kategori_obat := KategoriObats{}
	rows, _ := rw.db.Query(selectKategoriObat)
	defer rows.Close()
	for rows.Next() {
		var k KategoriObat
		err := rows.Scan(&k.KodeKategoriObat, &k.NamaKategoriObat, &k.Deskripsi, &k.KodeObat, &k.CreatedBy, &k.CreatedOn, &k.UpdateBy, &k.UpdateOn, &k.Status)
		if err != nil {
			fmt.Println("error query:", err)
			return kategori_obat, err
		}
		kategori_obat = append(kategori_obat, k)
	}
	//fmt.Println("db nya:", mahasiswa)
	return kategori_obat, nil
}

func (rw *dbReadWriter) UpdateKategoriObat(kar KategoriObat) error {
	fmt.Println("update successfuly")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(updateKategoriObat, kar.KodeKategoriObat, kar.NamaKategoriObat, kar.Deskripsi, kar.KodeObat, kar.CreatedBy, kar.CreatedOn, kar.UpdateBy, kar.UpdateOn, kar.Status, time.Now(), kar.NamaKategoriObat)

	fmt.Println("name:", kar.NamaKategoriObat, kar.Status)
	if err != nil {
		return err
	}

	return tx.Commit()
}
