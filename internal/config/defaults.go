package config

import (
	"github.com/spf13/viper"
)

func SetDefaults(v *viper.Viper) {
	v.SetDefault("http.listen", ":80")
	v.SetDefault("log.level", "5")

	v.SetDefault("db.host", "mysql-pismo")
	v.SetDefault("db.port", 3306)
	v.SetDefault("db.user", "root")
	v.SetDefault("db.pass", "abc123")
	v.SetDefault("db.name", "pismo")
}
