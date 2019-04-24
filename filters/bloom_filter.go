package filters

import (
	"errors"
	"hash"
	"math"

	"github.com/cespare/xxhash"
	"github.com/twmb/murmur3"
)

// Bloom Error Codes
var (
	ErrInvalidSize          = errors.New("invalid max size, max size should be greater than 0")
	ErrInvalidFalsePositive = errors.New("invalid false positive probability, it should be positive value 0")
)

// BloomFilter definition
type BloomFilter struct {
	bitArray      []uint64      // BloomFilter bit array
	bitArraySize  uint64        // Number of bits to use to persist elements - m
	maxSize       uint64        // Maximum Size of the bloom filter - n
	hashFunctions []hash.Hash64 // List of hash functions
}

// Reference from https://stackoverflow.com/questions/658439/how-many-hash-functions-does-my-bloom-filter-need
func optimalBitArraySize(maxSize uint64, falsePostive float64) uint64 {
	res := math.Ceil(-float64(maxSize) * math.Log(falsePostive) / (math.Ln2 * math.Ln2))
	return uint64(res)
}

// Reference from https://stackoverflow.com/questions/658439/how-many-hash-functions-does-my-bloom-filter-need
func optimalHashSize(maxSize, bitArraySize uint64) uint64 {
	res := math.Ceil(float64(maxSize) * math.Ln2 / float64(bitArraySize))
	return uint64(res)
}

func prepareHashFunctions(size uint64) []hash.Hash64 {
	hashFunctions := make([]hash.Hash64, size)
	var i uint64
	hashFunctions[0] = murmur3.New64()
	for i = 1; i < size; i++ {
		hashFunctions[i] = xxhash.New()
	}
	return hashFunctions

}

// NewBloomFilter returns new BloomFilter object, need to pass maximum size of the bloomfilter
func NewBloomFilter(maxSize uint64, falsePostive float64) (*BloomFilter, error) {
	if maxSize == 0 {
		return nil, ErrInvalidSize
	}

	if falsePostive <= 0 {
		return nil, ErrInvalidFalsePositive
	}

	// bitArraySize := optimalBitArraySize(maxSize, falsePostive)
	// totalHashFunctions := optimalHashSize(maxSize, bitArraySize)
	// bitArray := make([]uint64, bitArraySize)

	return nil, nil
}
