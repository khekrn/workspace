package filters

import "testing"

func TestAdd(t *testing.T) {
	bloom, _ := NewBloomFilter(1000, 0.1)
	
	v1 := []byte{1, 2, 3, 4, 5}
	bloom.Add((v1))

	if !bloom.Contains(v1){
		t.Error("expected item is true but actual value is false")
	}

	if bloom.Contains([]byte{10}){
		t.Error("expected item is false but actual value is true")
	}
}

func TestContains(b *testing.T) {
	
}