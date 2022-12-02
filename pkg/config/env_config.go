package config

import (
	"bbdate/pkg/logging"
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type (
	EnvMode string
	// NOTE DB接続情報をenv使用とする場合は複数データベースの情報を取得できるよう修正が必要なことに注意
	EnvConfig struct {
		EnvMode            EnvMode `envconfig:"ENV_MODE" default:"local"`
		HashSecret         string  `envconfig:"HASH_SECRET" default:"hogehoge"`
		RedisAddress       string  `envconfig:"REDIS_ADDRESS" default:"redis:6379"`
		RedisPass          string  `envconfig:"REDIS_PASS" default:""`
		DBReaderAddress    string  `envconfig:"DB_READER_ADDRESS" default:"db:3306"`
		DBWriterAddress    string  `envconfig:"DB_WRITER_ADDRESS" default:"db:3306"`
		DBName             string  `envconfig:"DB_NAME" default:"bbdate"`
		DBUser             string  `envconfig:"DB_USER" default:"root"`
		DBPassword         string  `envconfig:"DB_PASSWORD" default:"root"`
		CORSAcceptOrigin   string  `envconfig:"CORS_ACCEPT_ORIGIN" default:""`
		TmxSessionTimeout  int     `envconfig:"TMX_SESSION_TIMEOUT" default:"3600"`
		DefaultFromAddress string  `envconfig:"DEFAULT_FROM_ADDRESS" default:"from@example.com"`
		SmtpHost           string  `envconfig:"SMTP_HOST" default:"mail"`
		SmtpUserName       string  `envconfig:"SMTP_USERNAME" default:""`
		SmtpPassword       string  `envconfig:"SMTP_PASSWORD" default:""`
		SmtpPort           int     `envconfig:"SMTP_PORT" default:"1025"`
	}
)

// envmode
const (
	Local EnvMode = "local"
	Dev   EnvMode = "dev"
	Test  EnvMode = "test"
	Prod  EnvMode = "prod"
)

var (
	EnvStore *EnvConfig = &EnvConfig{}
)

func init() {
	if err := envconfig.Process("", EnvStore); err != nil {
		logging.Fatal("", fmt.Sprintf("failed to read value from env: error=> %v", err))
	}

	logEnvValues()
}

func logEnvValues() {
	logging.Info("", fmt.Sprintf("Read from Env: EnvMode=[%s] ", EnvStore.EnvMode))
	logging.Info("", fmt.Sprintf("Read from Env: CORSAcceptOrigin=[%s] ", EnvStore.CORSAcceptOrigin))
	logging.Info("", fmt.Sprintf("Read from Env: TmxSessionTimeout=[%d] ", EnvStore.TmxSessionTimeout))

	if EnvStore.EnvMode == Local {
		logging.Info("", fmt.Sprintf("Read from Env: DefaultFromAddress=[%s] ", EnvStore.DefaultFromAddress))
		logging.Info("", fmt.Sprintf("Read from Env: RedisAddress=[%s] ", EnvStore.RedisAddress))
		logging.Info("", fmt.Sprintf("Read from Env: RedisPass=[%s] ", EnvStore.RedisPass))
		logging.Info("", fmt.Sprintf("Read from Env: DBReaderAddress=[%s] ", EnvStore.DBReaderAddress))
		logging.Info("", fmt.Sprintf("Read from Env: DBWriterAddress=[%s] ", EnvStore.DBWriterAddress))
		logging.Info("", fmt.Sprintf("Read from Env: DBName=[%s] ", EnvStore.DBName))
		logging.Info("", fmt.Sprintf("Read from Env: DBUser=[%s] ", EnvStore.DBUser))
		logging.Info("", fmt.Sprintf("Read from Env: DBPassword=[%s] ", EnvStore.DBPassword))
		logging.Info("", fmt.Sprintf("Read from Env: SmtpHost=[%s] ", EnvStore.SmtpHost))
		logging.Info("", fmt.Sprintf("Read from Env: SmtpPort=[%d] ", EnvStore.SmtpPort))
	}
}
