package main

import (
    "./repositories"
    "encoding/json"
    "fmt"
    "github.com/gorilla/mux"
    "net/http"
    "strconv"
)

type Server struct {
    gameRepository repositories.GameRepository
    userGameRepository repositories.UserGameRepository
}

func NewServer(repository *repositories.GameRepository,
                userGameRepository *repositories.UserGameRepository) *Server {
    return &Server{
        gameRepository: *repository,
        userGameRepository: *userGameRepository,
    }
}

func (s *Server) Run() {

    router := mux.NewRouter()
    router.HandleFunc("/api/games", s.getGames).Methods("GET")
    router.HandleFunc("/api/myGames", s.getUsersGames).Methods("GET")
    router.PathPrefix("/").Handler(http.FileServer(http.Dir("../static")))

    httpServer := &http.Server{
        Addr:    ":3000",
        Handler: router,
    }

    httpServer.ListenAndServe()
}

func (s Server) getGames(w http.ResponseWriter, r *http.Request) {

    startStr := r.URL.Query().Get("start")
    if startStr == "" {
        startStr = "0"
    }

    start, err := strconv.Atoi(startStr)
    if err != nil {
        fmt.Fprintf(w, "Error converting start request param to integer: %s\n", err)
    }

    endStr := r.URL.Query().Get("count")
    if endStr == "" {
        endStr = "10"
    }

    count, err := strconv.Atoi(endStr)
    if err != nil {
        fmt.Fprintf(w, "Error converting count request param to integer: %s\n", err)
    }

    filter := r.URL.Query().Get("filter")

    if err := json.NewEncoder(w).Encode(s.gameRepository.Get(start, count, filter)); err != nil {
        fmt.Fprintf(w, "Error encoding JSON response: %s", err)
    }
}


func (s Server) getUsersGames(w http.ResponseWriter, r *http.Request) {

    startStr := r.URL.Query().Get("start")
    if startStr == "" {
        startStr = "0"
    }

    start, err := strconv.Atoi(startStr)
    if err != nil {
        fmt.Fprintf(w, "Error converting start request param to integer: %s\n", err)
    }

    endStr := r.URL.Query().Get("count")
    if endStr == "" {
        endStr = "10"
    }

    count, err := strconv.Atoi(endStr)
    if err != nil {
        fmt.Fprintf(w, "Error converting count request param to integer: %s\n", err)
    }

    filter := r.URL.Query().Get("filter")

    if err := json.NewEncoder(w).Encode(s.userGameRepository.Get(start, count, filter)); err != nil {
        fmt.Fprintf(w, "Error encoding JSON response: %s", err)
    }
}
