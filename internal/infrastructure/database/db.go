package database

import (
	"errors"
	"github.com/jmoiron/sqlx"
)

type DB struct {
	master     *sqlx.DB
	syncSlave  *sqlx.DB
	asyncSlave *sqlx.DB
}

type ConnectionSpec struct {
	Master     *sqlx.DB
	SyncSlave  *sqlx.DB
	AsyncSlave *sqlx.DB
}

func NewDatabase(spec *ConnectionSpec) (*DB, error) {
	if spec.Master == nil {
		return nil, errors.New("master must be defined")
	}

	result := &DB{
		master: spec.Master,
	}

	if spec.SyncSlave != nil {
		result.syncSlave = spec.SyncSlave
	}

	if spec.AsyncSlave != nil {
		result.asyncSlave = spec.AsyncSlave
	}

	return result, nil
}

func (d *DB) GetMaster() *sqlx.DB {
	return d.master
}

func (d *DB) GetSyncSlave() *sqlx.DB {
	return d.syncSlave
}

func (d *DB) GetAsyncSlave() *sqlx.DB {
	return d.asyncSlave
}
