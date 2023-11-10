package main

import (
   "fmt"
   "github.com/labstack/echo/v4"
   "github.com/labstack/echo/v4/middleware"
   "cuaca/routes"
)

func main() {
   e := echo.New()

   // Middleware
   e.Use(middleware.Logger())
   e.Use(middleware.Recover())

   // Routes
   routes.Init(e)

   port := 8080
   e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}
