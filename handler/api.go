package handler

import (
	"netflix/storage"
	"netflix/views/layout"

	"github.com/labstack/echo/v4"
)

type APIServer struct {
	listenAddress string
	store         storage.Storage
}

func NewAPIServer(listenAddress string, store storage.Storage) *APIServer {
	return &APIServer{
		listenAddress: listenAddress,
		store:         store,
	}
}

func (s *APIServer) Run() {
	app := echo.New()
    app.Static("/", "assests")

	//home
	app.GET("/", s.HandleHome)

	//gets full catalog
	app.GET("/catalog", s.HandleGetFullCatalog)

	//gets full watch list
	app.GET("/watch-list", s.HandleGetFullWatchlist)

	//searches catalog for a substring
	app.GET("/search", s.HandleSearch)

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

func (s *APIServer) HandleHome(c echo.Context) error {
	return render(c, layout.Base())
}
