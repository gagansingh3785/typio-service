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

	// setup application context
	err := setupTypingService()
	if err != nil {
		log.Fatalf("[MAIN] Failed to setup typing service: %v", err)
	}

	cliApp := &cli.Command{
		Name:        "typio-service",
		Description: "Backend service for typio project",
		Version:     "1.0.0",
		Commands: []*cli.Command{
			{
				Name:        "start-server",
				Description: "Start the typio service http server",
				Usage:       "typing-service start-server",
				Action: func(c context.Context, cmd *cli.Command) error {
					return server.StartHTTPServer()
				},
			},
			{
				Name:        "migrations:run",
				Description: "Run database migrations",
				Usage:       "typing-service migrations:run",
				Action: func(c context.Context, cmd *cli.Command) error {
					log.Println("[MAIN] Running database migrations...")
					return nil
				},
			},
			{
				Name:        "migrations:rollback",
				Description: "Rollback the last database migration",
				Usage:       "typing-service migrations:rollback",
				Action: func(c context.Context, cmd *cli.Command) error {
					log.Println("[MAIN] Rolling back the last database migration...")
					return nil
				},
			},
		},
	}
	err = cliApp.Run(context.Background(), os.Args)
	if err != nil {
		log.Fatalf("[MAIN] Error running typio service: %v", err)
	}
}
