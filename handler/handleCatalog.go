package handler

import (
	"fmt"
	"log"
	"netflix/models"
	"netflix/views/components"
	"netflix/views/components/forms"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

func (s *APIServer) HandleGetFullCatalog(c echo.Context) error {
	catalog, err := s.store.GetFullCatalog()
	if err != nil {
		log.Fatal(err)
	}
	page := c.QueryParam("page")
	pageNum, err := strconv.Atoi(page)
	if err != nil {
		fmt.Println(err)
		return err
	}
	pages := int(len(catalog) / 40)

	return render(c, components.Catalog(catalog, pageNum, 40, pages))
}

func (s *APIServer) HandleSearch(c echo.Context) error {
	catalog, err := s.store.SearchCatalog(c.FormValue("search"))
	if err != nil {
		log.Fatal(err)
	}
	pages := int(len(catalog) / 40)

	return render(c, components.Catalog(catalog, 1, 40, pages))
}

func (s *APIServer) HandleNewEntryForm(c echo.Context) error {
	return nil
}

func (s *APIServer) HandleUpdateEntryForm(c echo.Context) error {
	id := c.QueryParam("id")

	showId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	form, err := s.store.GetMediaEntryById(showId)

	return render(c, forms.MediaInfoForm(form))
}

func (s *APIServer) HandleUpdateCatalog(c echo.Context) error {
	return nil
}

func (s *APIServer) HandleUpdateCatalogEntry(c echo.Context) error {
	id, err := strconv.Atoi(c.FormValue("showID"))
	if err != nil {
		return err
	}

	date, err := time.Parse("2006-01-02", c.FormValue("dateAdded"))
	if err != nil {
		return nil
	}

	year, err := strconv.Atoi(c.FormValue("releaseYear"))
	if err != nil {
		return err
	}

	media := models.MediaInfo{
		ID:          id,
		Type:        c.FormValue("type"),
		Title:       c.FormValue("title"),
		Director:    c.FormValue("director"),
		Country:     c.FormValue("country"),
		DateAdded:   date,
		ReleaseYear: year,
		Rating:      c.FormValue("rating"),
		Duration:    c.FormValue("duration"),
		Category:    c.FormValue("category"),
	}

    err = s.store.UpdateMediaInfo(&media)
    if err != nil{
        return err
    }

	return render(c, components.MediaInfo(media))
}

func (s *APIServer) HandleDisableCatalogEntry(c echo.Context) error {
	return nil
}
