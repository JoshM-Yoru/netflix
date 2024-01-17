package handler

import (
	"netflix/models"
	"sync"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

// helper function to render .templ components
func render(c echo.Context, component templ.Component) error {
    return component.Render(c.Request().Context(), c.Response())
}

// cache functionality is currently unused
type SearchCache struct {
    Records []*models.MediaInfo
    SearchedTerm string
    Mu sync.Mutex
}


var CatalogCache []*models.MediaInfo
var CatalogSearchCache *SearchCache
type CacheLookUp map[string]*[]models.MediaInfo

// func (clu *CacheLookUp) CheckCache(searchTerm string) bool {
//     c := *clu
//
//     if v, ok := c[searchTerm]; ok {
//         fmt.Println("true")
//         return true
//     }
//     fmt.Println("false")
//     return false;
// }
//
