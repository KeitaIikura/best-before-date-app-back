package db

import (
	"bbdate/pkg/logging"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// コネクションのインターフェースと実体
type (
	IMySQL interface {
		GetReaderConn() *sql.DB
		GetWriterConn() *sql.DB
	}
	MySQL struct {
		ReaderConn *sql.DB
		WriterConn *sql.DB
	}
	MySQLConnectorConfig struct {
		DBUser       string
		DBPass       string
		DBReaderHost string
		DBWriterHost string
		DBName       string
	}
)

func NewMySQLConnector(config MySQLConnectorConfig) IMySQL {
	logging.Info("system", "NewMySQLConnector start")
	result := &MySQL{}
	var err error
	// reader
	err = result.establishConnection(config, true)
	if err != nil {
		logging.Fatal("system", fmt.Sprintf("new mysql connector on establishReaderConnection: %v", err))
	}
	// writer
	err = result.establishConnection(config, false)
	if err != nil {
		logging.Fatal("system", fmt.Sprintf("new mysql connector on establishWriterConnection: %v", err))
	}

	logging.Info("system", "NewMySQLConnector end")
	return result
}

func (m *MySQL) establishConnection(conf MySQLConnectorConfig, reader bool) error {
	var host string
	if reader {
		host = conf.DBReaderHost
	} else {
		host = conf.DBWriterHost
	}

	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true",
		conf.DBUser,
		conf.DBPass,
		host,
		conf.DBName,
	)
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("dsn(%s) open database connection error: %w", dsn, err)
	}

	// TODO コネクション数などオプションの設定
	// m.Connection.SetMaxOpenConns(10)
	// m.Connection.SetMaxIdleConns(10)
	// m.Connection.SetConnMaxLifetime(time.Duration(10) * time.Second)
	if reader {
		m.ReaderConn = conn
		return m.ReaderConn.Ping()
	} else {
		m.WriterConn = conn
		return m.WriterConn.Ping()
	}
}

func (m *MySQL) GetReaderConn() *sql.DB {
	return m.ReaderConn
}

func (m *MySQL) GetWriterConn() *sql.DB {
	return m.WriterConn
}
