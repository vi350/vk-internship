package registry

type Registry interface {
	Sync()
	RemoveInactive(int)
}
