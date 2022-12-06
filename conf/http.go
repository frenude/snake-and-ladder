package conf

type HTTPConfig struct {
	Host string `toml:"host"`
	Port uint   `toml:"port"`
}
