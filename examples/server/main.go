package main

import (
	"database/sql"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/JonF12/TemplComponentLib/examples"
	"github.com/JonF12/TemplComponentLib/examples/models"
	"github.com/JonF12/TemplComponentLib/src/dropzone"
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func render(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response().Writer)
}

func main() {
	e := echo.New()

	fileServer := http.FileServer(http.FS(examples.Files))
	e.Logger.SetLevel(log.INFO)
	e.GET("/assets/*", echo.WrapHandler(fileServer))
	e.GET("/", renderMain)
	e.GET("/form", renderForm)
	e.GET("/article", renderArticle)
	e.POST("/addcustomer", renderAddCustomer)
	e.POST("/dropzone-upload", dropzoneUpload)
	e.DELETE("/dropzone-delete/:id", dropzoneDelete)
	e.Logger.Fatal(e.Start(":3000"))
}

func renderAddCustomer(c echo.Context) error {
	// Parse form values and convert types
	paying, err := strconv.ParseBool(c.FormValue("customer-ispaying"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid paying status")
	}

	taxIdentifier := strings.TrimSpace(c.FormValue("customer-taxidentifier"))
	if taxIdentifier == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Tax identifier is required")
	}

	countryCode := strings.TrimSpace(c.FormValue("customer-countrycode"))
	if taxIdentifier == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Country code is required")
	}

	// Parse date
	payingDate, err := time.Parse("2006-01-02", c.FormValue("customer-payingdate"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid date format")
	}

	// Parse and validate image ID
	var imageID sql.NullInt64
	if rawImageID := c.FormValue("customer-image-id"); rawImageID != "" {
		id, err := strconv.ParseInt(rawImageID, 10, 64)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid image ID")
		}
		imageID = sql.NullInt64{Int64: id, Valid: true}
	}

	c.Logger().Info(imageID)
	// Create customer object
	customer := &models.Customer{
		Paying:        paying,
		CountryCode:   countryCode,
		TaxIdentifier: taxIdentifier,
		PayingDate:    payingDate,
		FileId:        imageID,
		FileName:      "chicken.png",
	}
	c.Logger().Info(customer.Paying)
	// Render the page with the customer data
	return render(c, examples.FormsPage(customer))
}

func renderMain(c echo.Context) error {
	return render(c, examples.Page())
}

func renderArticle(c echo.Context) error {
	return render(c, examples.Article())
}

func renderForm(c echo.Context) error {
	//simulate getting data here
	customer := &models.Customer{
		PayingDate:    time.Now(),
		Paying:        true,
		TaxIdentifier: "AAABBBCCC",
		CountryCode:   "ca",
		FileId:        sql.NullInt64{Int64: 124, Valid: true},
		FileName:      "chicken.png",
	}
	return render(c, examples.FormsPage(customer))
}

func dropzoneUpload(c echo.Context) error {
	headers := c.Request().Header
	props := &dropzone.Props{
		Label:     headers.Get("Dropzone-Label"),
		Id:        headers.Get("Dropzone-Id"),
		Required:  headers.Get("Dropzone-Required") == "true",
		Accept:    headers.Get("Dropzone-Accept"),
		Name:      headers.Get("Dropzone-Name"),
		UploadUrl: headers.Get("Dropzone-Uploadurl"),
		DeleteUrl: headers.Get("Dropzone-Deleteurl"),
		FileName:  headers.Get("Dropzone-Filename"),
		FileId:    headers.Get("Dropzone-FileId"),
	}
	file, err := c.FormFile(props.Name)
	if err != nil {
		props.ValidationText = err.Error()
		return dropzone.Render(props).Render(c.Request().Context(), c.Response().Writer)
	}
	//simulating addition here
	props.FileName = file.Filename
	props.Disabled = true
	props.FileId = "123"

	return render(c, dropzone.Render(props))
}

func dropzoneDelete(c echo.Context) error {
	headers := c.Request().Header
	props := &dropzone.Props{
		Label:     headers.Get("Dropzone-Label"),
		Id:        headers.Get("Dropzone-Id"),
		Required:  headers.Get("Dropzone-Required") == "true",
		Accept:    headers.Get("Dropzone-Accept"),
		Name:      headers.Get("Dropzone-Name"),
		UploadUrl: headers.Get("Dropzone-Uploadurl"),
		DeleteUrl: headers.Get("Dropzone-Deleteurl"),
		FileName:  headers.Get("Dropzone-Filename"),
		FileId:    headers.Get("Dropzone-FileId"),
	}

	//simulate deletion set this to "null"
	props.FileId = ""
	props.FileName = ""
	return render(c, dropzone.Render(props))
}
