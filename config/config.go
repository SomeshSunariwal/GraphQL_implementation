package config

const (
	PORT = "8080"

	HOST = "localhost"

	PG_PORT = 5432

	PG_USER = "postgres"

	PG_PASSWORD = "postgres"

	PG_DATABASE_NAME = "test"
)

type ENV_CONFIG struct {
	HOST             string
	PG_USER          string
	PG_PASSWORD      string
	PG_DATABASE_NAME string
}
