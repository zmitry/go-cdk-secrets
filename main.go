package main

import (
	"cdk-example/utils"
	"context"
	"log"

	_ "gocloud.dev/postgres/gcppostgres"
	_ "gocloud.dev/runtimevar/gcpsecretmanager"
)

func main() {
	ctx := context.Background()
	a := Args{}
	utils.Parse(&a, "gcpsecretmanager://projects/es-scalability-test/secrets")
	err := Run(ctx, a)
	if err != nil {
		log.Fatalf("failed to run: %v", err)
	}
}
