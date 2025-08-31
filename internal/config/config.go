package config

import (
	"encoding/json"
	"flag"
	"os"
	"path/filepath"
)

type Config struct {
	AppName    string `json:"app_name"`
	Version    string `json:"version"`
	Debug      bool   `json:"debug"`
	ConfigPath string `json:"-"`
	
	Theme      ThemeConfig      `json:"theme"`
	Keybindings KeybindingConfig `json:"keybindings"`
}

type ThemeConfig struct {
	ColorScheme string `json:"color_scheme"`
	UseEmoji    bool   `json:"use_emoji"`
}

type KeybindingConfig struct {
	Quit   []string `json:"quit"`
	Help   []string `json:"help"`
	Input  []string `json:"input"`
	Refresh []string `json:"refresh"`
}

func New() *Config {
	cfg := &Config{
		AppName: "Bubble Skeleton",
		Version: "0.1.0",
		Debug:   false,
		Theme: ThemeConfig{
			ColorScheme: "default",
			UseEmoji:    true,
		},
		Keybindings: KeybindingConfig{
			Quit:    []string{"q", "ctrl+c"},
			Help:    []string{"h", "?"},
			Input:   []string{"i"},
			Refresh: []string{"r"},
		},
	}

	flag.BoolVar(&cfg.Debug, "debug", false, "Enable debug mode")
	flag.StringVar(&cfg.ConfigPath, "config", "", "Path to config file")
	flag.Parse()

	if cfg.ConfigPath == "" {
		cfg.ConfigPath = cfg.defaultConfigPath()
	}

	cfg.Load()

	return cfg
}

func (c *Config) defaultConfigPath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return filepath.Join(homeDir, ".config", "bubble-skeleton", "config.json")
}

func (c *Config) Load() error {
	if c.ConfigPath == "" {
		return nil
	}

	data, err := os.ReadFile(c.ConfigPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	return json.Unmarshal(data, c)
}

func (c *Config) Save() error {
	if c.ConfigPath == "" {
		return nil
	}

	dir := filepath.Dir(c.ConfigPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(c.ConfigPath, data, 0644)
}