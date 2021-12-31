package MMR

import (
	"crypto/rand"
	"crypto/sha256"
	"testing"
)

func TestMMR_Correct(t *testing.T) {
	acc := NewMMR(sha256.New)
	length := 256
	elements := make([][]byte, length)
	rng := rand.Reader
	for i := 0; i < len(elements); i++ {
		elements[i] = make([]byte, length)
		for {
			_, err := rng.Read(elements[i])
			if err != nil {
				continue
			} else {
				break
			}
		}
	}
	first := acc.Add(elements[0])
	if !acc.MemVerify(elements[0], first) {
		t.Errorf("Not match with %d", elements[0])
	}
	for i := 1; i < len(elements); i++ {
		got := acc.Add(elements[i])
		if !acc.MemVerify(elements[i], got) {
			t.Errorf("Not match when  %d append with index %d", elements[i], i)
		}

		first.Update(elements[i], got)
		if !acc.MemVerify(elements[0], first) {
			t.Errorf("Not match when  %d append with index %d", elements[i], i)
		}
	}
}
