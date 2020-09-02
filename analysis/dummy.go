package analysis

import (
	"fmt"

	"github.com/gfleury/narrowTrade/models"
)

type DummyAnalyser struct {
}

func (a *DummyAnalyser) Init() {
}

func (a *DummyAnalyser) Analysis() []InstrumentAnalysis {
	return nil
}

func (a *DummyAnalyser) Indicator(i models.Instrument, indicatorName IndicatorName) ([]float64, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (a *DummyAnalyser) Quote(i models.Instrument) (*Quote, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (a *DummyAnalyser) OneAnalysis(i models.Instrument) (*InstrumentAnalysis, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (a *DummyAnalyser) Analyse(is []models.Instrument) error {
	return fmt.Errorf("Not implemented")
}
