package nanoid

import (
	"crypto/rand"
	"errors"
	"io"
)

var (
	errMinSymbols  = errors.New("alphabet must contain at least 2 unique symbols")
	errOnlyASCII   = errors.New("alphabet must not conain non-ASCII symbols")
	errDupeSymbols = errors.New("alphabet must not contain duplicate symbols")
)

// DefaultSize is the default ID length.
const DefaultSize = 21

// Base64 represents the standard base64 nanoid encoding.
var Base64, _ = NewEncoding("-_zyxwvutsrqponmlkjihgfedcba9876543210ZYXWVUTSRQPONMLKJIHGFEDCBA")

// Base58 is a more human-friendly base58 encoding.
var Base58, _ = NewEncoding("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")

// Base32 uses standard Base32 alphabet.
var Base32, _ = NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZ234567")

// Encoding defines the characters that consitute the output ID.
type Encoding struct{ alphabet string }

// NewEncoding inits a new Encoding.
func NewEncoding(alphabet string) (*Encoding, error) {
	if len(alphabet) < 2 {
		return nil, errMinSymbols
	}

	seen := make(map[byte]struct{})
	for _, r := range alphabet {
		if r > '\u007F' {
			return nil, errOnlyASCII
		}
		c := byte(r)
		if _, ok := seen[c]; ok {
			return nil, errDupeSymbols
		}
		seen[c] = struct{}{}
	}
	return &Encoding{alphabet: alphabet}, nil
}

// MustGenerate creates a new random ID or panics. It is equivalent to
// the expression:
//
//    nanoid.Must(Encoding.Generate(size))
func (e *Encoding) MustGenerate(size int) string {
	return Must(e.Generate(size))
}

// Generate generates a new ID from a cryptographically random source.
func (e *Encoding) Generate(size int) (string, error) {
	return e.FromReader(rand.Reader, size)
}

// FromReader generates a new ID from a reader.
// If size of <1 is passed, DefaultSize will be assumed.
func (e *Encoding) FromReader(r io.Reader, size int) (string, error) {
	if size < 1 {
		size = 21
	}

	bytes := make([]byte, size)
	if _, err := io.ReadFull(r, bytes); err != nil {
		return "", err
	}

	n := byte(len(e.alphabet))
	for i, c := range bytes {
		bytes[i] = e.alphabet[c%n]
	}

	return string(bytes), nil
}

// Must returns uuid if err is nil and panics otherwise.
func Must(id string, err error) string {
	if err != nil {
		panic(err)
	}
	return id
}

// New creates a new random ID or panics. New is a short-cut
// for:
//
//    nanoid.NewSize(DefaultSize)
func New() string {
	return NewSize(DefaultSize)
}

// NewSize creates a new random ID or panics. NewSize is a short-cut
// for:
//
//    nanoid.Base64.NewSize(size)
func NewSize(size int) string {
	return Base64.MustGenerate(size)
}
