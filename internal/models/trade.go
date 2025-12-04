package models

type Trade struct {
	EntryPrice   float64
	ExitPrice    float64
	Profit       float64
	RMultiple    float64
	TimestampIn  int64
	TimestampOut int64
}
