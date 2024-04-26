data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "./cmd/atlas/db/loader",
  ]
}

env "dev" {
  src = data.external_schema.gorm.url
  dev = "docker://postgres/16/dev?search_path=public"
  url = "postgres://postgres:postgres@localhost:5432/transactions?sslmode=disable"
  migration {
    dir = "file://db/migrations"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}

env "fly" {
  url = "postgres://postgres:<password>@localhost:5454/transactions?sslmode=disable"
  migration {
    dir = "file://db/migrations"
  }
}
