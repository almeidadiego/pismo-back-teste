package config

import (
	"database/sql"
	"fmt"
	"os"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql" // load mysql driver
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Load() {
	v := viper.GetViper()
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	SetDefaults(v)
}

func InitLog() {
	log.SetFormatter(&log.JSONFormatter{})

	log.SetOutput(os.Stdout)

	level := viper.GetUint32("log.level")
	log.SetLevel(log.Level(level))
}

func InitDatabase() (*sql.DB, error) {
	host := viper.GetString("db.host")
	port := viper.GetInt("db.port")
	user := viper.GetString("db.user")
	password := viper.GetString("db.pass")
	name := viper.GetString("db.name")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&charset=utf8",
		user, password, host, port, name)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// Open doesn't open a connection. Validate DSN data:
	for i := 0; i < 30; i++ {
		err = db.Ping()
		if err != nil {
			time.Sleep(time.Second * 5)
		} else {
			break
		}
	}

	if err != nil {
		return nil, err
	}

	return db, nil
}
