package views

import (
	"netflix/models"
)

type CatalogContext struct {
    Catalog []*models.MediaInfo
    Page int
    PageSize int
    Pages int
    FullPageReq bool
}
type WatchListContext struct {
    Catalog []*models.WatchListInfo
    Page int
    PageSize int
    Pages int
    FullPageReq bool
}
