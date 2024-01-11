package handler

import (
	"fmt"
	"log"
	"netflix/views/components"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (s *APIServer) HandleGetFullWatchlist(c echo.Context) error  {
	page := c.QueryParam("page")
	pageNum, err := strconv.Atoi(page)
	if err != nil {
		fmt.Println(err)
		return err
	}

	watchList, err := s.store.GetWatchList()
	if err != nil {
		log.Fatal(err)
	}

	pages := int(len(watchList) / 40)

	return render(c, components.WatchList(watchList, pageNum, 40, pages))
}

func (s *APIServer) HandleUpdateWatchList(c echo.Context) error {
	id := c.QueryParam("id")
	idNum, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

    err = s.store.CreateWatchListEntry(idNum)
	if err != nil {
		return err
	}

	watchList, err := s.store.GetWatchList()
	if err != nil {
		log.Fatal(err)
	}

	pages := int(len(watchList) / 40)

	return render(c, components.WatchList(watchList, 1, 40, pages))
}

func (s *APIServer) HandleUpdateWatchListEntry(c echo.Context) error {
	id := c.QueryParam("id")
	idNum, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

    err = s.store.UpdateWatchList(idNum)
	if err != nil {
		return err
	}

    media, err := s.store.GetWatchListEntryById(idNum)
	if err != nil {
		return err
	}

	return render(c, components.WatchListMediaInfo(*media))
}

func (s *APIServer) HandleDeleteWatchlistEntry(c echo.Context) error {
    return nil
}
