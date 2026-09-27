package pokecache

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestAddGet(t *testing.T) {
	interval := 5 * time.Second
	tests := map[string]struct {
		input struct {
			key string
			val []byte
		}
		want []byte
	}{
		"simple": {
			input: struct {
				key string
				val []byte
			}{
				key: "https://example.com",
				val: []byte("testdata"),
			},
			want: []byte("testdata"),
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(tc.input.key, tc.input.val)
			got, ok := cache.Get(tc.input.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := NewCache(baseTime)
	key := "https://example.com"
	cache.Add(key, []byte("testdata"))

	_, ok := cache.Get(key)
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get(key)
	if ok {
		t.Errorf("expected to not find key")
		return
	}

}
