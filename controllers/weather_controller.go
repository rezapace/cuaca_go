package controllers

import (
	"net/http"
	"cuaca/utils"
	"github.com/labstack/echo/v4"
)

// WeatherForecastHandler adalah fungsi untuk menangani permintaan ramalan cuaca.
func WeatherForecastHandler(c echo.Context) error {
	// Mengambil kunci API WeatherAPI.com. Ganti dengan kunci API Anda sendiri.
	apiKey := "cc0a364fa16d48d0941135837230311"

	// Mengambil nilai parameter 'location' dari permintaan.
	location := c.QueryParam("location")
	if location == "" {
		// Jika lokasi tidak diberikan, kembalikan respons bad request.
		return c.String(http.StatusBadRequest, "Location is required")
	}

	// Mengambil data cuaca menggunakan fungsi FetchWeatherData dari paket utils.
	weather, err := utils.FetchWeatherData(apiKey, location)
	if err != nil {
		// Jika terjadi kesalahan saat mengambil data cuaca, kembalikan respons internal server error.
		return c.String(http.StatusInternalServerError, "Error fetching weather data")
	}

	// Mengatur format data cuaca dan mengirimkannya sebagai respons dalam format JSON.
	return c.JSON(http.StatusOK, weather)
}
