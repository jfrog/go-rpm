package rpm

import (
	"io/ioutil"
	"strings"
	"testing"
)

// TestTestdataModuleExcludesFixtures guards the nested module that keeps the RPM
// fixtures out of the published module zip. The fixtures themselves are still
// read by relative path.
func TestTestdataModuleExcludesFixtures(t *testing.T) {
	b, err := ioutil.ReadFile("testdata/go.mod")
	if err != nil {
		t.Fatalf("testdata/go.mod: %v", err)
	}

	var modulePath, goLine string
	for _, line := range strings.Split(string(b), "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "//") {
			continue
		}
		switch {
		case strings.HasPrefix(line, "module "):
			modulePath = strings.TrimSpace(strings.TrimPrefix(line, "module "))
		case strings.HasPrefix(line, "go "):
			goLine = line
		case strings.HasPrefix(line, "require"):
			t.Fatalf("testdata/go.mod must not require modules, found %q", line)
		default:
			t.Fatalf("unexpected line in testdata/go.mod: %q", line)
		}
	}
	if modulePath != "github.com/jfrog/go-rpm/v2/testdata" {
		t.Fatalf("module path = %q, want github.com/jfrog/go-rpm/v2/testdata", modulePath)
	}
	if goLine != "go 1.15" {
		t.Fatalf("go line = %q, want %q", goLine, "go 1.15")
	}
}
