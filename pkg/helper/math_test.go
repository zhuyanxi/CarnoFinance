package helper

import (
	"strconv"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestRSRS(t *testing.T) {
	data := []float64{1, 1.01, 1.02, 1.03, 1.04, 1.05, 1.06}
	// go ver: 10.332265969835445
	// py ver: 10.332265969835467
	result := RSRS(data)
	logrus.Infoln(result)
	res := strconv.FormatFloat(result, 'f', 10, 64)
	if res != "10.3322659698" {
		t.Error("RSRS failed")
	}
}
