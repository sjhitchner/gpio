package gpio

import (
	"fmt"
)

func OpenPin(n int, mode Mode) (Pin, error) {
	return nil, fmt.Errorf("OpenPin not implemented on OSX")
}
