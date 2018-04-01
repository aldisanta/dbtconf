package dbtconf

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"

	"github.com/BurntSushi/toml"
)

// Config is Config
type Config struct {
	DB Database `toml:"database"`
}

// Database is Database
type Database struct {
	User     string
	Password string
	Dbname   string
	Sslmode  string
}

// Connection is Connection
func Connection(configFile string) (string, error) {
	var config Config
	var buffer bytes.Buffer

	cfg, err := ioutil.ReadFile(configFile)
	// read toml file
	if err != nil {
		return "", err
	}
	// decode toml file
	if _, err := toml.Decode(fmt.Sprintf("%s", cfg), &config); err != nil {
		return "", err
	}

	dbConf := reflect.ValueOf(&config.DB)
	dbConfTyp := reflect.Indirect(dbConf).Type()
	dbConfVal := dbConf.Elem()
	for index := 0; index < dbConfVal.NumField(); index++ {
		buffer.WriteString(strings.ToLower(dbConfTyp.Field(index).Name))
		buffer.WriteString("=")
		buffer.WriteString(dbConfVal.Field(index).Interface().(string))
		buffer.WriteString(" ")
	}
	return strings.TrimSpace(buffer.String()), nil
}
