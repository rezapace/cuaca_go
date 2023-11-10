package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"github.com/rezapace/cuaca_go/models/weather_model"
)


func FetchWeatherData(apiKey, location string) (*models.WeatherResponse, error) {
   url := fmt.Sprintf("https://api.weatherapi.com/v1/forecast.json?key=%s&q=%s&days=1&aqi=no&alerts=no", apiKey, location)
   response, err := http.Get(url)
   if err != nil {
      return nil, err
   }
   defer response.Body.Close()

   data, err := ioutil.ReadAll(response.Body)
   if err != nil {
      return nil, err
   }

   var weather models.WeatherResponse
   if err := json.Unmarshal(data, &weather); err != nil {
      return nil, err
   }

   return &weather, nil
}
