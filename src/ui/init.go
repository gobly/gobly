package ui

import "core"

func CreateContext(prefix string, r *core.Router) {
	r.AddSubRouter(prefix, initialize)
}

func initialize(s *core.Router) {
	s.AddStatic("/static", "static")
}
