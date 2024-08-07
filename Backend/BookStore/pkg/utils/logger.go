package utils

import (
	"github.com/labstack/gommon/log"
	"os"
)

// Logger adalah instance logger untuk aplikasi.
var Logger = log.New("app")

// InitLogger menginisialisasi konfigurasi logger.
func InitLogger() {
	Logger.SetOutput(os.Stdout) // Menetapkan output logger ke standar output (stdout)
	Logger.SetLevel(log.INFO)   // Mengatur level log ke INFO
}
