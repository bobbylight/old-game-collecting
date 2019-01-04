package main

import (
    "./repositories"
)

func main() {

    server := NewServer(repositories.NewGameRepository(), repositories.NewUserGameRepository())
    server.Run()
}
