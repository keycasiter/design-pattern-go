package behavior_pattern

type Memento struct {
	state string
}

func NewMemento(state string) *Memento {
	return &Memento{state: state}
}

func (m *Memento) GetState() string {
	return m.state
}

type StateMachine struct {
	mementos []*Memento
}

func NewStateMachine() *StateMachine {
	return &StateMachine{}
}

func (s *StateMachine) Add(m *Memento) {
	s.mementos = append(s.mementos, m)
}

func (s *StateMachine) GetMemento(state string) (*Memento, int) {
	for idx, m := range s.mementos {
		if m.GetState() == state {
			return m, idx
		}
	}
	return nil, -1
}
