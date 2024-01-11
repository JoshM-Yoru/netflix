package storage

import (
	"database/sql"
	"fmt"
	"netflix/models"
	"strings"

	_ "github.com/lib/pq"
)

// database function interface
type Storage interface {
	GetFullCatalog() ([]*models.MediaInfo, error)
	SearchCatalog(string) ([]*models.MediaInfo, error)
	SearchWatchList(string) ([]*models.WatchListInfo, error)
	GetMediaEntryById(int) (*models.MediaInfo, error)
	GetWatchListEntryById(int) (*models.WatchListInfo, error)
	UpdateMediaInfo(*models.MediaInfo) error
	CreateMediaEntry(*models.MediaInfo) error
	DisableMediaEntry(int) error
	GetWatchList() ([]*models.WatchListInfo, error)
	CreateWatchListEntry(int) error
	UpdateWatchList(int) error
	DeleteWatchListEntry(int) error
}

// struct for the database connection
type PostgresStore struct {
	db *sql.DB
}

// function used to connect to the local database
// returns the pointer to the database connection
func NewPostgresStore() (*PostgresStore, error) {
	conStr := "user=postgres dbname=netflix sslmode=disable"
	db, err := sql.Open("postgres", conStr)
	if err != nil {
		return nil, err
	}

	// used to verify the connection is active
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{
		db: db,
	}, nil
}

// database init function used to create the watched_media table if it does not exist
func (s *PostgresStore) Init() error {
	watchedTable := s.createWatchedMediaTable()
	if watchedTable != nil {
		return watchedTable
	}
	return nil
}

// function to create watched_media table
func (s *PostgresStore) createWatchedMediaTable() error {
	query := `create table if not exists watched_media(
        pk_id serial primary key,
        fk_id integer references catalog(pk_id),
        watched boolean
    )`

	_, err := s.db.Query(query)

	return err
}

func (s *PostgresStore) GetFullCatalog() ([]*models.MediaInfo, error) {

	rows, err := s.db.Query(`
        select pk_id, type, title, director, country, date_added, release_year, rating, duration, listed_in 
        from catalog 
        where is_deleted = false 
        order by title asc
        `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	catalog := []*models.MediaInfo{}
	for rows.Next() {
		mediaInfo, err := scanIntoMediaInfo(rows)
		if err != nil {
			return nil, err
		}
		catalog = append(catalog, mediaInfo)
	}

	return catalog, nil
}

func (s *PostgresStore) SearchCatalog(title string) ([]*models.MediaInfo, error) {
	searchPattern := fmt.Sprintf("%%%s%%", strings.ToLower(title))

	rows, err := s.db.Query(`
        select pk_id, type, title, director, country, date_added, release_year, rating, duration, listed_in 
        from catalog 
        where is_deleted = false and title ilike $1 
        order by title asc`, searchPattern)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	catalog := []*models.MediaInfo{}
	for rows.Next() {
		mediaInfo, err := scanIntoMediaInfo(rows)
		if err != nil {
			return nil, err
		}
		catalog = append(catalog, mediaInfo)
	}

	return catalog, nil
}

func (s *PostgresStore) GetWatchList() ([]*models.WatchListInfo, error) {

	rows, err := s.db.Query(`
        select c.type, c.title, c.director, c.country, c.date_added, c.release_year, c.rating, c.duration, c.listed_in, wm.pk_id, wm.watched 
        from catalog c 
        right join watched_media wm 
        on c.pk_id = wm.fk_id 
        order by c.title asc
        `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	watchList := []*models.WatchListInfo{}
	for rows.Next() {
		mediaInfo, err := scanIntoWatchListInfo(rows)
		if err != nil {
			return nil, err
		}
		watchList = append(watchList, mediaInfo)
	}

	return watchList, nil
}

func (s *PostgresStore) SearchWatchList(title string) ([]*models.WatchListInfo, error) {
	searchPattern := fmt.Sprintf("%%%s%%", strings.ToLower(title))

	rows, err := s.db.Query(`
        select c.type, c.title, c.director, c.country, c.date_added, c.release_year, c.rating, c.duration, c.listed_in, wm.pk_id, wm.watched 
        from catalog c 
        right join watched_media wm 
        on c.pk_id = wm.fk_id 
        where c.title ilike $1 and c.pk_id = wm.fk_id 
        order by c.title asc`, searchPattern)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	watchList := []*models.WatchListInfo{}
	for rows.Next() {
		mediaInfo, err := scanIntoWatchListInfo(rows)
		if err != nil {
			return nil, err
		}
		watchList = append(watchList, mediaInfo)
	}

	return watchList, nil
}

func (s *PostgresStore) GetMediaEntryById(id int) (*models.MediaInfo, error) {

	rows, err := s.db.Query(`
        select pk_id, type, title, director, country, date_added, release_year, rating, duration, listed_in 
        from catalog 
        where is_deleted = false and pk_id = $1`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	rows.Next()

	mediaInfo, err := scanIntoMediaInfo(rows)
	if err != nil {
		return nil, err
	}

	return mediaInfo, nil
}

func (s *PostgresStore) GetWatchListEntryById(id int) (*models.WatchListInfo, error) {

	rows, err := s.db.Query(`
        select c.type, c.title, c.director, c.country, c.date_added, c.release_year, c.rating, c.duration, c.listed_in, wm.pk_id, wm.watched 
        from catalog c 
        right join watched_media wm 
        on c.pk_id = wm.fk_id 
        where wm.pk_id = $1`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	rows.Next()

	mediaInfo, err := scanIntoWatchListInfo(rows)
	if err != nil {
		return nil, err
	}

	return mediaInfo, nil
}

func (s *PostgresStore) CreateMediaEntry(info *models.MediaInfo) error {
	query := `insert into catalog (type, title, director, country, date_added, release_year, rating, duration, listed_in, is_deleted)
                values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

	_, err := s.db.Query(
		query,
		info.Type,
		info.Title,
		info.Director,
		info.Country,
		info.DateAdded.Format("2006-01-02"),
		info.ReleaseYear,
		info.Rating,
		info.Duration,
		info.Category,
		false,
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *PostgresStore) CreateWatchListEntry(id int) error {
	query := `insert into watched_media (fk_id, watched)
                values ($1, $2)`

	_, err := s.db.Query(
		query,
		id,
		false,
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *PostgresStore) DisableMediaEntry(id int) error {
	query := `update catalog 
              set is_deleted = true 
              where pk_id = $1`

	_, err := s.db.Query(query, id)
	if err != nil {
		return err
	}

	return nil
}

// unused
func (s *PostgresStore) DeleteMediaEntry(id int) error {
	query := `delete from catalog 
              where show_id = $1`

	_, err := s.db.Query(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *PostgresStore) DeleteWatchListEntry(id int) error {
	query := `delete from watched_media 
              where pk_id = $1`

	_, err := s.db.Query(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *PostgresStore) UpdateWatchList(id int) error {
	query := `update watched_media 
              set watched = case 
                      when watched = true then false 
                      else true 
                  end 
              where pk_id = $1`

	_, err := s.db.Query(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *PostgresStore) UpdateMediaInfo(info *models.MediaInfo) error {
	query := `update catalog 
              set type = $1, title = $2, director = $3, country = $4, date_added = $5, release_year = $6, rating = $7, duration = $8, listed_in = $9 
              where pk_id = $10`

	_, err := s.db.Query(
		query,
		info.Type,
		info.Title,
		info.Director,
		info.Country,
		info.DateAdded.Format("2006-01-02"),
		info.ReleaseYear,
		info.Rating,
		info.Duration,
		info.Category,
		info.ID,
	)
	if err != nil {
		return err
	}

	return nil
}

// helper function used to get the data from the sql row and map it to the MediaInfo stuct
func scanIntoMediaInfo(rows *sql.Rows) (*models.MediaInfo, error) {
	media := new(models.MediaInfo)
	err := rows.Scan(
		&media.ID,
		&media.Type,
		&media.Title,
		&media.Director,
		&media.Country,
		&media.DateAdded,
		&media.ReleaseYear,
		&media.Rating,
		&media.Duration,
		&media.Category,
	)

	return media, err
}

// helper function used to get the data from the sql row and map it to the WatchListInfo stuct
func scanIntoWatchListInfo(rows *sql.Rows) (*models.WatchListInfo, error) {
	media := new(models.WatchListInfo)
	err := rows.Scan(
		&media.Media.Type,
		&media.Media.Title,
		&media.Media.Director,
		&media.Media.Country,
		&media.Media.DateAdded,
		&media.Media.ReleaseYear,
		&media.Media.Rating,
		&media.Media.Duration,
		&media.Media.Category,
		&media.ID,
		&media.Watched,
	)

	return media, err
}
