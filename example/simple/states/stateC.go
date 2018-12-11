package states

import (
	"fmt"
	"github.com/Enflick/gohsm"
)

type StateC struct {
	parentState *StateA
	entered     bool
	exited      bool
}

func NewStateC(srv hsm.Service, parentState *StateA) *StateC {
	hsm.Precondition(srv, parentState != nil, fmt.Sprintf("NewStateC: parentState cannot be nil"))

	return &StateC{
		parentState: parentState,
	}
}

func (s *StateC) Name() string {
	return "C"
}

func (s *StateC) OnEnter(srv hsm.Service, event hsm.Event) hsm.State {
	hsm.Precondition(srv, !s.entered, fmt.Sprintf("State %s has already been entered", s.Name()))
	srv.Logger().Debug("->C;")
	s.entered = true
	return s
}

func (s *StateC) OnExit(srv hsm.Service, event hsm.Event) hsm.State {
	hsm.Precondition(srv, !s.exited, fmt.Sprintf("State %s has already been entered", s.Name()))
	srv.Logger().Debug("<-C;")
	s.exited = true
	return s.ParentState()
}

func (s *StateC) EventHandler(srv hsm.Service, event hsm.Event) hsm.Transition {
	switch event.ID() {
	case ex.ID():
		return hsm.NewExternalTransition(event, NewStateC(srv, s.parentState), action6)
	case ey.ID():
		return hsm.NewInternalTransition(event, action7)
	default:
		return hsm.NilTransition
	}
}

func (s *StateC) Entered() bool {
	return s.entered
}

func (s *StateC) Exited() bool {
	return s.exited
}

func (s *StateC) ParentState() hsm.State {
	return s.parentState
}

func action6(srv hsm.Service) {
	srv.Logger().Debug("Action6")
	LastActionIdExecuted = 6
}

func action7(srv hsm.Service) {
	srv.Logger().Debug("Action7")
	LastActionIdExecuted = 7
}
