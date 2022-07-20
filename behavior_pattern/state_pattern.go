package behavior_pattern

import "fmt"

type State interface {
	Do() string
	GetState() string
}

type StateContext struct {
	state State
}

func NewStateContext(state State) *StateContext {
	return &StateContext{state: state}
}

func (ctx *StateContext) Run() {
	fmt.Printf("state:%s is %s \n", ctx.state.GetState(), ctx.state.Do())
}

type StartState struct {
	state string
}

func NewStartState() *StartState {
	return &StartState{state: "start"}
}

func (s *StartState) Do() string {
	return "starting"
}

func (s *StartState) GetState() string {
	return s.state
}

type EndState struct {
	state string
}

func NewEndState() *EndState {
	return &EndState{state: "end"}
}

func (s *EndState) Do() string {
	return "ending"
}

func (s *EndState) GetState() string {
	return s.state
}
