package main

import (
	"github.com/JonF12/TemplComponentLib/examples"
	"net/http"
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func render(c echo.Context, component templ.Component) error {
 return component.Render(c.Request().Context(), c.Response().Writer)
}

func main() {
  e := echo.New()

  fileServer := http.FileServer(http.FS(examples.Files))
  e.GET("/assets/*", echo.WrapHandler(fileServer))
  e.GET("/", renderMain)

  e.Logger.Fatal(e.Start(":3000"))
}

func renderMain(c echo.Context) error {
  render(c, examples.Page())
  return nil
}
