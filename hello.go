package main

import (
    "database/sql"
    "fmt"
    "github.com/gorilla/mux"
    _ "github.com/lib/pq"
    "log"
    "net/http"
)

type compound struct {
    compound_nm   string
    chemotype     sql.NullString
    s10           sql.NullFloat64
    source        sql.NullString
    smiles        sql.NullString
    reference1    sql.NullString
    reference1url sql.NullString
    hidden        bool
    solubility    sql.NullFloat64
}

func main() {

    connStr := "postgres://postgres:postgres@localhost/postgres?sslmode=disable"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err)
    }

    rows, err := db.Query(`select * from compound`)
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    for rows.Next() {

        var c compound
        if err := rows.Scan(&c.compound_nm, &c.chemotype, &c.s10, &c.source, &c.smiles,
            &c.reference1, &c.reference1url, &c.hidden, &c.solubility); err != nil {
            log.Fatal(err)
        }

        var smiles string
        if c.smiles.Valid {
            smiles = c.smiles.String
        } else {
            smiles = "<none>"
        }
        fmt.Printf("%s has smiles string %s\n", c.compound_nm, smiles)
    }

    db.Close()

    router := mux.NewRouter()
    router.HandleFunc("/api/people", GetPeople).Methods("GET")
    router.PathPrefix("/").Handler(http.FileServer(http.Dir("static")))

    log.Println("Listening...")
    http.ListenAndServe(":3000", router)
}

func GetPeople(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello world again")
}
