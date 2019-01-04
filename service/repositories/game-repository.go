package repositories

import (
    "database/sql"
    _ "github.com/lib/pq"
    "log"
    "strings"
)

type gameRecord struct {
    id            sql.NullInt64
    name          sql.NullString
    boxArt        sql.NullString
    naRelDate     sql.NullString
    euRelDate     sql.NullString
    loosePrice    sql.NullInt64
    licensed      sql.NullBool
    wikipediaUrl  sql.NullString
    publisher     sql.NullString
}

type GameRecord struct {
    Id            int      `json:"id,omitempty"`
    Name          string   `json:"name,omitempty"`
    BoxArt        *string  `json:"boxArt,omitempty"`
    NaRelDate     *string  `json:"naRelDate,omitempty"`
    EuRelDate     *string  `json:"euRelDate,omitempty"`
    LoosePrice    *int     `json:"loosePrice,omitempty"`
    Licensed      *bool    `json:"licensed,omitempty"`
    WikipediaUrl  *string  `json:"wikipediaUrl,omitempty"`
    Publishers    []string `json:"publishers,omitEmpty"`
}

type GameRepository struct {
    database *sql.DB
}

func NewGameRepository() *GameRepository {

    connStr := "postgres://postgres:postgres@localhost/postgres?search_path=game_collector&sslmode=disable"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
       log.Fatal(err)
    }

    return &GameRepository{ database: db }
}

const gamePageQuery string = `select g.*, string_agg(p.publisher, '#,#')
  from game g 
       left join game_x_publisher
         on g.game_id = game_x_publisher.game_id
       inner join publisher p
         on game_x_publisher.publisher_id = p.publisher_id
  group by g.game_id
  offset $1 limit $2`

const filteredGamePageQuery string = `select g.*, string_agg(p.publisher, '#,#')
  from game g 
       left join game_x_publisher
         on g.game_id = game_x_publisher.game_id
       inner join publisher p
         on game_x_publisher.publisher_id = p.publisher_id
  where LOWER(g.game) like $1
  group by g.game_id
  offset $2 limit $3`

func (gr GameRepository) Get(start int, count int, filter string) *PagedDataRep {

    var rows *sql.Rows
    var err error
    if filter != "" {
        filter = filter + "%"
        rows, err = gr.database.Query(filteredGamePageQuery, filter, start, count)
    } else {
        rows, err = gr.database.Query(gamePageQuery, start, count)
    }
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    var games []*GameRecord

    for rows.Next() {

        var row gameRecord
        if err := rows.Scan(&row.id, &row.name, &row.boxArt, &row.naRelDate, &row.euRelDate,
                &row.loosePrice, &row.licensed, &row.wikipediaUrl, &row.publisher); err != nil {
            log.Fatal(err)
        }

        games = append(games, createGameRecord(row))
    }

    var total int
    query := "select count(1) from game"
    if filter != "" {
        query = query + " where LOWER(game) like $1"
        err = gr.database.QueryRow(query, filter).Scan(&total)
    } else {
        err = gr.database.QueryRow(query).Scan(&total)
    }
    if err != nil {
        log.Fatal(err)
    }

    return &PagedDataRep{
        Data: games,
        Start: start,
        Total: total,
    }
}

func createGameRecord(row gameRecord) (gr *GameRecord) {

    id := row.id.Int64
    name := row.name.String

    var boxArt *string = nil
    if row.boxArt.Valid {
        boxArt = &row.boxArt.String
    }

    var naRelDate *string = nil
    if row.naRelDate.Valid {
        naRelDate = possiblyStripTimeStamp(row.naRelDate.String)
    }

    var euRelDate *string = nil
    if row.euRelDate.Valid {
        euRelDate = possiblyStripTimeStamp(row.euRelDate.String)
    }

    var loosePrice *int = nil
    if row.loosePrice.Valid {
        i := int(row.loosePrice.Int64)
        loosePrice = &i
    }

    var licensed *bool = nil
    if row.licensed.Valid {
        licensed = &row.licensed.Bool
    }

    var wikipediaUrl *string = nil
    if row.wikipediaUrl.Valid {
        wikipediaUrl = &row.wikipediaUrl.String
    }

    var publishers []string
    if row.publisher.Valid {
        publishers = strings.Split(row.publisher.String, "#,#")
    }

    return &GameRecord{
        Id: int(id),
        Name: name,
        BoxArt: boxArt,
        NaRelDate: naRelDate,
        EuRelDate: euRelDate,
        LoosePrice: loosePrice,
        Licensed: licensed,
        WikipediaUrl: wikipediaUrl,
        Publishers: publishers,
    }
}

func possiblyStripTimeStamp(iso8601 string) *string {
    if strings.HasSuffix(iso8601, "T00:00:00Z") {
        iso8601 = iso8601[0:10]
    }
    return &iso8601
}
