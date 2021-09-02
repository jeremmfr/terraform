package plugins

import (
	"github.com/hashicorp/terraform/internal/addrs"
	"github.com/hashicorp/terraform/internal/providers"
	"github.com/hashicorp/terraform/internal/provisioners"
)

// FinderTestingOverrides is for use with NewFinderForTests to
// bypass all of the usual discovery behaviors inside unit and integration
// tests.
type FinderTestingOverrides struct {
	Providers    map[addrs.Provider]providers.Factory
	Provisioners map[string]provisioners.Factory
}

// NewFinderForTests creates and returns a provider that skips all of the
// usual discovery steps and instead just "finds" exactly the mock plugin
// components given in the argument.
//
// A test-oriented finder can still handle all of the other "With..." methods,
// but they will have no observable effect. The goal is to allow the usual
// codepaths to try to customize the Finder in the ways they usually would,
// but for those calls to be effectively no-ops so that the testing overrides
// can still shine through regardless.
//
// After calling NewFinderForTests, all objects reachable through the given
// overrides object belong to the finder and must not be read or written by
// the caller.
func NewFinderForTests(overrides FinderTestingOverrides) Finder {
	return Finder{
		testingOverrides: &overrides,
	}
}
