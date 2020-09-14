package trader

type AvailableCash struct {
	AvailableCash float64
}

func (ac *AvailableCash) GetSplit(n int) float64 {
	return ac.AvailableCash / float64(n)
}

func (ac *AvailableCash) GetTotal() float64 {
	return ac.AvailableCash
}

func (ac *AvailableCash) SetAvailableCash(available float64) {
	ac.AvailableCash = available
}

func (ac *AvailableCash) Spent(spent float64) {
	ac.AvailableCash -= spent
}
