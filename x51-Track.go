//-- p --
package songTrack

//-- i --
type Track struct {
	inst func (clap <-chan []string, flap chan<- []string)
	clap chan []string
	flap chan []string
}
func Track_Create (instrc func (clap <-chan []string, flap chan<- []string)) (*Track) {
	return &Track {
		inst: instrc,
		clap: make (chan []string),
		flap: make (chan []string),
	}
}
func (s *Track) Run () {
	go s.inst (s.clap, s.flap)
}
func (s *Track) _CLAP_Fill (stream []string) () {
	s.clap <- stream
}
func (s *Track) _FLAP_Read () (stream []string) {
	select {
		case stream = <- s.flap: {}
		default: {}
	}
	return
}
