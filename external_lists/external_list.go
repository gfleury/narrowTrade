package external_lists

type ExternalList interface {
	GetSymbols() ([]string, error)
}
