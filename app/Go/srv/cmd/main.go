package main

import (
	"crypto/tls"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"net/http"
	"os/user"
	"sports/backend/srv/cmd/config"
	"sports/backend/srv/controllers/dashboard"
	"sports/backend/srv/routes"
	"sports/backend/srv/server"
	"sports/backend/srv/utils"
)

var cfg config.Config

// initLogger initializes the zap logger with reasonable
// defaults and replaces the global logger.
func initLogger() error {
	// Initialize the logs encoder.
	encoder := zap.NewProductionEncoderConfig()
	encoder.EncodeTime = zapcore.ISO8601TimeEncoder
	encoder.EncodeDuration = zapcore.StringDurationEncoder

	// Initialize the logger.
	logger, err := zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         "console",
		EncoderConfig:    encoder,
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}.Build()
	if err != nil {
		return err
	}

	// Then replace the globals.
	zap.ReplaceGlobals(logger)

	return nil
}

// Load up configuration.
func loadConfiguration() error {
	viper.AddConfigPath("./srv/cmd/config")
	viper.SetConfigName("configuration")

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		return err
	}

	return nil
}

// initialize the database connection and the HTTP router.
func initializeAPI(server *server.Server, driver, username, password, port, host, database string) error {
	var err error

	server.DB, err = utils.GetDBConnection(driver, username, password, port, host, database)
	if err != nil {
		return err
	}

	// Database migration
	server.DB.AutoMigrate(
		&user.User{},
	)

	server.Router = mux.NewRouter()
	routes.InitializeRoutes(server)
	server.HTTPClient = &http.Client{}

	return nil
}

func main() {
	// Global logging synchronizer.
	// This ensures the logged data is flushed out of the buffer before program exits.
	defer zap.S().Sync()

	err := initLogger()
	if err != nil {
		zap.S().Fatal(err)
	}

	err = loadConfiguration()
	if err != nil {
		zap.S().Fatal(err)
	}

	// Set up the dashboard Websocket API module
	dashboard := &dashboard.Dashboard{
		ConnHub: make(map[string]*dashboard.Connection),
		Results: make(chan *dashboard.Result),
		Join:    make(chan *dashboard.Connection),
		Leave:   make(chan *dashboard.Connection),
	}

	srv := server.Server{}
	srv.Addr = cfg.APIAddress
	srv.Dashboard = dashboard

	err = initializeAPI(
		&srv,
		cfg.DBDriver,
		cfg.DBUsername,
		cfg.DBPassword,
		cfg.DBPort,
		cfg.DBHost,
		cfg.DBName,
	)
	if err != nil {
		zap.S().Fatal(err)
	}

	// Disable cert verification to use self-signed certificates for internal service needs.
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	err = run(&srv)
	if err != nil {
		zap.S().Fatal(err)
	}
}

func run(srv *server.Server) error {
	go srv.Dashboard.Run()

	log.Printf("Server listening on %s", srv.Addr)

	zap.S().Fatal(http.ListenAndServeTLS(srv.Addr,
		"./srv/rsa.crt",
		"./srv/rsa.key",
		srv.Router))

	return nil
}
