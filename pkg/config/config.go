package config

import (
	"embed"
	"errors"
	"fmt"
	"os"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/hcl"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/fs"
)

func ProvideConfig() (*Config, error) {
	k, err := Load("config.hcl")
	if err != nil {
		return nil, fmt.Errorf("config.InitConfig: %w", err)
	}

	c, err := UnmarshalConfig(k)
	if err != nil {
		return nil, fmt.Errorf("config.InitConfig: %w", err)
	}

	return c, nil
}

// errors
var (
	ErrEmptyPathList = errors.New("empty paths list")
)

//go:generate cp ../../default.hcl default.hcl
//go:embed default.hcl
var eConfigFile embed.FS

// config.Load: load and merge configs
// reads from embedded config first
// then other config files with the
// later ones having highest precidence
func Load(paths ...string) (*koanf.Koanf, error) {
	if len(paths) == 0 {
		return nil, fmt.Errorf("config.Load: %w", ErrEmptyPathList)
	}

	k := koanf.New("/")

	err := k.Load(fs.Provider(eConfigFile, "default.hcl"), hcl.Parser(false))
	if err != nil {
		return nil, fmt.Errorf("config.Load: %w", err)
	}

	for _, path := range paths {
		err = k.Load(file.Provider(path), hcl.Parser(false))
		if err != nil {
			if os.IsNotExist(err) {
				fmt.Fprintf(os.Stderr, "config.Load: File %q not found", path)
				continue
			}
			return nil, fmt.Errorf("config.Load: %w", err)
		}
	}

	return k, nil
}

type Config struct {
	Mode  string
	Debug bool
}

// config.UnmarshalConfig: unmarshal koanf config
func UnmarshalConfig(k *koanf.Koanf) (*Config, error) {
	var config Config
	err := k.Unmarshal("", &config)
	if err != nil {
		return nil, fmt.Errorf("config.UnmarshallConfig: %w", err)
	}

	return &config, nil
}
