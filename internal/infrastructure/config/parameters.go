package config

const (
	DatabaseDriver          = "DATABASE_DRIVER"
	DatabaseHost            = "DATABASE_HOST"
	DatabaseName            = "DATABASE_NAME"
	DatabaseUser            = "DATABASE_USER"
	DatabasePassword        = "DATABASE_PASSWORD"
	DatabasePortInContainer = "DATABASE_PORT_CONTAINER"

	MasterDatabaseDSN     = "MASTER_DATABASE_DSN"
	SlaveSyncDatabaseDSN  = "SYNC_SLAVE_DATABASE_DSN"
	SlaveAsyncDatabaseDSN = "ASYNC_SLAVE_DATABASE_DSN"
)
