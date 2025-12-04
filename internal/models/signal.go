package models

type Signal int

const (
	SignalNone Signal = iota
	SignalBuy
	SignalSell
	SignalExit
)
