package nanoid_test

import (
	"strings"
	"testing"

	"github.com/bsm/nanoid"
)

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := nanoid.New()
		if len(s) != 21 {
			b.Fatalf("expected 21 chars, but got %d (%s)", len(s), s)
		}
	}
}

func BenchmarkNewSize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := nanoid.NewSize(16)
		if len(s) != 16 {
			b.Fatalf("expected 16 chars, but got %d (%s)", len(s), s)
		}
	}
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
	t.Log(nanoid.Base32.MustGenerate(16))

	s := nanoid.Must(nanoid.Base32.FromReader(strings.NewReader("mOckrand0m"), 10))
	if x := nanoid.ID("NPDLSBOEQN"); x != s {
		t.Fatalf("expected %q but got %q", x, s)
	}

	s = nanoid.Must(nanoid.Base58.FromReader(strings.NewReader("mOckrand0m"), 10))
	if x := nanoid.ID("tNirygujqt"); x != s {
		t.Fatalf("expected %q but got %q", x, s)
	}

	s = nanoid.Must(nanoid.Base64.FromReader(strings.NewReader("mOckrand0m"), 10))
	if x := nanoid.ID("Sm2UN4R1PS"); x != s {
		t.Fatalf("expected %q but got %q", x, s)
	}
}
