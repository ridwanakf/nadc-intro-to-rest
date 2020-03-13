package app

import (
	"context"
	"database/sql"
	"os"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/ridwanakf/nadc-intro-to-rest/constant"
	"github.com/ridwanakf/nadc-intro-to-rest/internal/config"
	"gopkg.in/yaml.v2"
)

type NadcRest struct {
	Repos    *Repos
	UseCases *Usecases
}

func NewNadcRest() (*NadcRest, error) {
	cfg, err := readConfig(constant.ConfigProjectFilepath)
	if err != nil {
		return nil, errors.Wrapf(err, "error getting config")
	}

	db, err := initDB(cfg)
	if err != nil {
		return nil, errors.Wrapf(err, "error connect db")
	}

	app := new(NadcRest)

	app.Repos, err = newRepos(db)
	if err != nil {
		return nil, errors.Wrap(err, "errors invoking newRepos")
	}

	app.UseCases = newUsecases(app.Repos)

	return app, nil
}

func (a *NadcRest) Close() []error {
	var errs []error

	errs = append(errs, a.Repos.Close()...)
	errs = append(errs, a.UseCases.Close()...)

	return errs
}

func readConfig(cfgPath string) (*config.Config, error) {
	f, err := os.Open(cfgPath)
	if err != nil {
		return nil, errors.Wrapf(err, "config file not found")
	}
	defer f.Close()

	var cfg config.Config

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		return nil, errors.Wrapf(err, "error reading config")
	}

	return &cfg, nil
}

func initDB(cfg *config.Config) (*sql.DB, error) {

	// Initialize SQL DB
	db, err := sql.Open(cfg.DB.Driver, cfg.DB.Address)
	if err != nil {
		return nil, err
	}

	// Check if db connected
	if err = db.PingContext(context.Background()); err != nil {
		return nil, err
	}

	// Set db params
	db.SetMaxIdleConns(cfg.DB.MaxConns)
	db.SetMaxOpenConns(cfg.DB.MaxIdleConns)

	return db, nil
}
