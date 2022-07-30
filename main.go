package main

import (
	"cdk-example/utils"
	"context"
	"log"
	"os"

	_ "gocloud.dev/postgres/gcppostgres"
	_ "gocloud.dev/runtimevar/gcpsecretmanager"
)

func main() {
	utils.TryLoad(".env.local")
	ctx := context.Background()
	a := Args{}
	utils.Parse(&a, os.Getenv("SECRET_VAULT_URL"))
	err := Run(ctx, a)
	if err != nil {
		log.Fatalf("failed to run: %v", err)
	}
}
