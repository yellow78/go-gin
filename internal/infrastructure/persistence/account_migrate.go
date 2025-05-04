package persistence

import (
	"fmt"
	"log"
	"time"

	pkgsql "go-gin/pkg/db"

	mysqldriver "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations() {
	// ⚙️ 1. 設定資料庫連線
	sqlconfig := mysqldriver.Config{
		User:                 "digimon",
		Passwd:               "digimon123",
		Net:                  "tcp",
		Addr:                 "localhost:3306",
		DBName:               "digimon_game",
		ParseTime:            true,
		Loc:                  time.Local,
		AllowNativePasswords: true,
	}

	dbManger := pkgsql.NewDBManager()
	dbManger.InitDB(&sqlconfig, "digimon")

	digimonDB, ok := dbManger.GetDB("digimon")
	if !ok {
		log.Fatalf("failed to get database connection")
	}

	defer digimonDB.Close()

	// ⚙️ 2. 設定 migrate database driver
	driver, err := mysql.WithInstance(digimonDB, &mysql.Config{})
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
