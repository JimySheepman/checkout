package config

import (
	"errors"
	"flag"
	"os"
	"strings"

	"github.com/magiconair/properties"
)

var Config *Configurations

func Load() (cfg *Configurations, err error) {
	var path string
	for i, arg := range os.Args {
		if arg == "-v" {
			return nil, nil
		}
		path, err = parseCfg(os.Args, i)
		if err != nil {
			return nil, err
		}
		if path != "" {
			break
		}
	}
	p, err := loadProperties(path)
	if err != nil {
		return nil, err
	}

	Config, err = load(p)
	return Config, err
}

var errInvalidConfig = errors.New("invalid or missing path to config file")

func parseCfg(args []string, i int) (path string, err error) {
	if len(args) == 0 || i >= len(args) || !strings.HasPrefix(args[i], "-cfg") {
		return "", nil
	}
	arg := args[i]
	if arg == "-cfg" {
		if i >= len(args)-1 {
			return "", errInvalidConfig
		}
		return args[i+1], nil
	}

	if !strings.HasPrefix(arg, "-cfg=") {
		return "", errInvalidConfig
	}

	path = arg[len("-cfg="):]
	switch {
	case path == "":
		return "", errInvalidConfig
	case path[0] == '\'':
		path = strings.Trim(path, "'")
	case path[0] == '"':
		path = strings.Trim(path, "\"")
	}
	if path == "" {
		return "", errInvalidConfig
	}
	return path, nil
}

func loadProperties(path string) (p *properties.Properties, err error) {
	if path == "" {
		return properties.NewProperties(), nil
	}
	if strings.HasPrefix(path, "http://") || strings.HasPrefix(path, "https://") {
		return properties.LoadURL(path)
	}
	return properties.LoadFile(path, properties.UTF8)
}

func load(p *properties.Properties) (cfg *Configurations, err error) {
	cfg = &Configurations{}

	f := NewFlagSet(os.Args[0], flag.ExitOnError)

	// dummy values which were parsed earlier
	f.String("cfg", "", "Path or URL to config file")
	f.Bool("v", false, "Show version")

	// config values

	// server
	f.StringVar(&cfg.Server.ServerType, "server.restserver.addr", Default.Server.ServerType, "server type")

	f.StringVar(&cfg.Server.RestServer.Addr, "server.rpcapiaddr", Default.Server.RestServer.Addr, "rest server addr")
	f.IntVar(&cfg.Server.RestServer.PprofEnable, "server.restserver.pprofenable", Default.Server.RestServer.PprofEnable, "rest server pprof enable")

	f.StringVar(&cfg.Server.FileServer.InputPath, "server.fileserver.inputpath", Default.Server.FileServer.InputPath, "file server input path")
	f.StringVar(&cfg.Server.FileServer.OutputPath, "server.fileserver.outputpath", Default.Server.FileServer.OutputPath, "file server output path")

	// MongoDB params
	f.StringVar(&cfg.MongoDB.Addr, "mongodb.addr", Default.MongoDB.Addr, "MongoDB addr")
	f.IntVar(&cfg.MongoDB.Port, "mongodb.port", Default.MongoDB.Port, "MongoDB port")
	f.StringVar(&cfg.MongoDB.User, "mongodb.user", Default.MongoDB.User, "MongoDB user")
	f.StringVar(&cfg.MongoDB.Password, "mongodb.password", Default.MongoDB.Password, "MongoDB password")
	f.IntVar(&cfg.MongoDB.Timeout, "mongodb.timeout", Default.MongoDB.Timeout, "MongoDB timeout")
	f.StringVar(&cfg.MongoDB.Name, "mongodb.name", Default.MongoDB.Name, "MongoDB name")
	f.StringVar(&cfg.MongoDB.Collection, "mongodb.collection", Default.MongoDB.Collection, "MongoDB collection")

	// Logger
	f.StringVar(&cfg.LoggerConfig.AppName, "loggerconfig.appname", Default.LoggerConfig.AppName, "Logger app name")
	f.IntVar(&cfg.LoggerConfig.LogLevel, "loggerconfig.loglevel", Default.LoggerConfig.LogLevel, "Logger log level")
	f.StringVar(&cfg.LoggerConfig.EnvironmentType, "loggerconfig.environmenttype", Default.LoggerConfig.EnvironmentType, "Logger environment type")
	f.StringVar(&cfg.LoggerConfig.LogEncoding, "loggerconfig.logencoding", Default.LoggerConfig.LogEncoding, "Logger log encoding")
	f.StringVar(&cfg.LoggerConfig.GraylogAddr, "loggerconfig.graylogaddr", Default.LoggerConfig.GraylogAddr, "Logger Graylog Addr")

	// filter out -test flags
	var args []string
	for _, a := range os.Args[1:] {
		if strings.HasPrefix(a, "-test.") {
			continue
		}
		args = append(args, a)
	}

	prefixes := []string{""}
	if err := f.ParseFlags(args, os.Environ(), prefixes, p); err != nil {
		return nil, err
	}

	return cfg, nil
}
