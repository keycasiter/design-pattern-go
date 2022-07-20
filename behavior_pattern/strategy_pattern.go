package behavior_pattern

type Strategy interface {
	Handle(x, y int) int
}

type StrategyContext struct {
	strategy Strategy
}

func NewStrategyContext(strategy Strategy) *StrategyContext {
	return &StrategyContext{strategy: strategy}
}

func (s *StrategyContext) Run(x, y int) int {
	return s.strategy.Handle(x, y)
}

type SumStrategy struct {
}

func NewSumStrategy() *SumStrategy {
	return &SumStrategy{}
}

func (s *SumStrategy) Handle(x, y int) int {
	return x + y
}

type MultiplyStrategy struct {
}

func NewMultiplyStrategy() *MultiplyStrategy {
	return &MultiplyStrategy{}
}

func (s *MultiplyStrategy) Handle(x, y int) int {
	return x * y
}
