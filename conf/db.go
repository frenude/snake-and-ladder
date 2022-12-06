package conf

type DBConfig struct {
	Host         string `toml:"host"`
	Port         int32  `toml:"port"`
	User         string `toml:"user"`
	Password     string `toml:"password" json:"-"` // 密码屏蔽 json 标签,不要打印到日志
	Charset      string `toml:"charset"`
	DataBase     string `toml:"data_base"`
	MaxConn      int    `toml:"max_conn"`       // 最大连接数
	MaxIdle      int    `toml:"max_idle"`       // 最大闲置数
	ConnTimeout  int    `toml:"conn_timeout"`   // 超时时间,单位毫秒
	ReadTimeout  int    `toml:"read_timeout"`   // 读超时时间,单位毫秒
	WriteTimeout int    `toml:"write_timeout"`  // 写超时时间,单位毫秒
	ConnLifeTime int    `toml:"conn_life_time"` // 最大生命周期,单位秒
	ParseTime    bool   `toml:"parse_time"`     // 查询结果是否自动解析为时间
	TimeZone     string `toml:"time_zone"`      // mysql可填Local自动获取系统时区,其他填具体的时区(Asia/Shanghai)
	IsPostgres   bool   `toml:"is_postgres"`    // 是否 postgres 数据库
}
