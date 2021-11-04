package opus

import (
	"github.com/pion/mediadevices/pkg/codec"
	"testing"
)

func TestShouldImplementBitRateControl(t *testing.T) {
	e := &encoder{}
	if _, ok := e.Controller().(codec.BitRateController); !ok {
		t.Error()
	}
}

func TestShouldImplementKeyFrameControl(t *testing.T) {
	t.SkipNow() // TODO: Implement key frame control

	e := &encoder{}
	if _, ok := e.Controller().(codec.KeyFrameController); !ok {
		t.Error()
	}
}
