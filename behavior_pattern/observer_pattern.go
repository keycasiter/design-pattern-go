package behavior_pattern

import "fmt"

type Observer struct {
	name    string
	subject *Subject
}

func NewObserver(name string, subject *Subject) *Observer {
	return &Observer{name: name, subject: subject}
}

func (o *Observer) Update() {
	fmt.Printf("%s receive change state:%s\n", o.name, o.subject.GetState())
}

type Subject struct {
	observers []*Observer
	state     string
}

func NewSubject() *Subject {
	return &Subject{}
}

func (s *Subject) Attach(observer *Observer) {
	s.observers = append(s.observers, observer)
}

func (s *Subject) GetState() string {
	return s.state
}

func (s *Subject) SetState(state string) {
	s.state = state
	s.notifyAll()
}

func (s *Subject) notifyAll() {
	for _, observer := range s.observers {
		observer.Update()
	}
}
