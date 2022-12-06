package conf

import (
	"github.com/fsnotify/fsnotify"
	"github.com/google/martian/v3/log"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"os"
	"snake-and-ladder/utils"
	"strings"
	"sync/atomic"
)

// configFileName 定义业务配置文件名称
const configFileName = "config.toml"

// Config 配置结构，与本地或者配置中心的toml配置文件对应
type Config struct {
	DB     DBConfig     `toml:"db"`
	HTTP   HTTPConfig   `toml:"http"`
	Redis  RedisConfig  `toml:"redis"`
	Server ServerConfig `toml:"server"`
}

var (
	// globalConf 为保证并发安全，使用原子变量存储全局配置文件
	globalConf atomic.Value
	//unmarshaler = config.GetUnmarshaler("toml")
)

func Init() error {
	log.Infof("start init config")
	defer func() {
		log.Infof("config init success: %s", utils.AnyToString(GetConf()))
	}()

	configPath := os.Getenv("CONFIG_PATH")
	if configPath != "" {
		err := readLocalConfigFile(configPath)
		if err != nil {
			log.Errorf("get env CONFIG_PATH: %s, but read failed: %v", configPath, err)
			return err
		}

	} else {
		err := readLocalConfigFile(configFileName)
		if err != nil {
			log.Errorf("get env form config file: %s, but read failed: %v", configFileName, err)
			return err
		}
	}

	return nil
}

// GetConf 获取全局配置文件
func GetConf() *Config {
	return globalConf.Load().(*Config)
}
func readLocalConfigFile(configPath string) error {
	viper.SetConfigFile(configPath)
	// 支持读取环境变量，设置环境变量时，需使用下划线作为分隔符
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	// 将读到的配置文件存储到全局变量中
	err = storeLocalConf()
	if err != nil {
		return err
	}

	// 监听本地配置文件，热加载
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Infof("config watcher: %s", e.String())
		err := storeLocalConf()
		if err != nil {
			log.Errorf("config watcher, store conf: %v", err)
		}
	})
	viper.WatchConfig()
	log.Infof("config watcher running")
	return nil
}

func storeLocalConf() error {
	conf := &Config{}
	err := viper.Unmarshal(conf, func(decoderConfig *mapstructure.DecoderConfig) {
		decoderConfig.TagName = "toml"
	})
	if err != nil {
		return err
	}
	globalConf.Store(conf)
	return nil
}
