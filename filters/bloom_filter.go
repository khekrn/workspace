package filters

import (
	"bytes"
	"encoding/binary"
	"errors"
	"hash"
	"math"
	"math/rand"

	"github.com/OneOfOne/xxhash"
)

// Bloom Error Codes
var (
	ErrInvalidSize          = errors.New("invalid max size, max size should be greater than 0")
	ErrInvalidFalsePositive = errors.New("invalid false positive probability, it should be positive value 0")
)

// BloomFilter struct definition
type BloomFilter struct {
	bitArray      []bool        // BloomFilter bit array
	bitArraySize  int           // Number of bits to use to persist elements - m
	maxSize       int           // Maximum Size of the bloom filter - n
	hashFunctions []hash.Hash64 // List of hash functions
}

func (bl *BloomFilter) fetchIndices(item []byte) []int {
	indexes := make([]int, len(bl.hashFunctions))

	for i, hash := range bl.hashFunctions {
		var hashValue uint64
		hashBinary := hash.Sum(item)
		buf := bytes.NewBuffer(hashBinary)
		binary.Read(buf, binary.LittleEndian, &hashValue)

		bitToBeSet := hashValue % uint64(bl.bitArraySize)
		indexes[i] = int(bitToBeSet)
	}

	return indexes
}

// Add adding an elements to the filter
func (bl *BloomFilter) Add(item []byte) {
	indices := bl.fetchIndices(item)
	for _, index := range indices {
		bl.bitArray[index] = true
	}
}

// Contains checking whether the element exist in the set or not
func (bl *BloomFilter) Contains(item []byte) bool {
	indices := bl.fetchIndices(item)
	for _, index := range indices {
		bitValue := bl.bitArray[index]
		if !bitValue {
			return false
		}
	}
	return true
}

// Reference from https://stackoverflow.com/questions/658439/how-many-hash-functions-does-my-bloom-filter-need
func optimalBitArraySize(maxSize int, falsePostive float64) int {
	res := math.Ceil(-float64(maxSize) * math.Log(falsePostive) / (math.Ln2 * math.Ln2))
	return int(res)
}

// Reference from https://stackoverflow.com/questions/658439/how-many-hash-functions-does-my-bloom-filter-need
func optimalHashSize(maxSize, bitArraySize int) int {
	res := math.Ceil(float64(maxSize) * math.Ln2 / float64(bitArraySize))
	return int(res)
}

func prepareHashFunctions(size int) []hash.Hash64 {
	hashFunctions := make([]hash.Hash64, size)
	for i := 0; i < size; i++ {
		hashFunctions[i] = xxhash.NewS64(rand.Uint64())
	}
	return hashFunctions

}

// NewBloomFilter returns new BloomFilter object, need to pass maximum size of the bloomfilter
func NewBloomFilter(maxSize int, falsePostive float64) (*BloomFilter, error) {
	if maxSize == 0 {
		return nil, ErrInvalidSize
	}

	if falsePostive <= 0 {
		return nil, ErrInvalidFalsePositive
	}

	bitArraySize := optimalBitArraySize(maxSize, falsePostive)
	totalHashFunctions := optimalHashSize(maxSize, bitArraySize)
	bitArray := make([]bool, bitArraySize)
	bloomFilter := BloomFilter{bitArray, bitArraySize, maxSize, prepareHashFunctions(totalHashFunctions)}
	return &bloomFilter, nil
}
