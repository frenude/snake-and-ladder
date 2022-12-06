package conf

type RedisConfig struct {
	Addrs        []string `toml:"addrs"`             // 地址列表,使用 host:port 模式
	Port         uint     `toml:"port"`              // 地址列表,使用 host:port 模式
	DB           int      `toml:"db"`                // db 号
	Password     string   `toml:"password" json:"-"` // 密码,屏蔽 json 标签,不要打印到日志
	MaxRetries   int      `toml:"max_retries"`       // 最大重试次数
	DialTimeout  int      `toml:"dial_timeout"`      // 连接超时时间,单位毫秒
	ReadTimeout  int      `toml:"read_timeout"`      // 读超时时间,单位毫秒
	WriteTimeout int      `toml:"write_timeout"`     // 写超时时间,单位毫秒
	PoolSize     int      `toml:"pool_size"`         // 连接池大小
	MinIdleConns int      `toml:"min_idle_conns"`    // 最小空闲连接数
	MaxConnAge   int      `toml:"max_conn_age"`      // 最大生命周期,单位秒
	MasterName   string   `toml:"master_name"`       // 哨兵模式时填写
}
