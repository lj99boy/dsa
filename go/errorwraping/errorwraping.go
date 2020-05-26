package errorwraping

import (
	"errors"
	"fmt"
)

func ErrorWrap() {
	e := errors.New("test e1")
	newE := fmt.Errorf("wraping e1 %w", e)

}
