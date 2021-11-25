package core

type EngineI interface {
	Sync()
}

type RepositoryI interface {
	GetSyncAdapters() []SyncAdapterI
	AddSyncAdapter(adapterI SyncAdapterI)
	Flush() error
	SetBlock(blockNum int64)
}