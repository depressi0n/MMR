package MMR

import (
	"crypto/rand"
	"crypto/sha256"
	"testing"
)

func TestNewMerkleTree(t *testing.T) {
	length := 6
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
	tree, proofs := NewMerkleTree(sha256.New, elements, []int{0, 1, 2, 3, 4, 5})
	for i, witness := range proofs {
		if !Verify(elements[i], witness, tree.Root) {
			t.Errorf("Not match when  %d append with index %d", elements[i], i)
		}
	}
}
