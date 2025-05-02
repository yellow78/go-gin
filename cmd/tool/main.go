package main

import "go-gin/internal/infrastructure/persistence"

func main() {
	persistence.RunMigrations()
}
