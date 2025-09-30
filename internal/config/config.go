package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type SerialConfig struct {
	Port      string `yaml:"port"`
	BaudRate  int    `yaml:"baud_rate"`
	DataBits  int    `yaml:"data_bits"`
	Parity    string `yaml:"parity"`
	StopBits  int    `yaml:"stop_bits"`
	TimeoutMs int    `yaml:"timeout_ms"`
}

type SlaveConfig struct {
	ID                byte   `yaml:"id"`
	Name              string `yaml:"name"`
	PollIntervalMs    int    `yaml:"poll_interval_ms"`
	LevelRegisterAddr uint16 `yaml:"level_register_addr"`
	StartCoilAddr     uint16 `yaml:"start_coil_addr"`
	StopCoilAddr      uint16 `yaml:"stop_coil_addr"`
}

type InfluxConfig struct {
	URL   string `yaml:"url"`
	Token string `yaml:"token"`
	Org   string `yaml:"org"`
	Bucket string `yaml:"bucket"`
}

type PostgresConfig struct {
	DSN string `yaml:"dsn"`
}

type AppConfig struct {
	HTTPAddr string        `yaml:"http_addr"`
	Serial   SerialConfig  `yaml:"serial"`
	Slaves   []SlaveConfig `yaml:"slaves"`
	Influx   InfluxConfig  `yaml:"influx"`
	Postgres PostgresConfig `yaml:"postgres"`
}

func Load(path string) (*AppConfig, error) {
	cfg := &AppConfig{}
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read config: %w", err)
	}
	if err := yaml.Unmarshal(b, cfg); err != nil {
		return nil, fmt.Errorf("parse yaml: %w", err)
	}
	if v := os.Getenv("APP_ADDR"); v != "" {
		cfg.HTTPAddr = v
	}
	if v := os.Getenv("INFLUX_URL"); v != "" {
		cfg.Influx.URL = v
	}
	if v := os.Getenv("INFLUX_TOKEN"); v != "" {
		cfg.Influx.Token = v
	}
	if v := os.Getenv("INFLUX_ORG"); v != "" {
		cfg.Influx.Org = v
	}
	if v := os.Getenv("INFLUX_BUCKET"); v != "" {
		cfg.Influx.Bucket = v
	}
	if v := os.Getenv("POSTGRES_DSN"); v != "" {
		cfg.Postgres.DSN = v
	}
	if v := os.Getenv("SERIAL_PORT"); v != "" {
		cfg.Serial.Port = v
	}
	return cfg, nil
}