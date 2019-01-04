package repositories

import (
    "database/sql"
    _ "github.com/lib/pq"
    "log"
)

type UserGameRepository struct {
    database *sql.DB
}

func NewUserGameRepository() *UserGameRepository {

    connStr := "postgres://postgres:postgres@localhost/postgres?search_path=game_collector&sslmode=disable"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
       log.Fatal(err)
    }

    return &UserGameRepository{ database: db }
}

func (gr UserGameRepository) Get(start int, count int, filter string) *PagedDataRep {

    query := `select g.*, string_agg(p.publisher, '#,#')
      from user_x_game u
        join game g
          on u.game_id = g.game_id
        left join game_x_publisher
          on g.game_id = game_x_publisher.game_id
        inner join publisher p
          on game_x_publisher.publisher_id = p.publisher_id
        where u.user_id = 1
        group by g.game_id
        offset $1 limit $2`

    filteredQuery := `select g.*, string_agg(p.publisher, '#,#')
      from user_x_game u
        join game g
            on u.game_id = g.game_id
        left join game_x_publisher
          on g.game_id = game_x_publisher.game_id
        inner join publisher p
          on game_x_publisher.publisher_id = p.publisher_id
        where u.user_id = 1 and LOWER(g.game) like $1
        group by g.game_id
        offset $2 limit $3`

    var rows *sql.Rows
    var err error
    if filter != "" {
        filter = filter + "%"
        rows, err = gr.database.Query(filteredQuery, filter, start, count)
    } else {
        rows, err = gr.database.Query(query, start, count)
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
    query = "select count(1) from user_x_game u where u.user_id = 1"
    if filter != "" {
        query = query + " and LOWER(game) like $1"
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
