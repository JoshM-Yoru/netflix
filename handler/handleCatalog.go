package handler

import (
	"fmt"
	"log"
	"netflix/models"
	"netflix/views"
	"netflix/views/components"
	"netflix/views/components/forms"
	"netflix/views/layout"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

func (s *APIServer) HandleGetFullCatalog(c echo.Context) error {
    if len(c.QueryParam("search")) > 0 {
        s.HandleCatalogSearch(c)
    }

	page := c.QueryParam("page")
	pageNum, err := strconv.Atoi(page)
	if err != nil {
		return err
	}

    var catalog []*models.MediaInfo

    if len(CatalogCache) > 0  {
        catalog = CatalogCache
    } else {
        catalog, err = s.store.GetFullCatalog()
        if err != nil {
            log.Fatal(err)
        }
        CatalogCache = catalog
    }

	var pages int

	if len(catalog)%20 == 0 {
		pages = len(catalog) / 20
	} else {
		pages = int(len(catalog)/20) + 1
	}

	// used to make sure that a page that does not exist in the catalog can not be accessed
	if pageNum > pages {
		return render(c, components.BadRequest())
	}

	ctx := views.CatalogContext{
		Catalog:  catalog,
		Page:     pageNum,
		PageSize: 20,
		Pages:    pages,
	}

	if c.Request().Header.Get("HX-Request") == "true" {
		ctx.FullPageReq = false
		return render(c, components.Catalog(ctx))
	} else {
		ctx.FullPageReq = true
		return render(c, layout.CatalogFP(ctx))
	}
}

func (s *APIServer) HandleCatalogSearch(c echo.Context) error {

	catalog, err := s.store.SearchCatalog(c.FormValue("search"))
	if err != nil {
		log.Fatal(err)
	}

    fmt.Println(c.QueryParams())

	var pages int

	if len(catalog)%20 == 0 {
		pages = len(catalog) / 20
	} else {
		pages = int(len(catalog)/20) + 1
	}

	ctx := views.CatalogContext{
		Catalog:  catalog,
		Page:     1,
		PageSize: 20,
		Pages:    pages,
	}

	return render(c, components.Catalog(ctx))
}

func (s *APIServer) HandleNewEntryForm(c echo.Context) error {
	return render(c, forms.AddMediaInfoForm())
}

func (s *APIServer) HandleUpdateEntryForm(c echo.Context) error {
	id := c.QueryParam("id")

	showId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	form, err := s.store.GetMediaEntryById(showId)
	if err != nil {
		return err
	}

	return render(c, forms.MediaInfoForm(form))
}

func (s *APIServer) HandleUpdateCatalog(c echo.Context) error {
	date, err := time.Parse("2006-01-02", c.FormValue("dateAdded"))
	if err != nil {
		return nil
	}

	fmt.Println("test")

	year, err := strconv.Atoi(c.FormValue("releaseYear"))
	if err != nil {
		return err
	}

	media := models.MediaInfo{
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

	err = s.store.CreateMediaEntry(&media)
	if err != nil {
		return err
	}

	catalog, err := s.store.GetFullCatalog()
	if err != nil {
		log.Fatal(err)
	}

	var pages int

	if len(catalog)%20 == 0 {
		pages = len(catalog) / 20
	} else {
		pages = int(len(catalog)/20) + 1
	}

	ctx := views.CatalogContext{
		Catalog:  catalog,
		Page:     1,
		PageSize: 20,
		Pages:    pages,
	}

	return render(c, components.Catalog(ctx))
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
	if err != nil {
		return err
	}

	return render(c, components.MediaInfo(media))
}

func (s *APIServer) HandleDisableCatalogEntry(c echo.Context) error {
	id := c.QueryParam("id")
	numID, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	err = s.store.DisableMediaEntry(numID)
	if err != nil {
		return err
	}

	catalog, err := s.store.GetFullCatalog()
	if err != nil {
		return err
	}

	var pages int

	if len(catalog)%20 == 0 {
		pages = len(catalog) / 20
	} else {
		pages = int(len(catalog)/20) + 1
	}

	ctx := views.CatalogContext{
		Catalog:  catalog,
		Page:     1,
		PageSize: 20,
		Pages:    pages,
	}

	return render(c, components.Catalog(ctx))
}
