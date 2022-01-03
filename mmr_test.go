package MMR

import (
	"crypto/rand"
	"crypto/sha256"
	"testing"
	"time"
)

func TestMMR_Correct(t *testing.T) {
	num := 714100
	elements := make([][]byte, num)
	rng := rand.Reader
	for i := 0; i < num; i++ {
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
	totalUpdateTime := int64(0)
	totalAddTime := int64(0)
	totalVerfyTime := int64(0)
	for i := 0; i < 100; i++ {
		acc := NewMMR(sha256.New)
		first := acc.Add(elements[0])

		for i := 1; i < num; i++ {
			proof := acc.Add(elements[i])
			//endTime = time.Now()
			//t.Logf("%dns for adding a element to the accumulator with %d elements\n", endTime.UnixNano()-startTime.UnixNano(), i+1)

			startTime := time.Now()
			first.Update(elements[i], proof)
			endTime := time.Now()
			totalUpdateTime += endTime.UnixNano() - startTime.UnixNano()
			//fmt.Fprintf(f, "%dns for update a element to the accumulator with %d elements\n", endTime.UnixNano()-startTime.UnixNano(), i+1)
		}

		startTime := time.Now()
		x := acc.Add(elements[0])
		endTime := time.Now()
		totalAddTime += endTime.UnixNano() - startTime.UnixNano()

		startTime = time.Now()
		first.Update(elements[0], x)
		endTime = time.Now()
		totalUpdateTime += endTime.UnixNano() - startTime.UnixNano()

		startTime = time.Now()
		if !acc.MemVerify(elements[0], first) {
			t.Errorf("err")
		}
		endTime = time.Now()
		totalVerfyTime += endTime.UnixNano() - startTime.UnixNano()
	}
	t.Logf("Element Number is %d, add an element to accumulator using %dns, total update time is %dns, and verify it with %dns\n", num, totalAddTime/100, totalUpdateTime/100, totalVerfyTime/100)

	//f, err := os.OpenFile("/home/cs331/go/src/github.com/depressi0n/MMR/mmr_test.log", os.O_CREATE|os.O_RDWR, 0644)
	//if err != nil {
	//	return
	//}
	//defer f.Close()
	////
	//for length := 1; length < num; length += 5000 {
	//	fmt.Fprintf(f, "------START(%d)------\n", length)
	//	acc := NewMMR(sha256.New)
	//	// add the first element
	//	//startTime := time.Now()
	//	first := acc.Add(elements[0])
	//	//endTime := time.Now()
	//	//fmt.Fprintf(f, "%dns for adding a element to the accumulator with %d elements\n", endTime.UnixNano()-startTime.UnixNano(), 1)
	//
	//	//startTime = time.Now()
	//	//acc.MemVerify(elements[0], first)
	//	//endTime = time.Now()
	//	//fmt.Fprintf(f, "%dns for verify a element to the accumulator with %d elements\n", endTime.UnixNano()-startTime.UnixNano(), 1)
	//
	//	totalVerifyTime := int64(0)
	//	for i := 1; i < length-1; i++ {
	//		//startTime = time.Now()
	//		proof := acc.Add(elements[i])
	//		//endTime = time.Now()
	//		//t.Logf("%dns for adding a element to the accumulator with %d elements\n", endTime.UnixNano()-startTime.UnixNano(), i+1)
	//
	//		startTime := time.Now()
	//		first.Update(elements[i], proof)
	//		endTime := time.Now()
	//		totalVerifyTime += endTime.UnixNano() - startTime.UnixNano()
	//		//fmt.Fprintf(f, "%dns for update a element to the accumulator with %d elements\n", endTime.UnixNano()-startTime.UnixNano(), i+1)
	//
	//	}
	//
	//	startTime := time.Now()
	//	proof := acc.Add(elements[length-1])
	//	endTime := time.Now()
	//	fmt.Fprintf(f, "%dns for adding a element to the accumulator with %d elements\n", endTime.UnixNano()-startTime.UnixNano(), length)
	//	startTime = time.Now()
	//	first.Update(elements[length-1], proof)
	//	endTime = time.Now()
	//	totalVerifyTime += endTime.UnixNano() - startTime.UnixNano()
	//	//endTime = time.Now()
	//	//t.Logf("%dns for adding a element to the accumulator with %d elements\n", endTime.UnixNano()-startTime.UnixNano(), i+1)
	//	fmt.Fprintf(f, "%dns for update a element to the accumulator with %d elements\n", totalVerifyTime, length)
	//	startTime = time.Now()
	//	acc.MemVerify(elements[0], first)
	//	endTime = time.Now()
	//	fmt.Fprintf(f, "%dns for verify a element to the accumulator with %d elements\n", endTime.UnixNano()-startTime.UnixNano(), length)
	//	fmt.Fprintf(f, "------END(%d)------\n", length)
	//}

	//repeat := int64(1)
	//for cnt := int64(0); cnt < repeat; cnt++ {
	//
	//}

	//first := acc.Add(elements[0])
	//if !acc.MemVerify(elements[0], first) {
	//	t.Errorf("Not match when  %d append with index %d", elements[0], 0)
	//}
	//for i := 1; i < len(elements); i++ {
	//	startTime := time.Now()
	//	got := acc.Add(elements[i])
	//	endTime := time.Now()
	//	fmt.Printf("%ds for adding a element to the accumulator with %d elements", endTime.Unix()-startTime.Unix(), i)
	//	if !acc.MemVerify(elements[i], got) {
	//		t.Errorf("Not match when  %d append with index %d", elements[i], i)
	//	}
	//
	//	first.Update(elements[i], got)
	//	if !acc.MemVerify(elements[0], first) {
	//		t.Errorf("Not match when  %d append with index %d", elements[i], i)
	//	}
	//}

}
