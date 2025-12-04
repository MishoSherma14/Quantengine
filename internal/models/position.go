package models

type Position struct {
	IsOpen     bool
	EntryPrice float64
	Size       float64
	StopLoss   float64
	EntryTime  int64
}
