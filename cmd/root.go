package cmd

import (
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/blendle/zapdriver"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	configFile string
	config     Config
)

var rootCommand = &cobra.Command{
	Use: "udhaar",
}

func init() {
	cobra.OnInitialize(func() {
		if len(configFile) > 0 {
			viper.SetConfigFile(configFile)
		} else {
			viper.AddConfigPath(".")
			viper.SetConfigName("config")
		}
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
		viper.SetEnvPrefix("UDHAAR")
		viper.AutomaticEnv()
		if err := viper.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
				log.Fatalf("failed to read config file: %v", err)
			}
		}
		if err := viper.Unmarshal(&config); err != nil {
			log.Fatal(err)
		}
	})

	rootCommand.AddCommand(
		serveCommand(),
		migrateCommand(),
	)

	flags := rootCommand.PersistentFlags()
	flags.StringVarP(&configFile, "config", "c", "", "config file path")

	flags.Bool("dev", false, "development mode")
	bindPFlag(flags, "dev")
	flags.Bool("pprof", false, "expose pprof http interface")
	bindPFlag(flags, "pprof")
}

func Execute() error {
	return rootCommand.Execute()
}

func getLogger() (logger *zap.Logger) {
	if config.DevMode {
		return getCLILogger()
	}
	cfg := zap.Config{
		Level:            zap.NewAtomicLevelAt(zapcore.InfoLevel),
		Encoding:         "json",
		EncoderConfig:    zapdriver.NewProductionEncoderConfig(),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
	logger, _ = cfg.Build(zapdriver.WrapCore(zapdriver.ServiceName("udhaar")))
	return
}

func getCLILogger() (logger *zap.Logger) {
	level := zap.NewAtomicLevel()
	if config.DevMode {
		level = zap.NewAtomicLevelAt(zap.DebugLevel)
	}
	cfg := zap.Config{
		Level:       level,
		Development: config.DevMode,
		Encoding:    "console",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "T",
			LevelKey:       "L",
			NameKey:        "N",
			CallerKey:      "C",
			MessageKey:     "M",
			StacktraceKey:  "S",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.CapitalColorLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
	logger, _ = cfg.Build()
	return
}

func bindPFlag(flags *pflag.FlagSet, key string, flag ...string) {
	if len(flag) == 0 {
		flag = []string{key}
	}
	if err := viper.BindPFlag(key, flags.Lookup(flag[0])); err != nil {
		panic(err)
	}
}

func waitSIGINT() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	signal.Stop(quit)
	close(quit)
	for range quit {
		continue
	}
}
