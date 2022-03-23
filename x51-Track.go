//-- p --
package songTrack
import (
	"time"
	"fmt")
type Track struct {
	inst func (clap <-chan []string, flap chan<- []string)
	clap chan []string
	flap chan []string
}

//-- r --

//-- i --
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

func (s *Track) CLAP_Fill (stream []string, waitDrtn uint32) (fillStatus bool) {
	if waitDrtn == 0 {
		s.clap <- stream
		fillStatus = true
		return
	}
	
	chnl := make (chan bool)
	go func () {
		drtn := fmt.Sprintf ("%dms", waitDrtn)
		drtm, _ := time.ParseDuration (drtn)
		time.Sleep (drtm)
		chnl <- true
	} ()
	select {
		case s.clap <- stream: {fillStatus = true}
		case _ = <- chnl: {fillStatus = false}
	}
	return
}

func (s *Track) FLAP_Read (waitDrtn uint32) (readStatus bool, stream []string) {
	if waitDrtn == 0 {
		stream = <- s.flap
		readStatus = true
		return
	}
	
	chnl := make (chan bool)
	go func () {
		drtn := fmt.Sprintf ("%dms", waitDrtn)
		drtm, _ := time.ParseDuration (drtn)
		time.Sleep (drtm)
		chnl <- true
	} ()
	select {
		case stream = <- s.flap: {readStatus = true}
		case _ = <- chnl: {readStatus = false}
	}
	return
}
