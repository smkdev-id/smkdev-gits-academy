package controllers

import "gorm.io/gorm"

// Inisialisasi instance DB yang akan digunakan di controller
var db *gorm.DB

// Initialize mengatur koneksi database untuk controller
func Initialize(dbConn *gorm.DB) {
	db = dbConn
}
