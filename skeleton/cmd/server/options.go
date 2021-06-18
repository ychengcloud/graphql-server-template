package main

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	server "{{ .Extra.pkgpath }}/internal"
	"{{ .Extra.pkgpath }}/internal/handlers"
	"{{ .Extra.pkgpath }}/internal/middleware"
	"{{ .Extra.pkgpath }}/internal/models"
	"{{ .Extra.pkgpath }}/internal/database"
	"{{ .Extra.pkgpath }}/internal/models"
	"{{ .Extra.pkgpath }}/pkg/log"
	"{{ .Extra.pkgpath }}/pkg/transports/http"
)

func NewHttpOptions(v *viper.Viper) (*http.Options, error) {
	var (
		err error
		o   = new(http.Options)
	)

	if err = v.UnmarshalKey("http", o); err != nil {
		return nil, err
	}

	return o, err
}

func NewServerOptions(v *viper.Viper, logger *zap.Logger) (*server.Options, error) {
	var err error
	o := new(server.Options)
	if err = v.UnmarshalKey("app", o); err != nil {
		return nil, errors.Wrap(err, "unmarshal app option error")
	}

	logger.Info("load application options success")

	return o, err
}

func NewHandlersOptions(v *viper.Viper, logger *zap.Logger) (*handlers.Options, error) {
	var err error
	o := new(handlers.Options)
	if err = v.UnmarshalKey("http", o); err != nil {
		return nil, errors.Wrap(err, "unmarshal app option error")
	}

	logger.Info("load application options success")

	return o, err
}

func NewLogOptions(v *viper.Viper) (*log.Options, error) {
	var (
		err error
		o   = new(log.Options)
	)

	if err = v.UnmarshalKey("log", o); err != nil {
		return nil, err
	}
	fmt.Println("new log:", o.Filename)
	return o, err
}

func NewDatabaseOptions(v *viper.Viper, logger *zap.Logger) (*database.Options, error) {
	var err error
	o := &database.Options{
		MysqlOptions:  &database.MysqlOptions{},
		SqliteOptions: &database.SqliteOptions{},
	}
	if err = v.UnmarshalKey("db", o); err != nil {
		return nil, errors.Wrap(err, "unmarshal db option error")
	}

	logger.Info("load database options success",
		zap.String("User", o.MysqlOptions.User),
		zap.String("Host", o.MysqlOptions.Host),
		zap.Int("Port", o.MysqlOptions.Port),
	)
	return o, err
}

func NewSerivceOptions(v *viper.Viper) (*models.ServiceOptions, error) {
	var (
		err error
		o   = new(models.ServiceOptions)
	)

	if err = v.UnmarshalKey("services", o); err != nil {
		return nil, err
	}
	fmt.Println("new services:", o.KratosHost)
	return o, err
}

func NewAuthOptions(v *viper.Viper) (*middleware.Options, error) {
	var (
		err error
		o   = new(middleware.Options)
	)

	if err = v.UnmarshalKey("auth", o); err != nil {
		return nil, err
	}
	fmt.Println("new auth options:", o)
	return o, err
}

func NewAuth(authOpts *middleware.Options, opts *models.ServiceOptions) (middleware.Auth, error) {
	if authOpts.Skip {
		return nil, nil
	}
	
	switch authOpts.AuthType {
	case "kratos":
		return middleware.NewAuthKratos(opts), nil
	default:
		return nil, fmt.Errorf("Unknown auth type")
	}
}
