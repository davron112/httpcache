package test_test

import (
	"testing"

	"github.com/davron112/httpcache"
	"github.com/davron112/httpcache/test"
)

func TestMemoryCache(t *testing.T) {
	test.Cache(t, httpcache.NewMemoryCache())
}
