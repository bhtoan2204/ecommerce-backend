package postgresql

import (
	"context"
	"fmt"
	"time"

	persistent_object "payment/app/infrastructure/persistent/postgresql/persistent_object"
	"payment/package/logger"
	"payment/package/settings"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDB struct {
	db     *gorm.DB
	config *settings.Config
}

func NewPostgresDB(ctx context.Context, config *settings.Config) *PostgresDB {
	p := &PostgresDB{config: config}
	p.initDB(ctx)
	p.migrateTables()
	return p
}

func (p *PostgresDB) GetDB() *gorm.DB {
	p.ensureConnection()
	return p.db
}

func (p *PostgresDB) initDB(ctx context.Context) {
	log := logger.FromContext(ctx)
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		p.config.PostgresConfig.Host,
		p.config.PostgresConfig.Port,
		p.config.PostgresConfig.Username,
		p.config.PostgresConfig.Password,
		p.config.PostgresConfig.Database,
	)
	log.Info("Connecting to Postgres DB", zap.String("dsn", dsn))
	var err error
	for i := 1; i <= 5; i++ {
		p.db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			log.Info("Connected to DB",
				zap.Int("attempt", i),
			)
			p.setPool()
			return
		}
		log.Warn("Initial DB connection failed",
			zap.Int("attempt", i),
			zap.Error(err),
		)
		time.Sleep(2 * time.Second)
	}
	log.Fatal("Failed to connect to DB after 5 attempts",
		zap.Error(err),
	)
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
		&persistent_object.Outbox{},
		&persistent_object.Payment{},
		&persistent_object.Refund{},
		&persistent_object.Settlement{},
		&persistent_object.WalletTransaction{},
		&persistent_object.Wallet{},
		&persistent_object.WebhookLog{},
	}
	err := p.db.AutoMigrate(models...)
	if err != nil {
		panic(err)
	}
}

func (p *PostgresDB) ensureConnection() {
	log := logger.DefaultLogger()
	sqlDB, err := p.db.DB()
	if err != nil {
		log.Info("failed to get sql.DB handle",
			zap.Error(err),
		)
		p.reconnectDB()
		return
	}

	if err := sqlDB.Ping(); err != nil {
		log.Info("database disconnected",
			zap.Error(err),
		)
		p.reconnectDB()
	}
}

func (p *PostgresDB) reconnectDB() {
	log := logger.DefaultLogger()
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
			log.Info("Database reconnected successfully",
				zap.Int("attempt", i),
			)
			return
		}
		log.Warn("Reconnect attempt failed",
			zap.Int("attempt", i),
			zap.Error(err),
		)
		time.Sleep(2 * time.Second)
	}
	log.Fatal("Failed to reconnect to database after 3 attempts",
		zap.Error(err),
	)
}
