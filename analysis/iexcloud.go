package analysis

import (
	"fmt"

	"github.com/gfleury/narrowTrade/models"
	iex "github.com/goinvest/iexcloud/v2"
)

type IEXAnalyser struct {
	Client *iex.Client
}

func (a *IEXAnalyser) Analysis() []InstrumentAnalysis {
	return []InstrumentAnalysis{}
}

func (a *IEXAnalyser) Indicator(models.Instrument, IndicatorName) ([]float64, error) {
	return nil, fmt.Errorf("Not implemented yet")
}

func (a *IEXAnalyser) Quote(models.Instrument) (*Quote, error) {
	return nil, fmt.Errorf("Not implemented yet")
}

func (a *IEXAnalyser) OneAnalysis(models.Instrument) (*InstrumentAnalysis, error) {
	return nil, fmt.Errorf("Not implemented yet")
}

func (a *IEXAnalyser) Analyse([]models.Instrument) error {
	return fmt.Errorf("Not implemented yet")
}
