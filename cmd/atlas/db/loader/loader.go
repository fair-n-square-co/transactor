package main

import (
	"fmt"
	"io"
	"os"

	_ "ariga.io/atlas-go-sdk/recordriver"
	"ariga.io/atlas-provider-gorm/gormschema"
	"github.com/fair-n-square-co/transactions/internal/db/models"
	"github.com/fair-n-square-co/transactions/internal/db/models/config"
)

const (
	dialect = "postgres"
)

func load() {
	stmts, err := gormschema.New(
		dialect, gormschema.WithConfig(config.GetGormConfig()),
	).Load(models.GetAllModels()...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load gorm schema: %v\n", err)
		os.Exit(1)
	}
	io.WriteString(os.Stdout, stmts)
}
