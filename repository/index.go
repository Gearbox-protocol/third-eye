package repository

import (
	"gorm.io/gorm"
	"github.com/Gearbox-protocol/gearscan/models/sync_adapter"
)
type Repository struct {
	db *gorm.DB
	syncAdapters []*sync_adapter.SyncAdapter
}
type RepositoryI interface {
	GetSyncAdapters() []*sync_adapter.SyncAdapter
	AddSyncAdapters(adapter *sync_adapter.SyncAdapter)
}

func NewRepository(db *gorm.DB) RepositoryI {
	return &Repository{
		db: db,
	}
}

func (repo *Repository) GetSyncAdapters() []*sync_adapter.SyncAdapter {
	return repo.syncAdapters
}

func (repo *Repository) AddSyncAdapters(adapter *sync_adapter.SyncAdapter) {
	repo.syncAdapters = append(repo.syncAdapters, adapter)
}