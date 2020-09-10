package nanoid_test

import (
	"io"
	"strings"
	"testing"

	"github.com/bsm/nanoid"
)

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if s := nanoid.New(); len(s) != 21 {
			b.Fatalf("expected 21 chars, but got %d (%s)", len(s), s)
		}
	}
}

func BenchmarkNewSize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if s := nanoid.NewSize(16); len(s) != 16 {
			b.Fatalf("expected 16 chars, but got %d (%s)", len(s), s)
		}
	}
}

func BenchmarkNew_parallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if s := nanoid.New(); len(s) != 21 {
				b.Fatalf("expected 21 chars, but got %d (%s)", len(s), s)
			}
		}
	})
}

func TestNew(t *testing.T) {
	s := nanoid.New()
	if len(s) != 21 {
		t.Fatalf("expected 21 chars, but got %d (%s)", len(s), s)
	}
	if s2 := nanoid.New(); s2 == s {
		t.Fatal("received duplicates")
	}
}

func TestAlphabet_FromReader(t *testing.T) {
	s := nanoid.Must(nanoid.Base32.FromReader(mockEntropy(), 10))
	if x := "npdlsboeqn"; x != s {
		t.Fatalf("expected %q but got %q", x, s)
	}

	s = nanoid.Must(nanoid.Base58.FromReader(mockEntropy(), 10))
	if x := "tNirygujqt"; x != s {
		t.Fatalf("expected %q but got %q", x, s)
	}

	s = nanoid.Must(nanoid.Base64.FromReader(mockEntropy(), 10))
	if x := "Sm2UN4R1PS"; x != s {
		t.Fatalf("expected %q but got %q", x, s)
	}
}

func mockEntropy() io.Reader {
	return strings.NewReader("mOckrand0m")
}
