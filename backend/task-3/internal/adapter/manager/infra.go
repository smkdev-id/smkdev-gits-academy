package manager

import (
	"EkoEdyPurwanto/task-3/config"
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
		i.Cfg.PostgresConfig.Host,
		i.Cfg.PostgresConfig.Port,
		i.Cfg.PostgresConfig.User,
		i.Cfg.PostgresConfig.Password,
		i.Cfg.PostgresConfig.Name,
		i.Cfg.PostgresConfig.SSLMode,
		i.Cfg.PostgresConfig.TimeZone,
	)
	db, err := sql.Open(i.Cfg.PostgresConfig.Driver, dsn)
	if err != nil {
		return err
	}
	i.DB = db
	fmt.Printf("successfully connected to database: 🚀%s🚀\n", i.Cfg.PostgresConfig.Name)

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
