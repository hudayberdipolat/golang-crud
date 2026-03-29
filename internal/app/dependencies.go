package app

import (
	"github.com/hudayberdipolat/golang-crud/pkg/config"
	"gorm.io/gorm"
)

type Dependency struct {
	Config   *config.Config
	DbConfig *gorm.DB
}

func GetAppDependencies() (*Dependency, error) {
	getConfig, err := config.GetConfig()
	if err != nil {
		return nil, err
	}

	// getDBConfig := dbconfig.NewDatabaseConfig(getConfig)
	// getDbConnection, err := getDBConfig.GetDBConnection()
	// if err != nil {
	// 	return nil, err
	// }

	return &Dependency{
		Config:   getConfig,
		DbConfig: nil,
	}, nil
}
