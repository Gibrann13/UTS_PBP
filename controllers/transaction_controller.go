package controllers

import (
	m "Modul_2/models"
	n "Modul_2/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func UpdateTransaction(w http.ResponseWriter, r *http.Request, id int, userid int, prodid int, quantity int) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}
	query := "UPDATE transactions SET userid = ?, productid = ?, quantity = ? WHERE ID = ?;"

	var response m.UserResponse
	_, err = db.Exec(query, userid, prodid, quantity, id)
	if err != nil {
		response.Status = 400
		response.Message = "Insert Failed"
		json.NewEncoder(w).Encode(response)
	} else {
		response.Status = 200
		response.Message = "Success"
		json.NewEncoder(w).Encode(response)
	}
}

func DeleteTransaction(w http.ResponseWriter, r *http.Request, userid int, prodid int) {

	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err == nil {
		return
	}

	// query := "DELETE FROM users WHERE Name = ? AND price = ?;"
	query := "DELETE FROM transactions WHERE userid = ? AND productid = ?;"

	var response m.UserResponse
	_, err = db.Exec(query, userid, prodid)

	if err == nil {
		response.Status = 200
		response.Message = "Success"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		response.Status = 400
		response.Message = "Failed"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

func InsertTransaction(w http.ResponseWriter, r *http.Request, userid int, productid int, quantity int) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}
	// cek produkid ada di database gaa
	var count int
	var responses m.TransactionResponse
	err = db.QueryRow("SELECT COUNT(*) FROM products WHERE id = ?", productid).Scan(&count)
	if err != nil {
		if err != nil {
			responses.Status = 400
			responses.Message = "Insert Failed"
		} else {
			responses.Status = 200
			responses.Message = "Success"
		}
		return
	}
	// insert baru kalo tada
	if count == 0 {
		_, err = db.Exec("INSERT INTO products (id, price) VALUES (?, ?)", productid, "")
		if err != nil {
			if err != nil {
				responses.Status = 400
				responses.Message = "Insert Failed"
			} else {
				responses.Status = 200
				responses.Message = "Success"
			}
			return
		}
	}

	query := "Insert INTO transactions (userid, productid, quantity) VALUES (?,?,?)"

	var response n.TransactionResponse
	_, err = db.Exec(query, userid, productid, quantity)

	if err != nil {
		response.Status = 400
		response.Message = "Insert Failed"
	} else {
		response.Status = 200
		response.Message = "Success"
	}
}

func GetAllTransactions(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()
	// m.cobaHitung()

	query := "SELECT * FROM transactions"

	userid := r.URL.Query()["userid"]
	productid := r.URL.Query()["productid"]

	if userid != nil {
		fmt.Println(userid[0])
		query = "WHERE userid='" + userid[0] + "'"
	}
	if productid != nil {
		if productid[0] != "" {
			query += " AND"
		} else {
			query += "WHERE"
		}
		query += " productID='" + productid[0] + "'"
	}

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		//send error response
		return
	}

	var transaction n.Transaction
	var transactions []n.Transaction
	for rows.Next() {
		if err := rows.Scan(&transaction.ID, &transaction.UserID.ID, &transaction.ProductID.ID, &transaction.Quantity); err != nil {
			log.Println(err)
			//send error response
			return
		} else {
			transactions = append(transactions, transaction)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	// var response UsersResponse
	// if len(users) < 5
	var response n.TransactionsResponse
	response.Status = 200
	response.Message = "Success"
	response.Data = transactions
	json.NewEncoder(w).Encode(response)

}
