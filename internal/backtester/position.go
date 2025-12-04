package backtester

type Position struct {
	IsOpen     bool
	IsLong     bool
	EntryPrice float64
	Size       float64
	StopLoss   float64
	Target1    float64

	PartialClosed bool
	TrailingOn    bool
	TrailingStop  float64
}
