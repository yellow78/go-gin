package main

import "go-gin/internal/infrastructure/migrations"

func main() {
	migrations.RunMigrations()
}
