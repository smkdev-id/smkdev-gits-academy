package models

// Todo mewakili sebuah tugas dengan ID, deskripsi, dan status penyelesaian.
type Todo struct {
	ID        int    `json:"id"`        // ID adalah pengenal unik untuk item Todo.
	Todo      string `json:"todo"`      // Todo adalah deskripsi dari tugas yang harus diselesaikan.
	Completed bool   `json:"completed"` // Completed menunjukkan apakah tugas telah diselesaikan.
}
