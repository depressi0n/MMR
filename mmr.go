package MMR

import (
	"bytes"
)

type MMR struct {
	Peaks [][]byte // from low to high
	Len   uint64
	Hash  HashFunc
	Zero  []byte
}
type Witness struct {
	Hash HashFunc
	Dir  []bool // left -> False, right -> True
	Wit  [][]byte
}

func (w *Witness) Append(dir bool, wit []byte) {
	w.Dir = append(w.Dir, dir)
	w.Wit = append(w.Wit, wit)
	return
}

// Update is MemWitUpdate in accumulator
func (w *Witness) Update(element []byte, witness *Witness) {
	if len(witness.Dir) < len(w.Dir) {
		//fmt.Printf("0")
		return
	}
	yAncestors := GetAncestor(element, witness)
	i := len(w.Dir)
	for ; i < len(yAncestors)-1; i++ {
		w.Dir = append(w.Dir, true)
		w.Wit = append(w.Wit, yAncestors[i])
	}
	//fmt.Printf("1")
	return
}
func NewMMR(hash HashFunc) *MMR {
	f := hash()
	zero := make([]byte, f.Size())
	return &MMR{
		Peaks: [][]byte{},
		Len:   0,
		Hash:  hash,
		Zero:  zero,
	}
}

// Add append an element to the accumulator
func (m *MMR) Add(element []byte) *Witness {
	updatedPeaks := make([][]byte, len(m.Peaks))
	hashFunc := m.Hash()
	for i := 0; i < len(m.Peaks); i++ {
		updatedPeaks[i] = make([]byte, hashFunc.Size())
	}
	res := &Witness{
		Hash: m.Hash,
		Dir:  make([]bool, 0, len(m.Peaks)),
		Wit:  make([][]byte, 0, len(m.Peaks)),
	}
	pos := 0
	z := make([]byte, hashFunc.Size())
	hashFunc.Reset()
	copy(z, hashFunc.Sum(element))
	for pos < len(m.Peaks) && m.Peaks[pos] != nil {
		if bytes.Equal(m.Peaks[pos], m.Zero) {
			break
		}
		hashFunc.Reset()
		hashFunc.Write(m.Peaks[pos])
		hashFunc.Write(z)
		copy(z, hashFunc.Sum(nil))
		updatedPeaks[pos] = nil

		tmp := make([]byte, len(m.Peaks[pos]))
		copy(tmp, m.Peaks[pos])
		res.Append(false, tmp)

		pos++
	}
	// update the peaks
	if pos == len(m.Peaks) {
		updatedPeaks = append(updatedPeaks, z)
	} else {
		updatedPeaks[pos] = z
		pos++
		for pos < len(m.Peaks) {
			copy(updatedPeaks[pos], m.Peaks[pos])
			pos++
		}
	}
	m.Peaks = updatedPeaks
	m.Len++
	return res
}
func (m *MMR) MemVerify(element []byte, witness *Witness) bool {
	ancestors := GetAncestor(element, witness)
	root := ancestors[len(ancestors)-1]
	for i := 0; i < len(m.Peaks); i++ {
		if bytes.Equal(m.Peaks[i], root) {
			return true
		}
	}
	return false
}
func GetAncestor(element []byte, witness *Witness) [][]byte {
	res := make([][]byte, 0, len(witness.Dir)+1)
	f := witness.Hash()
	f.Reset()

	c := make([]byte, f.Size())
	copy(c, f.Sum(element))
	tmp := make([]byte, f.Size())
	copy(tmp, c)
	res = append(res, tmp)
	pos := 0
	for pos < len(witness.Dir) {
		f.Reset()
		if witness.Dir[pos] { //true -> right
			f.Write(c)
			f.Write(witness.Wit[pos])
		} else { //false -> left
			f.Write(witness.Wit[pos])
			f.Write(c)
		}
		copy(c, f.Sum(nil))
		tmp := make([]byte, f.Size())
		copy(tmp, c)
		res = append(res, tmp)
		pos++
	}
	return res
}
