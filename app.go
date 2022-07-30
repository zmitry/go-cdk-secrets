package main

import (
	"cdk-example/services/hello"
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/starius/api2"
	"gocloud.dev/postgres"
	_ "gocloud.dev/secrets/gcpkms"
)

type Args struct {
	// key where secret is stored
	PostgresURLKey string `env:"POSTGRES_URL_KEY" long:"postgres-url-key" default:"postgres-url"`
	// actual url for current service
	PostgresURL string `env:"POSTGRES_URL" long:"postgres-url" secret-key:"PostgresURLKey" optional:"true"`
}

func Run(ctx context.Context, args Args) error {
	fmt.Println("args, be aware they are sensitive", args)
	// Replace this with your actual settings.
	db, err := postgres.Open(ctx, args.PostgresURL)
	if err != nil {
		return err
	}
	defer db.Close()
	log.Printf("creating table")

	// Create the table to make sure connection works
	_, err = db.Exec("CREATE TABLE if not exists messages (message text);")
	if err != nil {
		return fmt.Errorf("failed to create table: %v", err)
	}
	log.Printf("setting up services")
	service := hello.NewEchoService(hello.NewEchoRepository(db))
	routes := hello.GetRoutes(service)
	api2.BindRoutes(http.DefaultServeMux, routes)
	log.Printf("starting server on 127.0.0.1:8080")
	return http.ListenAndServe(":8080", nil)
}
