package utils

import (
	"github.com/go-playground/validator/v10"
)

// Variabel untuk instance validator
var validate *validator.Validate

// InitValidator menginisialisasi validator baru.
func InitValidator() {
	validate = validator.New()
}

// ValidateStruct memvalidasi struct berdasarkan tag validasi.
func ValidateStruct(s interface{}) error {
	return validate.Struct(s)
}

// ParseValidationErrors mengubah error validasi menjadi peta pesan kesalahan.
func ParseValidationErrors(err error) map[string]string {
	errors := make(map[string]string)
	for _, err := range err.(validator.ValidationErrors) {
		field := err.Field()    // Nama field yang gagal divalidasi
		tag := err.Tag()       // Tag validasi yang gagal

		// Menentukan pesan kesalahan berdasarkan tag validasi
		switch tag {
		case "required":
			errors[field] = field + " is required"
		case "min":
			errors[field] = field + " should be at least " + err.Param() + " characters long"
		case "gte":
			errors[field] = field + " should be greater than or equal to " + err.Param()
		default:
			errors[field] = field + " is invalid"
		}
	}
	return errors
}
