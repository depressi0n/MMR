package MMR

import "hash"

// Using Hash as an updatable module

type HashFunc func() hash.Hash
