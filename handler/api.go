package handler

import (
	"netflix/storage"
	"netflix/views/components"
	"netflix/views/layout"

	"github.com/labstack/echo/v4"
)

// struct holds the port number and
type APIServer struct {
	listenAddress string
	store         storage.Storage
}

// function used to fill the APIServer struct
func NewAPIServer(listenAddress string, store storage.Storage) *APIServer {
	return &APIServer{
		listenAddress: listenAddress,
		store:         store,
	}
}

// function used to run the server that includes all necessary endpoints
func (s *APIServer) Run() {
	app := echo.New()
	app.Static("/assets", "assets")
	app.Static("/scripts", "scripts")
	app.Static("/styles", "styles")

	CatalogSearchCache = &SearchCache{
		SearchedTerm: "",
	}

	//home
	app.GET("/", s.HandleHome)

	//gets full catalog
	//also handles searching catalog title for a substring
	app.GET("/catalog", s.HandleGetFullCatalog)

	//gets full watch list
	//also handles searching watch list title for a substring
	app.GET("/watch-list", s.HandleGetFullWatchlist)

	//gets form for new entry
	app.GET("/new-item", s.HandleNewEntryForm)

	//gets form for entry edit
	app.GET("/update-item", s.HandleUpdateEntryForm)

	//posts a new entry to the catalog
	app.POST("/catalog", s.HandleUpdateCatalog)

	//posts a new entry to the watchlist
	app.POST("/watch-list", s.HandleUpdateWatchList)

	//updates a catalog entry at the id
	app.PUT("/catalog", s.HandleUpdateCatalogEntry)

	//updates a watch list entry at the id
	app.PUT("/watch-list", s.HandleUpdateWatchListEntry)

	//used to "delete"/disable an entry from the catalog
	app.PATCH("/catalog", s.HandleDisableCatalogEntry)

	//deletes a watch list entry at id
	app.DELETE("/watch-list", s.HandleDeleteWatchlistEntry)

	app.Start(s.listenAddress)
}

// returns html for the '/' route
func (s *APIServer) HandleHome(c echo.Context) error {
	if c.Request().Header.Get("HX-Request") == "true" {
		return render(c, components.Home())
	} else {
		return render(c, layout.Base())
	}
}

func (s *APIServer) HandleBadPage(c echo.Context) error {
	return render(c, components.BadRequest())
}
