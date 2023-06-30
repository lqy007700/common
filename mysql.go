package common

import "github.com/asim/go-micro/v3/config"

type MysqlConfig struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Pwd      string `json:"pwd"`
	Database string `json:"database"`
	Port     int    `json:"port"`
}

func GetMysqlFormConsul(config config.Config, path ...string) (*MysqlConfig, error) {
	c := &MysqlConfig{}
	err := config.Get(path...).Scan(c)
	return c, err
}
