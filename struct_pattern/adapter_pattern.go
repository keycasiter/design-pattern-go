package struct_pattern

type EngineAction interface {
	Run() string
}

type FastEngine struct {
}

func (e *FastEngine) Run() string {
	return "fast engine run..."
}

type SlowEngine struct {
}

func (e *SlowEngine) Run() string {
	return "slow engine run..."
}

type EngineMode int

const (
	Fast_Engine = iota
	Slow_Engine
)

//适配器
type EngineAdapter struct {
	engine interface{}
}

func NewEngineAdapter(engineMode EngineMode) *EngineAdapter {
	adapter := &EngineAdapter{}
	switch engineMode {
	case Fast_Engine:
		adapter.engine = &FastEngine{}
	case Slow_Engine:
		adapter.engine = &SlowEngine{}
	default:
		panic("unsupport engine mode")
	}
	return adapter
}

func (e *EngineAdapter) Run() string {
	switch e.engine.(type) {
	case *FastEngine:
		return e.engine.(*FastEngine).Run()
	case *SlowEngine:
		return e.engine.(*SlowEngine).Run()
	default:
		panic("unsupport engine")
	}
}
