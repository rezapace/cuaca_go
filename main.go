// Package main merupakan entry point dari aplikasi Go.
package main

// Mengimport paket-paket yang diperlukan.
import (
   "fmt" // Package untuk formatting dan output.
   "github.com/labstack/echo/v4" // Framework web Echo untuk Go.
   "github.com/labstack/echo/v4/middleware" // Middleware untuk Echo framework.
   "cuaca/routes" // Package routes yang mungkin berisi definisi route untuk aplikasi cuaca.
)

// Fungsi utama (main) yang akan dijalankan saat program dijalankan.
func main() {
   // Membuat objek Echo, yang merupakan instance dari framework web Echo.
   e := echo.New()

   // Middleware
   e.Use(middleware.Logger()) // Menggunakan middleware Logger untuk mencatat log HTTP requests.
   e.Use(middleware.Recover()) // Menggunakan middleware Recover untuk menangani pemulihan dari panic.

   // Routes
   routes.Init(e) // Memanggil fungsi Init dari package routes untuk mendaftarkan route pada objek Echo.

   // Menentukan port yang akan digunakan.
   port := 8080
   // Menggunakan logger dan menjalankan server pada port tertentu.
   e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}
