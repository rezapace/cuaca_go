// Package routes mendefinisikan file ini sebagai bagian dari paket (package) "routes".
package routes

// Mengimport paket eksternal yang dibutuhkan.
import (
   "github.com/labstack/echo/v4" // Framework web Echo untuk Go
   "cuaca/controllers" // Package controllers yang mungkin berisi fungsi-fungsi terkait logika bisnis cuaca.
)

// Init adalah fungsi untuk menginisialisasi route pada objek *echo.Echo.
func Init(e *echo.Echo) {
   // Mendaftarkan route HTTP GET pada path "/weather" ke fungsi WeatherForecastHandler dari package controllers.
   e.GET("/weather", controllers.WeatherForecastHandler)
}
