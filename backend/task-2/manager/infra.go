package manager

import (
	"EkoEdyPurwanto/task-2/config"
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/lib/pq"
)

type (
	InfraManager interface {
		Conn() *sql.DB
	}

	infraManager struct {
		DB  *sql.DB
		Cfg *config.Config
	}
)

func (i *infraManager) initDb() error {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
		i.Cfg.DbConfig.Host,
		i.Cfg.DbConfig.Port,
		i.Cfg.DbConfig.User,
		i.Cfg.DbConfig.Password,
		i.Cfg.DbConfig.Name,
		i.Cfg.DbConfig.SSLMode,
		i.Cfg.DbConfig.TimeZone,
	)
	db, err := sql.Open(i.Cfg.DbConfig.Driver, dsn)
	if err != nil {
		return err
	}
	i.DB = db
	fmt.Printf("successfully connected to database: 🚀%s🚀\n", i.Cfg.DbConfig.Name)

	return nil
}

func (i *infraManager) Conn() *sql.DB {
	return i.DB
}

var (
	instance     InfraManager
	once         sync.Once
	creationLock sync.Mutex
)

// NewInfraManager constructor use Singeleton design pattern
func NewInfraManager(cfg *config.Config) (InfraManager, error) {
	once.Do(func() {
		creationLock.Lock()
		defer creationLock.Unlock()

		if instance == nil {
			conn := &infraManager{
				Cfg: cfg,
			}
			err := conn.initDb()
			if err != nil {
				// If error occurs during initialization, set instance to nil
				instance = nil
				return
			}
			instance = conn
		}
	})
	return instance, nil
}
