package mysql

import (
	"database/sql"
	"errors"
	"sync"

	"github.com/go-sql-driver/mysql"
)

type DBManager struct {
	dbMap   map[string]*sql.DB
	onceMap sync.Map // map[string]*sync.Once
	mu      sync.RWMutex
}

func NewDBManager() *DBManager {
	return &DBManager{
		dbMap: make(map[string]*sql.DB),
	}
}

func (m *DBManager) InitDB(cfg *mysql.Config, name string) error {
	val, _ := m.onceMap.LoadOrStore(name, &sync.Once{})
	once := val.(*sync.Once)

	var initErr error
	once.Do(func() {
		connector, err := mysql.NewConnector(cfg)
		if err != nil {
			initErr = err
			return
		}

		db := sql.OpenDB(connector)
		if err := db.Ping(); err != nil {
			initErr = err
			return
		}

		m.mu.Lock()
		defer m.mu.Unlock()
		m.dbMap[name] = db
	})

	return initErr
}

func (d *DBManager) GetDB(name string) (*sql.DB, bool) {
	d.mu.Lock()
	defer d.mu.Unlock()
	db, ok := d.dbMap[name]
	return db, ok
}

func (d *DBManager) CloseDB(name string) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	db, ok := d.dbMap[name]
	if !ok {
		return errors.New("database not found")
	}

	if err := db.Close(); err != nil {
		return err
	}

	delete(d.dbMap, name)
	d.onceMap.Delete(name)
	return nil
}

func (d *DBManager) CloseAll() error {
	d.mu.Lock()
	defer d.mu.Unlock()

	var firstErr error
	for name, db := range d.dbMap {
		if err := db.Close(); err != nil && firstErr == nil {
			firstErr = err
		}
		delete(d.dbMap, name)
		d.onceMap.Delete(name)
	}
	return firstErr
}
