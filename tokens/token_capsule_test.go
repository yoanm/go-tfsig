package tokens_test

import (
	"testing"

	"github.com/yoanm/go-tfsig/tokens"
)

func TestContainsCapsule_nil(t *testing.T) {
	t.Parallel()

	if tokens.ContainsCapsule(nil) != false {
		t.Errorf("expected false, got  true")
	}
}
