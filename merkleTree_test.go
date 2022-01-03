package MMR

import (
	"crypto/rand"
	"crypto/sha256"
	"testing"
	"time"
)

func TestNewMerkleTree(t *testing.T) {
	rng := rand.Reader
	num := int64(3000)
	elements := make([][]byte, num)
	for i := 0; i < len(elements); i++ {
		elements[i] = make([]byte, 32)
		for {
			_, err := rng.Read(elements[i])
			if err != nil {
				continue
			} else {
				break
			}
		}
	}
	totalBuildingTime := int64(0)
	totalVerfyTime := int64(0)
	for i := 0; i < 100; i++ {
		startTime := time.Now()
		tree, proof := NewMerkleTree(sha256.New, elements, []int{0})
		endTime := time.Now()
		totalBuildingTime += endTime.UnixNano() - startTime.UnixNano()

		startTime = time.Now()
		if !Verify(elements[0], proof[0], tree.Root) {
			t.Errorf("err")
		}
		endTime = time.Now()
		totalVerfyTime += endTime.UnixNano() - startTime.UnixNano()
	}
	t.Logf("Element Number is %d, Constrcut a merkle tree used %dns, and verify an element with %dns\n", num, totalBuildingTime/100, totalVerfyTime/100)
	//f, err := os.OpenFile("/home/cs331/go/src/github.com/depressi0n/MMR/merkle_test.log", os.O_CREATE|os.O_RDWR, 0644)
	//if err != nil {
	//	return
	//}
	//defer f.Close()
	//for length := 0; length < 100000; length += 100 {
	//	total := int64(0)
	//	verify := int64(0)
	//	repeat := int64(100)
	//	for cnt := int64(0); cnt < repeat; cnt++ {
	//		startTime := time.Now()
	//		//tree, proofs := NewMerkleTree(sha256.New, elements, []int{0, 1, 2, 3, 4, 5})
	//		tree, proof := NewMerkleTree(sha256.New, elements[:length], []int{0})
	//		endTime := time.Now()
	//		total += endTime.UnixNano() - startTime.UnixNano()
	//
	//		if length == 0 {
	//			continue
	//		}
	//
	//		startTime = time.Now()
	//		Verify(elements[0], proof[0], tree.Root)
	//		endTime = time.Now()
	//		verify += endTime.UnixNano() - startTime.UnixNano()
	//		//for i, witness := range proofs {
	//		//	if !Verify(elements[i], witness, tree.Root) {
	//		//		t.Errorf("Not match when  %d append with index %d", elements[i], i)
	//		//	}
	//		//}
	//	}
	//	_, err := fmt.Fprintf(f, "Element Number is %d, Constrcut a merkle tree used %dns, and verify an element with %dns\n", length, total/repeat, verify/repeat)
	//	if err != nil {
	//		return
	//	}
	//	//t.Logf("Element Number is %d, Constrcut a merkle tree used %dns, and verify an element with %dns\n", length, total/repeat, verify/repeat)
	//
	//}
}
