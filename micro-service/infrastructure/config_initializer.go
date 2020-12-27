package infrastructure

import (
    "fmt"
    "github.com/fsnotify/fsnotify"
    "github.com/spf13/viper"
)

// Config struct
type Config struct {
    Name string
}

// InitConfig ()
func InitConfig(cfg string) error {
    c := Config{
        Name: cfg,
    }

    if err := c.initConfig(); err != nil {
        return err
    }

    c.watchConfig()

    return nil
}

func (c *Config) initConfig() error {
    if c.Name != "" {
        viper.SetConfigFile(c.Name)
    } else {
        viper.AddConfigPath("resource/conf")
        viper.SetConfigName("config")
    }

    viper.SetConfigType("yaml")
	
    err := viper.ReadInConfig() 
    if err != nil {
        panic(fmt.Errorf("fatal error config file: %s ", err))
    }
    fmt.Println(viper.GetString("name"))

    return nil
}

func (c *Config) watchConfig() {
    viper.WatchConfig()
    viper.OnConfigChange(func(e fsnotify.Event) {
        fmt.Println("Config file changed:", e.Name)
    })
}