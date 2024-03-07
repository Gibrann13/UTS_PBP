package controllers

import (
	"Modul_2/models"
	m "Modul_2/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func LoginUser(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	email := "gib@gmail.com"
	password := "66666"
	// email := r.FormValue("email")
	// password := r.FormValue("password")

	// Mendapatkan nilai platform dari header
	platform := r.Header.Get("platform")

	fmt.Println("masok1")
	query := "SELECT * FROM users WHERE email = '" + email + "' AND password = '" + password + "'"
	rows, _ := db.Query(query)

	fmt.Println("masukkk")
	var userFound = false
	if rows.Next() == true {
		userFound = true
	}
	fmt.Println("masukk")
	if userFound {
		// Set header platform
		w.Header().Set("platform", platform)

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/plain")
		// Menyesuaikan pesan respons sesuai dengan platform
		fmt.Fprintf(w, "Welcome "+email+"\n")
		fmt.Fprintf(w, "Success login from %s", platform)

		// Contoh untuk mengatur header Content-Type application/json
		w.Header().Set("Content-Type", "application/json")

		var response m.UserResponse
		response.Status = 200

	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Invalid email or password")
	}
	fmt.Println("masok4")
}

func GetDetailUserTransaction(w http.ResponseWriter, r *http.Request, userID int) {
	db := connect()
	defer db.Close()

	var query string
	if userID != 0 {
		query = "SELECT t.id, t.quantity, u.id, u.name, u.age, u.address, p.id, p.name, p.price FROM transactions t JOIN users u ON t.userid = u.id JOIN products p ON t.productid = p.id WHERE u.id = " + strconv.Itoa(userID)
	} else {
		query = "SELECT t.id, t.quantity, u.id, u.name, u.age, u.address, p.id, p.name, p.price FROM transactions t JOIN users u ON t.userid = u.id JOIN products p ON t.productid = p.id"
	}

	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, "Failed to execute query", http.StatusInternalServerError)
		return
	}

	var transactions []m.Transaction

	for rows.Next() {
		var transaction m.Transaction
		var user m.User
		var product m.Product

		err := rows.Scan(&transaction.ID, &transaction.Quantity, &user.ID, &user.Name, &user.Age, &user.Address, &product.ID, &product.Name, &product.Price)
		if err != nil {
			http.Error(w, "Failed to scan row", http.StatusInternalServerError)
			return
		}
		transaction.UserID = user
		transaction.ProductID = product
		transactions = append(transactions, transaction)
	}

	w.Header().Set("Content-Type", "application/json")
	var response = m.TransactionsResponse{}
	response.Status = 200
	response.Message = "Success"
	response.Data = transactions
	json.NewEncoder(w).Encode(response)

}
func UpdateUser(w http.ResponseWriter, r *http.Request, name string, newname string) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}
	query := "UPDATE users SET Name = ? WHERE Name = ?;"

	var response m.UserResponse
	_, err = db.Exec(query, newname, name)
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

func DeleteUser(w http.ResponseWriter, r *http.Request, name string, age int) {

	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}

	// query := "DELETE FROM users WHERE Name = ? AND Age = ?;"
	query := "DELETE FROM users WHERE Name = ?;"

	var response m.UserResponse
	_, err = db.Exec(query, name)

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
func InsertUser(w http.ResponseWriter, r *http.Request, name string, age int, address string) {
	db := connect()
	defer db.Close()

	query := "Insert INTO users (name, age, address) VALUES (?,?,?)"

	var err error
	_, err = db.Exec(query, name, age, address)
	if err != nil {
		log.Println(err)

		http.Error(w, "Gagal Insert", http.StatusInternalServerError)
		return
	} else {
		fmt.Println("User inserted successfully.")
	}
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()
	// m.cobaHitung()

	query := "SELECT * FROM users"

	name := r.URL.Query()["name"]
	age := r.URL.Query()["age"]
	if name != nil {
		fmt.Println(name[0])
		query = "WHERE name='" + name[0] + "'"
	}
	if age != nil {
		if name[0] != "" {
			query += " AND"
		} else {
			query += "WHERE"
		}
		query += " age='" + age[0] + "'"
	}

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		//send error response
		return
	}

	var user m.User
	var users []m.User
	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Name, &user.Age, &user.Address); err != nil {
			log.Println(err)
			//send error response
			return
		} else {
			users = append(users, user)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	// var response UsersResponse
	// if len(users) < 5
	var response m.UsersResponse
	response.Status = 200
	response.Message = "Success"
	response.Data = users
	json.NewEncoder(w).Encode(response)

}

// GORM Used
// GORM Insert
func InsertUserGORM(w http.ResponseWriter, r *http.Request) {
	// Parse input dari request
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "failed to parse form data", http.StatusBadRequest)
		return
	}
	name := r.Form.Get("name")
	age, _ := strconv.Atoi(r.Form.Get("age"))
	address := r.Form.Get("address")
	email := r.Form.Get("email")
	password := r.Form.Get("password")
	if email == "" || password == "" {
		http.Error(w, "please provide both email and password", http.StatusBadRequest)
		return
	}

	db := connectgorm()

	user := models.User{Name: name, Age: age, Address: address, Email: email, Password: password}

	result := db.Create(&user)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	var response = m.UserResponse{}
	response.Status = http.StatusOK
	response.Message = "Success"
	response.Data = user
	json.NewEncoder(w).Encode(response)
}

// GORM UpdateUser
func UpdateUserGORM(w http.ResponseWriter, r *http.Request) {
	// Parse input dari request
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "failed to parse form data", http.StatusBadRequest)
		return
	}

	// Ambil nama lama dan nama baru dari input form
	oldName := r.Form.Get("oldName")
	newName := r.Form.Get("newName")
	address := r.Form.Get("newAddress")
	if oldName == "" || newName == "" {
		http.Error(w, "please provide both oldName and newName", http.StatusBadRequest)
		return
	}

	// Koneksi ke database
	db := connectgorm()

	var user m.User
	result := db.Model(&user).Select("name", "age", "address").Where("name = ?", oldName).Updates(map[string]interface{}{"name": newName, "address": address})
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	var response = m.UserResponse{}
	response.Status = http.StatusOK
	response.Message = "Success"
	response.Data = user
	json.NewEncoder(w).Encode(response)
}

// GORM Delete User
func DeleteUserGORM(w http.ResponseWriter, r *http.Request) {
	// Parse input dari request
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "failed to parse form data", http.StatusBadRequest)
		return
	}

	// Ambil nama  dari input form
	name := r.Form.Get("name")

	if name == "" {
		http.Error(w, "please provide name to delete", http.StatusBadRequest)
		return
	}

	var user m.User
	// Koneksi ke database
	db := connectgorm()

	result := db.Where("name = ?", name).Delete(&user)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	var response = m.UserResponse{}
	response.Status = http.StatusOK
	response.Message = "Success Deleting the Data"
	response.Data = user
	json.NewEncoder(w).Encode(response)
}

// Select User GORMM
func SelectUserGORM(w http.ResponseWriter, r *http.Request) {
	// Parse input dari request
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "failed to parse form data", http.StatusBadRequest)
		return
	}
	var users []m.User
	// Koneksi ke database
	db := connectgorm()

	result := db.Find(&users)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	var response = m.UserResponse{}
	response.Status = http.StatusOK
	response.Message = "Success Get the users"
	//response.Data = users
	json.NewEncoder(w).Encode(users)
}
func UpdateRawGORM(w http.ResponseWriter, r *http.Request) {
	// Parse input dari request
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "failed to parse form data", http.StatusBadRequest)
		return
	}

	minAge := r.FormValue("min_age")
	maxAge := r.FormValue("max_age")
	newAddress := r.FormValue("new_address")

	if minAge == "" || maxAge == "" || newAddress == "" {
		http.Error(w, "please provide min_age and max_age and new_address", http.StatusBadRequest)
		return
	}

	var users []m.User
	db := connectgorm()
	// Raw SQL

	// Query SQL mentah
	rows, err := db.Raw("SELECT name, age, address FROM users WHERE age BETWEEN ? AND ?", minAge, maxAge).Rows()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Iterate over rows
	for rows.Next() {
		var user m.User
		// Scan baris ke variabel yang sesuai
		if err := db.ScanRows(rows, &user); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Lakukan pembaruan alamat untuk setiap pengguna
		user.Address = newAddress

		// Lakukan pembaruan menggunakan GORM
		if err := db.Model(&m.User{}).Where("age BETWEEN ? AND ?", minAge, maxAge).Update("address", newAddress).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		users = append(users, user)
	}

	w.Header().Set("Content-Type", "application/json")
	var response = m.UserResponse{}
	response.Status = http.StatusOK
	response.Message = "Success Update the Address for >min_age and <max_age"
	//response.Data = users
	json.NewEncoder(w).Encode(users)
}
