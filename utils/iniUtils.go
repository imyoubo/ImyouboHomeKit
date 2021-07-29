package utils

import (
	"github.com/go-ini/ini"
)

type IniConfig struct {
	path string
	Conf *ini.File
}

func LoadIni(path string) (*IniConfig, error) {
	cfg, err := ini.Load(path)
	return &IniConfig{path: path, Conf: cfg}, err
}

func (c *IniConfig) GetValue(section string, key string) string {
	return c.Conf.Section(section).Key(key).Value()
}

func (c *IniConfig) GetOrElse(section string, key string, def string) string {
	if v := c.GetValue(section, key); v != "" {
		return v
	} else {
		return def
	}
}

func (c *IniConfig) Reload() error {
	cfg, err := ini.Load(c.path)
	if err != nil {
		return err
	}
	c.Conf = cfg
	return nil
}
