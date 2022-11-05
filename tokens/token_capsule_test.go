package tokens

import (
	"testing"
)

func TestContainsCapsule_nil(t *testing.T) {
	if ContainsCapsule(nil) != false {
		t.Errorf("expected false, got  true")
	}
}
