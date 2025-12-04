package generator

type StrategyTemplate struct {
	Entry        IndicatorSpec
	Confirmation IndicatorSpec
	Baseline     IndicatorSpec
	Volume       IndicatorSpec
	Exit         IndicatorSpec
}

type IndicatorSpec struct {
	Name   string
	Params map[string]any
}
