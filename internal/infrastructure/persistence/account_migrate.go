package persistence

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations() {
	// ⚙️ 1. 設定資料庫連線
	dsn := "digimon:digimon123@tcp(localhost:3306)/digimon_game?multiStatements=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	// ⚙️ 2. 設定 migrate database driver
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatalf("failed to create migration driver: %v", err)
	}

	// ⚙️ 3. 指定 migrations 路徑
	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"mysql",
		driver,
	)
	if err != nil {
		log.Fatalf("failed to initialize migration: %v", err)
	}

	// ⚙️ 4. 執行 migration
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("migration failed: %v", err)
	}

	fmt.Println("✅ Migration completed successfully.")
}
