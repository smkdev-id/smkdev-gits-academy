package manager

import (
	"database/sql"
	"fmt"

	"github.com/alwi09/config"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type (
	MysqlManager interface {
		Conn() *sql.DB
	}

	mysqlManager struct {
		DB  *sql.DB
		Cfg *config.Config
	}
)

func NewMysqlManager(cfg *config.Config) (MysqlManager, error) {
	m := &mysqlManager{
		Cfg: cfg,
	}

	if err := m.initDb(); err != nil {
		return nil, err
	}

	return m, nil
}

func (m *mysqlManager) initDb() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		m.Cfg.DB.DB_USER,
		m.Cfg.DB.DB_PASSWORD,
		m.Cfg.DB.DB_HOST,
		m.Cfg.DB.DB_PORT,
		m.Cfg.DB.DB_NAME,
	)

	db, err := sql.Open(m.Cfg.DB.DB_DRIVER, dsn)
	if err != nil {
		return err
	}

	m.DB = db
	fmt.Printf("connected to database: %s\n", m.Cfg.DB.DB_NAME)

	return nil
}

func (m *mysqlManager) Conn() *sql.DB {
	return m.DB
}
