package dbconfig

import (
	"fmt"

	"github.com/hudayberdipolat/golang-crud/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	config *config.Config
}

func NewDatabaseConfig(config *config.Config) *DatabaseConfig {
	return &DatabaseConfig{
		config: config,
	}
}

func (databaseConfig DatabaseConfig) GetDBConnection() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		databaseConfig.config.DbConfig.DbHost,
		databaseConfig.config.DbConfig.DbUser,
		databaseConfig.config.DbConfig.DbPassword,
		databaseConfig.config.DbConfig.DbName,
		databaseConfig.config.DbConfig.DbPort,
		databaseConfig.config.DbConfig.DbSslMode,
		databaseConfig.config.DbConfig.DbTimeZone,
	)
	dbConnection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return dbConnection, nil
}
