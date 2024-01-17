package handler

import (
	"netflix/models"
	"sync"
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
	Records      []*models.MediaInfo
	SearchedTerm string
	Mu           sync.Mutex
	Expiration   time.Duration
}

var CatalogCache []*models.MediaInfo
var CatalogSearchCache *SearchCache
