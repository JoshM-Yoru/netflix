package handler

import (
	"fmt"
	"log"
    "netflix/views"
	"netflix/views/components"
	"netflix/views/layout"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (s *APIServer) HandleGetFullWatchlist(c echo.Context) error {
    if len(c.QueryParam("search")) > 0 {
        s.HandleWatchListSearch(c)
        return nil
    }

	watchList, err := s.store.GetWatchList()
	if err != nil {
		log.Fatal(err)
	}

    var pages int

    if len(watchList) % 20 == 0 {
        pages = len(watchList)/20
    } else {
        pages = int(len(watchList)/20) + 1
    }

	page := c.QueryParam("page")
	pageNum, err := strconv.Atoi(page)
	if err != nil {
        return render(c, layout.WatchListFP(views.WatchListContext{
        Catalog: watchList,
        Page: 1,
        PageSize: 20,
        Pages: pages,
            FullPageReq: true,
        }))
	}

	if pageNum > pages {
		return render(c, components.BadRequest())
	}

    ctx := views.WatchListContext {
        Catalog: watchList,
        Page: pageNum,
        PageSize: 20,
        Pages: pages,
    }

    if (c.Request().Header.Get("HX-Request") == "true"){
        ctx.FullPageReq = false
        return render(c, components.WatchList(ctx))
    } else {
        ctx.FullPageReq = true
        return render(c, layout.WatchListFP(ctx))
    }
}

func (s *APIServer) HandleWatchListSearch(c echo.Context) error {
	searchedTerm := c.FormValue("search")

	watchList, err := s.store.SearchWatchList(searchedTerm)
	if err != nil {
		log.Fatal(err)
	}

    var pages int

    if len(watchList) % 20 == 0 {
        pages = len(watchList)/20
    } else {
        pages = int(len(watchList)/20) + 1
    }

    ctx := views.WatchListContext {
        Catalog: watchList,
        Page: 1,
        PageSize: 20,
		SearchParam: fmt.Sprintf("&search=%s", searchedTerm),
        Pages: pages,
    }

    if (c.Request().Header.Get("HX-Request") == "true"){
        ctx.FullPageReq = false
        return render(c, components.WatchList(ctx))
    } else {
        ctx.FullPageReq = true
        return render(c, layout.WatchListFP(ctx))
    }
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

	media, err := s.store.GetWatchListEntryByFKId(idNum)
	if err != nil {
		return err
	}


	return render(c, components.FavoriteButton(media.Media.ID, media.ID))
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

	return render(c, components.WatchedToggle(*media))
}

func (s *APIServer) HandleDeleteWatchlistEntry(c echo.Context) error {
	id := c.QueryParam("id")
	idNum, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	err = s.store.DeleteWatchListEntry(idNum)
	if err != nil {
		return err
	}

	watchList, err := s.store.GetWatchList()
	if err != nil {
		return err
	}

    var pages int

    if len(watchList) % 20 == 0 {
        pages = len(watchList)/20
    } else {
        pages = int(len(watchList)/20) + 1
    }

    ctx := views.WatchListContext {
        Catalog: watchList,
        Page: 1,
        PageSize: 20,
        Pages: pages,
    }

	return render(c, components.WatchList(ctx))
}
