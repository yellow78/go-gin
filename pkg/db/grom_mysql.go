package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"sync"

	"gorm.io/gorm"

	mysqldriver "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
)

type gromManager struct {
	gromMap map[string]*gorm.DB
	onceMap sync.Map // map[string]*sync.Once
	mu      sync.RWMutex
}

func NewGromManager() *gromManager {
	return &gromManager{
		gromMap: make(map[string]*gorm.DB),
	}
}

func (g *gromManager) InitDBWithConfig(cfg *mysqldriver.Config, name string, gormCfg ...*gorm.Config) error {
	val, _ := g.onceMap.LoadOrStore(name, &sync.Once{})
	once := val.(*sync.Once)

	var initErr error
	once.Do(func() {
		var gcfg *gorm.Config
		if len(gormCfg) > 0 {
			gcfg = gormCfg[0]
		} else {
			gcfg = &gorm.Config{}
		}

		// 使用 mysql.NewConnector + sql.OpenDB
		connector, err := mysqldriver.NewConnector(cfg)
		if err != nil {
			initErr = fmt.Errorf("create connector failed: %w", err)
			return
		}

		stdDB := sql.OpenDB(connector)
		if err := stdDB.Ping(); err != nil {
			initErr = fmt.Errorf("ping failed: %w", err)
			return
		}

		// 透過已有的 sql.DB 實例建立 gorm.DB
		db, err := gorm.Open(mysql.New(mysql.Config{
			Conn: stdDB,
		}), gcfg)
		if err != nil {
			initErr = fmt.Errorf("gorm open failed: %w", err)
			return
		}

		g.mu.Lock()
		g.gromMap[name] = db
		g.mu.Unlock()
	})

	return initErr
}

func (g *gromManager) GetDB(name string) (*gorm.DB, bool) {
	g.mu.RLock()
	defer g.mu.RUnlock()
	db, ok := g.gromMap[name]
	return db, ok
}

func (g *gromManager) CloseDB(name string) error {
	g.mu.Lock()
	defer g.mu.Unlock()

	db, ok := g.gromMap[name]
	if !ok {
		return errors.New("database not found")
	}

	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("get sql.DB failed: %w", err)
	}

	if err := sqlDB.Close(); err != nil {
		return fmt.Errorf("close failed: %w", err)
	}

	delete(g.gromMap, name)
	g.onceMap.Delete(name)
	return nil
}

func (g *gromManager) CloseAll() error {
	g.mu.Lock()
	defer g.mu.Unlock()

	var firstErr error
	for name, db := range g.gromMap {
		sqlDB, err := db.DB()
		if err != nil && firstErr == nil {
			firstErr = fmt.Errorf("get sql.DB for %s failed: %w", name, err)
		} else if closeErr := sqlDB.Close(); closeErr != nil && firstErr == nil {
			firstErr = fmt.Errorf("close %s failed: %w", name, closeErr)
		}
		delete(g.gromMap, name)
		g.onceMap.Delete(name)
	}
	return firstErr
}
