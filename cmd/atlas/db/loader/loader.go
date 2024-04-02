package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	_ "ariga.io/atlas-go-sdk/recordriver"
	"ariga.io/atlas-provider-gorm/gormschema"
	"github.com/fair-n-square-co/transactions/internal/db/models"
	"github.com/fair-n-square-co/transactions/internal/db/models/config"
)

const (
	dialect = "postgres"
)

func createEnums() string {
	var stmt string
	enums := map[string][]string{
		"transaction_type":      {"payment", "settlement"},
		"Transaction_user_type": {"payer", "payee"},
	}
	for enum, values := range enums {
		stmt := fmt.Sprintf("CREATE TYPE %s AS ENUM ('%s');\n", enum, strings.Join(values, "', '"))
		io.WriteString(os.Stdout, stmt)
	}
	return stmt
}

func load() {
	stmts, err := gormschema.New(
		dialect, gormschema.WithConfig(config.GetGormConfig()),
	).Load(models.GetAllModels()...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load gorm schema: %v\n", err)
		os.Exit(1)
	}
	io.WriteString(os.Stdout, createEnums())
	io.WriteString(os.Stdout, stmts)
}
