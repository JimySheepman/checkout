package main

import (
	"checkout-case/config"
	utils "checkout-case/utitls"
	"log"
	"runtime"
)

func main() {
	cfg, err := config.Load()
	if err != nil || cfg == nil {
		log.Fatalf("[FATAL] %s", err)
	}

	l := InitLogger(cfg.LoggerConfig).Sugar()

	l.Info("checkout-case is started...")
	l.Infof("Runtime config " + utils.ToJSON(cfg))
	l.Infof("Go runtime is %s", runtime.Version())

	srv := new(service)
	srv.Start()
}
