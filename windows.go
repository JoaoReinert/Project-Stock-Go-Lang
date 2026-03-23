package main

import (
	"Desafio_Go_Lang/entities"
	"Desafio_Go_Lang/setup"
	"errors"
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/gorilla/mux"
	"github.com/kardianos/service"
	"io"
	"log/slog"
	"net/http"
	"os"
	"strings"
)

func Start() error {
	configsPath, action := loadFlags()

	config, err := readConfigFile(configsPath)
	if err != nil {
		return errors.Join(fmt.Errorf("failed to create windows service"), err)
	}

	windowsService, err := newService(*config)
	if err != nil {
		return errors.Join(fmt.Errorf("failed to create windows service"), err)
	}

	switch action {
	case "run", "":
		err = windowsService.Run()
		if err != nil {
			return errors.Join(fmt.Errorf("failed to run"), err)
		}

	case "uninstall":
		err = windowsService.Uninstall()
		if err != nil {
			return errors.Join(fmt.Errorf("failed to uninstall"), err)
		}

	case "install":
		err = windowsService.Install()
		if err != nil {
			return errors.Join(fmt.Errorf("failed to install"), err)
		}

	case "stop":
		err = windowsService.Stop()
		if err != nil {
			return errors.Join(fmt.Errorf("failed to stop"), err)
		}

	default:
		return errors.New("invalid action")
	}

	return nil
}

func loadFlags() (string, string) {
	var action string
	var configsPath string

	flag.StringVar(&configsPath, "configs", "", "the path to the application config file")
	flag.StringVar(&action, "action", "", "the action to execute")
	flag.Parse()

	if configsPath == "" {
		configsPath = os.Getenv("configs")
	}

	if configsPath == "" {
		panic("[configs] argument not found")
	}

	return configsPath, action
}

func readConfigFile(cfgPath string) (*entities.Config, error) {
	file, err := os.Open(cfgPath)
	if err != nil {
		return nil, errors.Join(fmt.Errorf("failed to open config file"), err)
	}

	b, err := io.ReadAll(file)
	if err != nil {
		return nil, errors.Join(fmt.Errorf("failed to read config file"), err)
	}
	defer file.Close()

	var config entities.Config

	_, err = toml.Decode(string(b), &config)
	if err != nil {
		return nil, errors.Join(fmt.Errorf("failed to decode config file"), err)
	}

	return &config, nil
}

type program struct {
	server *http.Server
	config entities.Config
}

func (p *program) Start(s service.Service) error {
	slog.Info("Start service")

	go p.run()
	return nil
}

func (p *program) Stop(s service.Service) error {
	slog.Info("Stop service")

	return nil
}

func (p *program) run() {
	router := mux.NewRouter()

	setup.SetupModules(router, p.config)

	server, err := setup.SetupServer(router, p.config)
	if err != nil {
		slog.Error("HTTP server failed", slog.String("error", err.Error()))
		panic(err)
	}

	p.server = server

	slog.Info("Starting HTTP server", slog.String("address", p.server.Addr))
	err = p.server.ListenAndServe()
	if err != nil {
		slog.Error("HTTP server failed", slog.String("error", err.Error()))
		panic(err)
	}
}

func newService(cfg entities.Config) (service.Service, error) {
	slog.Info("Creating service")

	// Load the received arguments
	var args []string

	// Clean the executable arguments
	if len(os.Args) > 1 {
		cleanArgs := make([]string, 0)

		for _, arg := range os.Args {
			if !strings.HasPrefix(arg, "-action") {
				cleanArgs = append(args, arg)
			}
		}

		args = cleanArgs
	}

	svcConfig := &service.Config{
		Name:        "Joao - Desafio Go Lang Authentication",
		DisplayName: "Joao Desafio Go Lang Authentication",
		Description: "Servidor utilizado para um desafio em Go Lang",
		Arguments:   args,
	}

	p := &program{config: cfg}

	s, err := service.New(p, svcConfig)
	if err != nil {
		return nil, err
	}

	return s, nil
}
