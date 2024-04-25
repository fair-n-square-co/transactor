[![Go](https://github.com/fair-n-square-co/transactions/actions/workflows/go.yml/badge.svg)](https://github.com/fair-n-square-co/transactions/actions/workflows/go.yml)
[![CodeQL](https://github.com/fair-n-square-co/transactions/actions/workflows/codeql.yml/badge.svg)](https://github.com/fair-n-square-co/transactions/actions/workflows/codeql.yml)


## Database

### DB tools

- [Atlas](https://atlasgo.io/guides/orms/gorm) - database migrations
- [GORM](https://gorm.io/docs/index.html) - ORM library

### DB migrations

1. Install Atlas
```bash
brew install ariga/tap/atlas
```

2. Update model files in `internal/db/models/` directory
   - All our models have this line at the top
    ```go
    package models

    import (
        "github.com/fair-n-square-co/transactions/internal/db/models/base"
    )

    type <Model Name> struct {
        base.PrimaryKey      // Always add this line for UUID as primary key
        base.DateTime        // Add this line if you want to use CreatedAt, UpdatedAt
        base.SoftDeleteModel // Add this line if you want to use soft delete
        // Add your fields
    }
    ```

3. Add the model to the `models` array in `internal/db/models/models.go` file
```go
var models = []any{
    // ... other models
    &<Model Name>{},
}
```
4. Run the following command to generate the migration file
```sh
atlas migrate diff <migration name> --env dev
```
5. Run the following command to apply the migration
```sh
atlas migrate up --env dev
```


## Tests

### Gomock

We use gomock to generate mock interfaces for testing. To generate a mock interface for a given interface, run the following command:
```sh
make gen/mock
```

### Installation

Install gomock
```sh
go install go.uber.org/mock/mockgen@latest
```

### Usage

Add your generate mockgen command at the top of your go file like:
```go
//go:generate mockgen -source=<file name> -destination=mocks/<file name> -package=mocks
```

