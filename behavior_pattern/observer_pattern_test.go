package behavior_pattern

import "testing"

func TestObserverPattern(t *testing.T) {
	subject := NewSubject()
	observer1 := NewObserver("#1", subject)
	observer2 := NewObserver("#2", subject)
	observer3 := NewObserver("#3", subject)
	observer4 := NewObserver("#4", subject)

	subject.Attach(observer1)
	subject.Attach(observer2)
	subject.Attach(observer3)
	subject.Attach(observer4)

	subject.SetState("start")
}
