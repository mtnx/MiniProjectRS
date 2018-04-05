package server

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	addObat          = `insert into obat(kode_obat, nama_obat, tanggal_kadaluwarsa, harga, createdby, createdon,updateby ,updateon, status)values (?,?,?,?,?,?,?,?,?)`
	selectObat       = `select kode_obat, nama_obat, tanggal_kadaluwarsa, harga, createdby, createdon,updateby ,updateon,status from obat Where status =1`
	updateObat       = `update obat set nama_obat=?, tanggal_kadaluwarsa=?, harga=?, createdby=?, createdon=?,updateby=? ,updateon=?,status=? where kode_obat=?`
	selectObatByNama = `select kode_obat, nama_obat, tanggal_kadaluwarsa, harga, createdby, createdon,updateby ,updateon,status from obat where nama_obat=?`
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

func (rw *dbReadWriter) AddObat(obat Obat) error {
	fmt.Println("insert")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(addObat, obat.KodeObat, obat.NamaObat, obat.TanggalKadaluwarsa, obat.Harga, obat.CreatedBy, obat.CreatedOn, obat.UpdateBy, obat.UpdateOn, obat.Status)
	//fmt.Println(err)
	if err != nil {
		return err

	}
	return tx.Commit()
}

func (rw *dbReadWriter) ReadObatByNama(nama_obat string) (Obat, error) {
	fmt.Println("show by nama")
	obat := Obat{NamaObat: nama_obat}
	err := rw.db.QueryRow(selectObatByNama, nama_obat).Scan(&obat.KodeObat, &obat.NamaObat, &obat.TanggalKadaluwarsa, &obat.Harga, &obat.CreatedBy, &obat.CreatedOn, &obat.UpdateBy, &obat.UpdateOn, &obat.Status)

	if err != nil {
		return Obat{}, err
	}

	return obat, nil
}

func (rw *dbReadWriter) ReadObat() (Obats, error) {
	fmt.Println("show all")
	obat := Obats{}
	rows, _ := rw.db.Query(selectObat)
	defer rows.Close()
	for rows.Next() {
		var k Obat
		err := rows.Scan(&k.KodeObat, &k.NamaObat, &k.TanggalKadaluwarsa, &k.Harga, &k.CreatedBy, &k.CreatedOn, &k.UpdateBy, &k.UpdateOn, &k.Status)
		if err != nil {
			fmt.Println("error query:", err)
			return obat, err
		}
		obat = append(obat, k)
	}
	//fmt.Println("db nya:", mahasiswa)
	return obat, nil
}

func (rw *dbReadWriter) UpdateObat(kar Obat) error {
	fmt.Println("update successfuly")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(updateObat, kar.NamaObat, kar.TanggalKadaluwarsa, kar.Harga, kar.CreatedBy, kar.CreatedOn, kar.UpdateBy, kar.UpdateOn, kar.Status, kar.KodeObat)

	fmt.Println("name:", kar.NamaObat, kar.Status)
	if err != nil {
		return err
	}

	return tx.Commit()
}
