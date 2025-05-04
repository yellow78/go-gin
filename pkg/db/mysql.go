package mysql

import (
	"database/sql"
	"log"
	"sync"

	"github.com/go-sql-driver/mysql"
)

type DBManager struct {
	dbMap   map[string]*sql.DB
	syncMap map[string]*sync.Once
	mu      sync.Mutex
}

func NewDBManager() *DBManager {
	return &DBManager{
		dbMap:   make(map[string]*sql.DB),
		syncMap: make(map[string]*sync.Once),
	}
}

func (m *DBManager) InitDB(cfg *mysql.Config, name string) {
	m.mu.Lock()
	if _, ok := m.syncMap[name]; !ok {
		m.syncMap[name] = &sync.Once{}
	}
	once := m.syncMap[name]
	m.mu.Unlock()

	once.Do(func() {
		connector, err := mysql.NewConnector(cfg)
		if err != nil {
			log.Fatalf("failed to create connector: %v", err)
		}

		db := sql.OpenDB(connector)
		if err := db.Ping(); err != nil {
			log.Fatalf("failed to ping database: %v", err)
		}

		m.mu.Lock()
		m.dbMap[name] = db
		m.mu.Unlock()

		log.Printf("Database %s connected (via OpenDB)", name)
	})
}

func (d *DBManager) GetDB(name string) (*sql.DB, bool) {
	d.mu.Lock()
	defer d.mu.Unlock()
	db, ok := d.dbMap[name]
	return db, ok
}

func (d *DBManager) CloseDB(name string) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if db, ok := d.dbMap[name]; ok {
		if err := db.Close(); err != nil {
			log.Printf("failed to close database %s: %v", name, err)
		} else {
			log.Printf("Database %s closed", name)
		}
		delete(d.dbMap, name)
		delete(d.syncMap, name)
	}
}

func (d *DBManager) CloseAll() {
	d.mu.Lock()
	defer d.mu.Unlock()
	for name, db := range d.dbMap {
		if err := db.Close(); err != nil {
			log.Printf("failed to close database %s: %v", name, err)
		} else {
			log.Printf("Database %s closed", name)
		}
	}

	d.dbMap = make(map[string]*sql.DB)
	d.syncMap = make(map[string]*sync.Once)
}
