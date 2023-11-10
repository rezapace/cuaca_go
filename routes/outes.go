package routes

import (
   "github.com/labstack/echo/v4"
   "cuaca_go/controllers"
)

func Init(e *echo.Echo) {
   e.GET("/weather", controllers.WeatherForecastHandler)
}
