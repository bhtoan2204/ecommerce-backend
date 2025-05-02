package postgresql

import (
	"fmt"
	"log"
	"time"

	persistentobject "user/app/infrastructure/persistent/postgresql/persistent_object"
	"user/package/settings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDB struct {
	db     *gorm.DB
	config *settings.Config
}

func NewPostgresDB(config *settings.Config) *PostgresDB {
	p := &PostgresDB{config: config}
	p.initDB()
	p.migrateTables()
	return p
}

func (p *PostgresDB) GetDB() *gorm.DB {
	p.ensureConnection()
	return p.db
}

func (p *PostgresDB) initDB() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		p.config.PostgresConfig.Host,
		p.config.PostgresConfig.Port,
		p.config.PostgresConfig.Username,
		p.config.PostgresConfig.Password,
		p.config.PostgresConfig.Database,
	)
	var err error
	p.db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	p.setPool()
	log.Println("Database connection established")
}

func (p *PostgresDB) setPool() {
	sqlDb, err := p.db.DB()
	if err != nil {
		panic(err)
	}
	sqlDb.SetMaxOpenConns(p.config.PostgresConfig.MaxOpenConns)
	sqlDb.SetMaxIdleConns(p.config.PostgresConfig.MaxIdleConns)
	sqlDb.SetConnMaxLifetime(time.Duration(p.config.PostgresConfig.MaxLifetime) * time.Second)
}

func (p *PostgresDB) migrateTables() {
	models := []interface{}{
		&persistentobject.User{},
		&persistentobject.Address{},
	}
	err := p.db.AutoMigrate(models...)
	if err != nil {
		panic(err)
	}
}

func (p *PostgresDB) ensureConnection() {
	sqlDB, err := p.db.DB()
	if err != nil {
		log.Printf("failed to get sql.DB handle: %v", err)
		p.reconnectDB()
		return
	}

	if err := sqlDB.Ping(); err != nil {
		log.Printf("database disconnected: %v. Reconnecting...", err)
		p.reconnectDB()
	}
}

func (p *PostgresDB) reconnectDB() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		p.config.PostgresConfig.Host,
		p.config.PostgresConfig.Port,
		p.config.PostgresConfig.Username,
		p.config.PostgresConfig.Password,
		p.config.PostgresConfig.Database,
	)

	var err error
	for i := 1; i <= 3; i++ {
		p.db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			log.Printf("Database reconnected successfully on attempt %d", i)
			return
		}
		log.Printf("Reconnect attempt %d failed: %v", i, err)
		time.Sleep(2 * time.Second)
	}
	log.Fatalf("Failed to reconnect to database after 3 attempts: %v", err)
}
