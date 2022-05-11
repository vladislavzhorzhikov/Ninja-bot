package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	DBUser       string `json:"db_user"`
	DBHost       string `json:"db_host"`
	DBPort       int    `json:"db_port"`
	DBPassword   string `json:"db_password"`
	DBName       string `json:"db_name"`
	RandomsTable string `json:"randoms_table"`
	DBSSLmode    string `json:"db_sslmode"`
	TGToken      string `json:"tg_token"`
}

func Get() (*Config, error) {
	bytes, err := ioutil.ReadFile("./.config/local.json")
	if err != nil {
		return nil, err
	}
	return read(bytes)
}

func read(data []byte) (*Config, error) {
	conf := Config{}
	err := json.Unmarshal(data, &conf)
	if err != nil {
		return nil, err
	}
	return &conf, nil
}

func (c *Config) GetTGToken() string {
	if c != nil {
		return c.TGToken
	}
	return ""
}

func (c *Config) GetDatabaseDSN() string {
	dsn := ""
	if c != nil {
		if c.DBHost != "" {
			dsn += fmt.Sprintf("host=%s", c.DBHost)
		}
		if c.DBPort != 0 {
			dsn += fmt.Sprintf(" port=%d", c.DBPort)
		}
		if c.DBUser != "" {
			dsn += fmt.Sprintf(" user=%s", c.DBUser)
		}
		if c.DBPassword != "" {
			dsn += fmt.Sprintf(" password=%s", c.DBPassword)
		}
		if c.DBName != "" {
			dsn += fmt.Sprintf(" dbname=%s", c.DBName)
		}
		if c.DBSSLmode != "" {
			dsn += fmt.Sprintf(" sslmode=%s", c.DBSSLmode)
		}
	}
	return dsn
}

func (c *Config) GetRandomTableName() string {
	if c != nil {
		return c.RandomsTable
	}
	return ""
}
