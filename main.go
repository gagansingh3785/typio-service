package main

import (
	"context"
	"log"
	"os"

	"github.com/gagansingh3785/typio-service/server"
	cli "github.com/urfave/cli/v3"
)

func main() {
	log.Println("[MAIN] Starting typio service...")

	if err := setupTypioService(); err != nil {
		log.Fatalf("[MAIN] Failed to setup typio service: %v", err)
	}

	cliApp := &cli.Command{
		Name:        "typio-service",
		Description: "Backend service for typio project",
		Version:     "1.0.0",
		Commands: []*cli.Command{
			{
				Name:        "start-server",
				Description: "Start the typio service http server",
				Usage:       "typio-service start-server",
				Action: func(c context.Context, cmd *cli.Command) error {
					return server.StartHTTPServer()
				},
			},
			{
				Name:        "migrations:create",
				Description: "Create a new database migration",
				Usage:       "typio-service migrations:create",
				Action: func(c context.Context, cmd *cli.Command) error {
					log.Println("[MAIN] Creating database migration...")
					return nil
				},
			},
			{
				Name:        "migrations:run",
				Description: "Run database migrations",
				Usage:       "typio-service migrations:run",
				Action: func(c context.Context, cmd *cli.Command) error {
					log.Println("[MAIN] Running database migrations...")
					return nil
				},
			},
			{
				Name:        "migrations:rollback",
				Description: "Rollback the last database migration",
				Usage:       "typio-service migrations:rollback",
				Action: func(c context.Context, cmd *cli.Command) error {
					log.Println("[MAIN] Rolling back the last database migration...")
					return nil
				},
			},
		},
	}

	if err := cliApp.Run(context.Background(), os.Args); err != nil {
		log.Fatalf("[MAIN] Error running typio service: %v", err)
	}
}
