package config

// Default config
var Default = &Configurations{
	Server: Server{
		ServerType: "rest",
		RestServer: RestServer{
			Addr:        ":8080",
			PprofEnable: 0,
		},
		FileServer: FileServer{
			InputPath:  "./input.txt",
			OutputPath: "./output.txt",
		},
	},
	MongoDB: MongoDB{
		Addr:       "localhost",
		Port:       27017,
		User:       "root",
		Password:   "example",
		Timeout:    5000,
		Name:       "shopping",
		Collection: "carts",
	},
	LoggerConfig: LoggerConfig{
		AppName:         "checkout-case",
		LogLevel:        0,
		LogEncoding:     "console",
		GraylogAddr:     "",
		EnvironmentType: "production",
	},
}
