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
