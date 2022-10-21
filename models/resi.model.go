package models

import (
	// "database/sql"
	"myapp/config"
	"net/http"
)

type Resi struct {
	Id              int    `json:"id"`
	Nomor           string `json:"nomor"`
	Pengirim        string `json:"pengirim"`
	Penerima        string `json:"penerima"`
	Alamat_Penerima string `json:"alamat_penerima"`
	Product         string `json:"product"`
	Product_Type    string `json:"product_type"`
	Status_Barang   string `json:"status_barang"`
	Estimasi        string `json:"estimasi"`
}

func DataResi() (Response, error) {
	var obj Resi
	var arrobj []Resi
	var res Response

	db := config.ConnectToDB()

	sqlStatement := "SELECT * FROM resi"

	rows, err := db.Query(sqlStatement)

	defer rows.Close()

	if err != nil {
		return res, err
	}
	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Nomor, &obj.Pengirim, &obj.Penerima, &obj.Alamat_Penerima,
			&obj.Product, &obj.Product_Type, &obj.Status_Barang, &obj.Estimasi)
		// err = rows.Scan(&obj.Nomor, &obj.Pengirim, &obj.Penerima, &obj.Alamat_Penerima, &obj.Product, &obj.Product_Type, &obj.Status_Barang, &obj.Estimasi)

		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}
	res.Status = http.StatusOK
	res.Message = "SUKSES"
	res.Data = arrobj

	return res, nil
}

func SimpanData(nomor string, pengirim string, penerima string, alamat_penerima string,
	product string, product_type string, status_barang string, estimasi string) (Response, error) {
	var res Response

	db := config.ConnectToDB()

	sqlStatement := "INSERT INTO resi (nomor, pengirim, penerima, alamat_penerima, product, product_type, status_barang, estimasi) VALUES(?, ?, ?, ?, ?, ?, ?, ?)"

	stmt, err := db.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(nomor, pengirim, penerima, alamat_penerima, product, product_type, status_barang, estimasi)

	if err != nil {
		return res, nil
	}

	getIdLast, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "SUKSES"
	res.Data = map[string]int64{
		"getIdLast": getIdLast,
	}

	return res, nil
}

func UpdateData(id int, nomor string, pengirim string, penerima string, alamat_penerima string,
	product string, product_type string, status_barang string, estimasi string) (Response, error) {

	var res Response

	db := config.ConnectToDB()

	sqlStatement := "UPDATE resi SET nomor = ?, pengirim = ?, penerima = ?, alamat_penerima = ?, product = ?, prodcut_type = ?, status_barang = ?, estimasi = ? WHERE id = ?"
	stmt, err := db.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(nomor, pengirim, penerima, alamat_penerima, product, product_type, status_barang, estimasi, id)

	if err != nil {
		return res, nil
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	// getIdLast, err := result.LastInsertId()
	// if err != nil {
	// 	return res, err
	// }

	res.Status = http.StatusOK
	res.Message = "SUKSES"
	res.Data = map[string]int64{
		"rows": rowsAffected,
		// "getIdLast": getIdLast,
	}

	return res, nil
}
