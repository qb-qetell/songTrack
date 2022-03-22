package songTrack

import (
	"fmt"
	"testing")
 
func TestPremier (t *testing.T) {
	t01 := Track_Create (func (clap <-chan []string, flap chan<- []string) {
		_x5100 := <- clap
		fmt.Printf ("Message from ST: \n%s\n%s\n%s\n", _x5100 [0], _x5100 [1],
			_x5100 [2])
		_x5200 := <- clap
		fmt.Printf ("Message from ST: \n%s\n%s\n%s\n", _x5200 [0], _x5200 [1],
			_x5200 [2])

		flap <- []string {"a", "b", "c"}
		flap <- []string {"x", "y", "z"}
	})

	t01.Run ()

	t01._CLAP_Fill ([]string {"1", "2", "3"})
	t01._CLAP_Fill ([]string {"7", "8", "9"})

	
   	_x5300 := []string {}
   	for {
   		_x5300 = t01._FLAP_Read ()
   		if _x5300 == nil {
   			continue
   		}
   		break
   	}
   	fmt.Printf ("Message from FT: \n%s\n%s\n%s\n", _x5300 [0], _x5300 [1], _x5300 [2])
   	_x5400 := []string {}
   	for {
   		_x5400 = t01._FLAP_Read ()
   		if _x5400 == nil {
   			continue
   		}
   		break
   	}
   	fmt.Printf ("Message from FT: \n%s\n%s\n%s\n", _x5400 [0], _x5400 [1], _x5400 [2])

}
