package MMR

import (
	"bytes"
)

type MerkleTree struct {
	Root []byte // from low to high
	Len  int
	Hash HashFunc
	Zero []byte
}

// NewMerkleTree construct a merkle tree from elements, and generate the merkle proof of indexes in proveIndex
func NewMerkleTree(hash HashFunc, element [][]byte, proveIndex []int) (*MerkleTree, map[int]*Witness) {
	if len(element) == 0 {
		return nil, nil
	}
	// compute the height of merkle tree
	height, count := 0, len(element)
	for count != 0 {
		count >>= 1
		height++
	}
	f := hash()
	tree := make([][][]byte, 1, height+1)
	// leaves

	tree[0] = make([][]byte, len(element))
	cur := 0
	for ; cur < len(element); cur++ {
		tree[0][cur] = make([]byte, f.Size())
		f.Reset()
		f.Write(element[cur])
		copy(tree[0][cur], f.Sum(nil))
	}
	// middle layer
	count = ((len(element) + 1) >> 1) << 1
	for count > 1 {
		cur = 0
		count = (count + 1) >> 1
		tmp := make([][]byte, count)
		for ; cur < count; cur++ {
			tmp[cur] = make([]byte, f.Size())
			f.Reset()
			f.Write(tree[len(tree)-1][cur<<1])
			if cur<<1+1 < len(tree[len(tree)-1]) {
				f.Write(tree[len(tree)-1][cur<<1+1])
			} else {
				f.Write(tree[len(tree)-1][cur<<1])
			}
			copy(tmp[cur], f.Sum(nil))
		}
		tree = append(tree, tmp)
	}
	// proof
	m := make(map[int]*Witness)
	for i := 0; i < len(proveIndex); i++ {
		t := &Witness{
			Hash: hash,
			Dir:  make([]bool, 0, len(tree)-1),
			Wit:  make([][]byte, 0, len(tree)-1),
		}
		index := proveIndex[i]
		if index >= len(element) {
			continue
		}
		//fmt.Printf("%d:", index)
		cur = 0
		for cur < len(tree)-1 {
			if index&1 == 0 {
				if index+1 >= len(tree[cur]) {
					t.Append(true, tree[cur][index])
					//fmt.Printf("%d->", index)
				} else {
					t.Append(true, tree[cur][index+1])
					//fmt.Printf("%d->", index+1)
				}
			} else {
				t.Append(false, tree[cur][index-1])
				//fmt.Printf("%d->", index-1)
			}
			cur += 1
			index >>= 1
		}
		//fmt.Println()
		m[proveIndex[i]] = t
	}

	zero := make([]byte, f.Size())
	return &MerkleTree{
		Root: tree[len(tree)-1][0],
		Len:  len(element),
		Hash: hash,
		Zero: zero,
	}, m
}

func Verify(element []byte, witness *Witness, root []byte) bool {
	path := GetAncestor(element, witness)
	if bytes.Equal(path[len(path)-1], root) {
		return true
	} else {
		return false
	}
}
