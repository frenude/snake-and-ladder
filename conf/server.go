package conf

type ServerConfig struct {
	Version             string `toml:"version"`
	API_SECRET          string `toml:"api_secret"`
	TOKEN_HOUR_LIFESPAN int    `toml:"token_hour_lifespan"`
}
