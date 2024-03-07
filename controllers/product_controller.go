package controllers

import (
	m "Modul_2/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func UpdateProduct(w http.ResponseWriter, r *http.Request, id int, name string, price int) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}
	query := "UPDATE products SET name = ?, price = ? WHERE id = ?;"

	var response m.ProductResponse
	_, err = db.Exec(query, name, price, id)
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

func DeleteProduct(w http.ResponseWriter, r *http.Request, id int) {

	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}

	//hapus semua id product di table transaction
	query1 := "DELETE FROM transactions WHERE productid= " + strconv.Itoa(id)
	_, err = db.Exec(query1)
	if err != nil {
		http.Error(w, "Failed to execute query", http.StatusInternalServerError)
		return
	}

	//-------------------------------
	//FIX GOSAH DIUBAH WE
	query := "DELETE FROM products WHERE id = ?;"

	var response m.UserResponse
	_, err = db.Exec(query, id)

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

func InsertProduct(w http.ResponseWriter, r *http.Request, name string, price int) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}
	query := "Insert INTO products (name, price) VALUES (?,?)"

	var response m.UserResponse
	_, err = db.Exec(query, name, price)
	if err != nil {
		response.Status = 400
		response.Message = "Insert Failed"
	} else {
		response.Status = 200
		response.Message = "Success"
	}
}

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()
	// m.cobaHitung()

	err := r.ParseForm()
	if err != nil {
		return
	}

	query := "SELECT * FROM products"

	name := r.URL.Query()["name"]
	price := r.URL.Query()["price"]
	if name != nil {
		fmt.Println(name[0])
		query = "WHERE name='" + name[0] + "'"
	}
	if price != nil {
		if name[0] != "" {
			query += " AND"
		} else {
			query += "WHERE"
		}
		query += " price='" + price[0] + "'"
	}

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		//send error response
		return
	}
	var response m.ProductsResponse
	var product m.Product
	var products []m.Product
	for rows.Next() {
		if err := rows.Scan(&product.ID, &product.Name, &product.Price); err != nil {
			log.Println(err)
			//send error response
			response.Status = 400
			response.Message = "Insert Failed"
		} else {
			products = append(products, product)
			response.Status = 200
			response.Message = "Success"
			response.Data = products
			json.NewEncoder(w).Encode(response)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	// var response UsersResponse
	// if len(users) < 5

	response.Status = 200
	response.Message = "Success"
	response.Data = products
	json.NewEncoder(w).Encode(response)

}
