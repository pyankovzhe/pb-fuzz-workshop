package kv

import (
	"testing"

	"github.com/stretchr/testify/require"
	"pgregory.net/rapid"
)

func dbCheck(tt *testing.T, fn func(KV, *rapid.T)) {
	rapid.Check(tt, func(t *rapid.T) {
		kv, err := NewInMemoryBadgerKV()
		require.NoError(t, err, "Unable to open a new instance of Badger DB")
		defer func() {
			require.NoError(t, kv.Close())
		}()

		fn(kv, t)
	})
}
