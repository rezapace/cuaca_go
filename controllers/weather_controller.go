package controllers

import (
	"net/http"
	"cuaca/utils"
	"github.com/labstack/echo/v4"
)

func WeatherForecastHandler(c echo.Context) error {
	apiKey := "cc0a364fa16d48d0941135837230311" // Ganti dengan kunci API WeatherAPI.com Anda
	location := c.QueryParam("location")
	if location == "" {
		return c.String(http.StatusBadRequest, "Location is required")
	}

	weather, err := utils.FetchWeatherData(apiKey, location)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error fetching weather data")
	}

	// Format data cuaca dan kirim sebagai respons sesuai format yang Anda inginkan.
	return c.JSON(http.StatusOK, weather)
}
