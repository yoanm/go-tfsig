package tfsig_test

import (
	"testing"

	"github.com/yoanm/go-tfsig"
)

func TestToTerraformIdentifier(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		Value string
		Want  string
	}{
		"Already valid":          {"a_valid-identifier", "a_valid-identifier"},
		"Contains special chars": {"id_@&é\"'(§è!çà)^$`ù=:;,<#°*¨£%+/.?>|\\", "id_----------------------------------"},
		"Start with a number":    {"0-id", "_-id"},
	}

	for tcname, tcase := range cases {
		tcase := tcase // For parallel execution

		t.Run(
			tcname,
			func(t *testing.T) {
				t.Parallel()

				got := tfsig.ToTerraformIdentifier(tcase.Value)
				if got != tcase.Want {
					t.Errorf("wrong result for case %q: got %v, want %v", t.Name(), got, tcase.Want)
				}
			},
		)
	}
}
