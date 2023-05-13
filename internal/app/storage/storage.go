package storage

const (
	InsertError         = "error performing insert"
	ReadError           = "error performing read"
	UpdateWithMapError  = "error performing update with map"
	FindError           = "error performing find"
	BeginError          = "error beginning transaction"
	PrepareError        = "error preparing transaction"
	CloseStatementError = "error closing statement"
	ExecError           = "error executing transaction"
	CommitError         = "error committing transaction"
	RollbackError       = "error rolling back transaction"
)

type Storage interface {
	Storage()
}
