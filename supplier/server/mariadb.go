package server

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	addSupplier          = `insert into supplier(kode_supplier, nama_supplier, jenis_supplier, createdby, createdon,updateby ,updateon,status)values (?,?,?,?,?,?,?,?,?)`
	selectSupplier       = `select kode_supplier, nama_supplier, jenis_supplier, createdby, createdon,updateby ,updateon,status from supplier Where status =1`
	updateSupplier       = `update supplier set kode_supplier=?, nama_supplier=?, jenis_supplier=?, createdby=?, createdon=?,updateby=? ,updateon=?,status=? where nama_supplier=?`
	selectSupplierByNama = `select kode_supplier, nama_supplier, jenis_supplier, createdby, createdon,updateby ,updateon,status from supplier where nama_supplier=?`
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

func (rw *dbReadWriter) AddSupplier(supplier Supplier) error {
	fmt.Println("insert")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(addSupplier, supplier.KodeSupplier, supplier.NamaSupplier, supplier.JenisSupplier, supplier.CreatedBy, supplier.CreatedOn, supplier.UpdateBy, supplier.UpdateOn, supplier.Status, time.Now())
	//fmt.Println(err)
	if err != nil {
		return err

	}
	return tx.Commit()
}

func (rw *dbReadWriter) ReadSupplierByNama(nama_supplier string) (Supplier, error) {
	fmt.Println("show by nama")
	supplier := Supplier{NamaSupplier: nama_supplier}
	err := rw.db.QueryRow(selectSupplierByNama, nama_supplier).Scan(&supplier.KodeSupplier, &supplier.NamaSupplier, &supplier.JenisSupplier, &supplier.CreatedBy, &supplier.CreatedOn, &supplier.UpdateBy, &supplier.UpdateOn, &supplier.Status)

	if err != nil {
		return Supplier{}, err
	}

	return supplier, nil
}

func (rw *dbReadWriter) ReadSupplier() (Suppliers, error) {
	fmt.Println("show all")
	supplier := Suppliers{}
	rows, _ := rw.db.Query(selectSupplier)
	defer rows.Close()
	for rows.Next() {
		var k Supplier
		err := rows.Scan(&k.KodeSupplier, &k.NamaSupplier, &k.JenisSupplier, &k.CreatedBy, &k.CreatedOn, &k.UpdateBy, &k.UpdateOn, &k.Status)
		if err != nil {
			fmt.Println("error query:", err)
			return supplier, err
		}
		supplier = append(supplier, k)
	}
	//fmt.Println("db nya:", mahasiswa)
	return supplier, nil
}

func (rw *dbReadWriter) UpdateSupplier(kar Supplier) error {
	fmt.Println("update successfuly")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(updateSupplier, kar.KodeSupplier, kar.NamaSupplier, kar.JenisSupplier, kar.CreatedBy, kar.CreatedOn, kar.UpdateBy, kar.UpdateOn, kar.Status, time.Now(), kar.NamaSupplier)

	fmt.Println("name:", kar.NamaSupplier, kar.Status)
	if err != nil {
		return err
	}

	return tx.Commit()
}
