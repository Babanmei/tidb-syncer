package mysql

type Result struct {
	Status uint16

	InsertId     uint64
	AffectedRows uint64

	*Resultset
}

type DstExecuter interface {
	DstExecute(query string, args ...interface{}) (*Result, error)
}
