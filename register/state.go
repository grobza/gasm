package register

type State struct {
	common []*reg
	ip     int
}

func NewState(regNum int) *State {
	s := &State{
		common: make([]*reg, regNum),
		ip:     0,
	}
	for it := 0; it < regNum; it++ {
		s.common[it] = new(reg)
	}

	return s
}

func (s *State) Common() []*reg {
	return s.common
}

func (s *State) IncIp() {
	s.ip++
}

func (s *State) Ip() int {
	return s.ip
}

type reg struct {
	val interface{}
}

func (r *reg) Set(val interface{}) {
	r.val = val
}

func (r *reg) Get() interface{} {
	return r.val
}
