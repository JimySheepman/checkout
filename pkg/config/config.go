package config

type Configurations struct {
	Server       Server
	MongoDB      MongoDB
	LoggerConfig LoggerConfig
}

type Server struct {
	RestServer RestServer
	FileServer FileServer
}

type RestServer struct {
	Addr        string
	PprofEnable int
}

type FileServer struct {
	InputPath  string
	OutputPath string
}

type LoggerConfig struct {
	AppName         string
	LogLevel        int
	LogEncoding     string
	GraylogAddr     string
	EnvironmentType string
}

type MongoDB struct {
	Addr       string
	Port       int
	User       string
	Password   string `json:"-"`
	Timeout    int
	Name       string
	Collection string
}
