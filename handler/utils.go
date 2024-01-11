package handler

import (
	"netflix/models"
	"time"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

// helper function to render .templ components
func render(c echo.Context, component templ.Component) error {
    return component.Render(c.Request().Context(), c.Response())
}

// cache functionality is currently unused
type SearchCache struct {
    records []*models.MediaInfo
    searchedTerm string
    expires time.Time
}

var CatalogSearchCache *SearchCache
