// Package utils mendefinisikan file ini sebagai bagian dari paket (package) "utils".
package utils

// Mengimport paket-paket eksternal yang dibutuhkan.
import (
	"encoding/json" // Package untuk encoding dan decoding JSON.
	"fmt"            // Package untuk formatting dan output.
	"io/ioutil"      // Package untuk membaca dan menulis file.
	"net/http"       // Package untuk membuat HTTP requests.
	"cuaca/models"   // Package models yang mungkin berisi definisi struktur data terkait cuaca.
)

// FetchWeatherData adalah fungsi untuk mengambil data cuaca dari API eksternal.
func FetchWeatherData(apiKey, location string) (*models.WeatherResponse, error) {
	// Membentuk URL dengan menggunakan API key dan lokasi untuk mengambil data cuaca.
	url := fmt.Sprintf("https://api.weatherapi.com/v1/forecast.json?key=%s&q=%s&days=1&aqi=no&alerts=no", apiKey, location)
	
	// Mengirimkan HTTP GET request ke URL.
	response, err := http.Get(url)
	if err != nil {
		return nil, err // Mengembalikan nilai nil dan error jika terjadi kesalahan pada request.
	}
	defer response.Body.Close() // Menutup body response setelah selesai.

	// Membaca data dari body response.
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err // Mengembalikan nilai nil dan error jika terjadi kesalahan pada pembacaan data.
	}

	// Membuat variabel weather untuk menampung hasil parsing JSON.
	var weather models.WeatherResponse
	// Meng-unmarshal data JSON ke dalam variabel weather.
	if err := json.Unmarshal(data, &weather); err != nil {
		return nil, err // Mengembalikan nilai nil dan error jika terjadi kesalahan pada proses unmarshalling.
	}

	// Mengembalikan data cuaca dan nilai error yang nil (jika tidak ada error).
	return &weather, nil
}
