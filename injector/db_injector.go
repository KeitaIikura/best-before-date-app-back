package injector

import (
	"bbdate/pkg/config"
	"bbdate/pkg/db"
	"bbdate/pkg/logging"
	"fmt"
)

func provideDBConnections() {

	dbConf := db.MySQLConnectorConfig{
		DBUser:       config.EnvStore.DBUser,
		DBPass:       config.EnvStore.DBPassword,
		DBReaderHost: config.EnvStore.DBReaderAddress,
		DBWriterHost: config.EnvStore.DBWriterAddress,
		DBName:       config.EnvStore.DBName,
	}

	if err := c.Provide(func() db.MySQLConnectorConfig {
		return dbConf
	}); err != nil {
		logging.Fatal("system", fmt.Sprintf("failed to provide =>%v", err))

	}
	if err := c.Provide(db.NewMySQLConnector); err != nil {
		logging.Fatal("system", fmt.Sprintf("failed to provide =>%v", err))
	}

}
