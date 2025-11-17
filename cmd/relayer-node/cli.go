package main

import (
	"context"
	"strconv"

	"github.com/ethereum/go-ethereum/log"
	"github.com/urfave/cli/v2"

	"github.com/Sandwichzzy/relayer-node/common/cliapp"
	"github.com/Sandwichzzy/relayer-node/common/opio"
	"github.com/Sandwichzzy/relayer-node/config"
	"github.com/Sandwichzzy/relayer-node/database"
	"github.com/Sandwichzzy/relayer-node/database/create_table"
	"github.com/Sandwichzzy/relayer-node/service"
)

var (
	ConfigFlag = &cli.StringFlag{
		Name:    "config",
		Value:   "./relayer-node.yaml",
		Aliases: []string{"c"},
		Usage:   "path to config file",
		EnvVars: []string{"RELAYER_NODE_CONFIG"},
	}
	MigrationsFlag = &cli.StringFlag{
		Name:    "migrations-dir",
		Value:   "./migrations",
		Usage:   "path to migrations folder",
		EnvVars: []string{"RELAYER_NODE_MIGRATIONS_DIR"},
	}
)

func runIndexer(ctx *cli.Context, shutdown context.CancelCauseFunc) (cliapp.Lifecycle, error) {
	return nil, nil
}

func runApi(ctx *cli.Context, _ context.CancelCauseFunc) (cliapp.Lifecycle, error) {
	log.Info("running api ...")
	cfg, err := config.New(ctx.String(ConfigFlag.Name))
	if err != nil {
		log.Error("failed to load config", "err", err)
		return nil, err
	}
	return service.NewAPI(ctx.Context, cfg)
}

func runMigrations(ctx *cli.Context) error {
	ctx.Context = opio.CancelOnInterrupt(ctx.Context)
	log.Info("running migrations...")
	cfg, err := config.New(ctx.String(ConfigFlag.Name))
	if err != nil {
		log.Error("failed to load config", "err", err)
		return err
	}
	db, err := database.NewDB(ctx.Context, cfg.MasterDb)
	if err != nil {
		log.Error("failed to connect to database", "err", err)
		return err
	}
	defer func(db *database.DB) {
		err := db.Close()
		if err != nil {
			return
		}
	}(db)
	err = db.ExecuteSQLMigration(ctx.String(MigrationsFlag.Name))
	if err != nil {
		return err
	}
	for i := range cfg.RPCs {
		log.Info("create chain table by chainId", "chainId", cfg.RPCs[i].ChainId)
		create_table.CreateTableFromTemplate(strconv.Itoa(int(cfg.RPCs[i].ChainId)), db)
	}
	log.Info("running migrations and create table from template success")
	return nil
}

func newCli(GitCommit string, GitDate string) *cli.App {
	flags := []cli.Flag{ConfigFlag}
	migrationFlags := []cli.Flag{MigrationsFlag, ConfigFlag}
	return &cli.App{
		Version:              "v0.0.1",
		Description:          "An indexer bridge events with a serving api layer",
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			{
				Name:        "api",
				Flags:       flags,
				Description: "Runs the api service",
				Action:      cliapp.LifecycleCmd(runApi),
			},
			{
				Name:        "index",
				Flags:       flags,
				Description: "Runs the indexing service",
				Action:      cliapp.LifecycleCmd(runIndexer),
			},
			{
				Name:        "migrate",
				Flags:       migrationFlags,
				Description: "Runs the database migrations",
				Action:      runMigrations,
			},
			{
				Name:        "version",
				Description: "print version",
				Action: func(ctx *cli.Context) error {
					cli.ShowVersion(ctx)
					return nil
				},
			},
		},
	}
}
