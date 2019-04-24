package hashmap

import (
	"testing"
)

func generateHashMap(size int) *HashMap {
	hashMap := New(size)
	for i := 0; i < size; i++ {
		hashMap.Put(string(i), i)
	}
	return hashMap
}

func generateMap(size int) map[string]interface{} {
	dict := make(map[string]interface{}, size)
	for i := 0; i < size; i++ {
		dict[string(i)] = i
	}
	return dict
}

func TestHashFunction(t *testing.T) {
	res := hashFunction("Burno", 10)
	if res < 0 || res >= 10 {
		t.Error("expected value is between 0 and 9 - ", res)
	}
}

func TestLength(t *testing.T) {
	hashMap := generateHashMap(10)
	if hashMap.Length() != 10 {
		t.Error("expected length is 10 - ", hashMap.Length())
	}

	hashMap.Delete(string(9))
	hashMap.Delete(string(5))

	if hashMap.Length() != 8 {
		t.Error("expected length is 8 - ", hashMap.Length())
	}

	hashMap = generateHashMap(0)
	if hashMap.Length() != 0 {
		t.Error("expected size is 0 for empty map - ", hashMap.Length())
	}
}

func TestPut(t *testing.T) {
	hashMap := generateHashMap(11)
	if hashMap.Length() != 11 {
		t.Error("expected length is 11 - ", hashMap.Length())
	}

	item, _ := hashMap.Get(string(8))
	if item != 8 {
		t.Error("expected 8 found - ", item)
	}
}

func TestGet(t *testing.T) {
	hashMap := generateHashMap(5)
	item, _ := hashMap.Get(string(1))
	if item != 1 {
		t.Error("expected 1 found - ", item)
	}
	item, _ = hashMap.Get(string(100))
	if item != nil {
		t.Error("expected nil found - ", item)
	}
}

func DeleteTest(t *testing.T) {
	hashMap := generateHashMap(5)
	success, _ := hashMap.Delete(string(100))
	if success == true {
		t.Error("expected false found true")
	}

	success, _ = hashMap.Delete(string(5))
	if success == false {
		t.Error("expected true found false")
	}
}

func BenchmarkHashMapPut(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generateHashMap(i)
	}
}

func BenchmarkMapPut(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generateMap(i)
	}
}

func BenchmarkHashMapGet(b *testing.B) {
	small := generateHashMap(100)
	medium := generateHashMap(1000)
	large := generateHashMap(10000)

	b.ResetTimer()

	for i := 0; i < 100; i++ {
		small.Get(string(i))
	}

	for i := 0; i < 1000; i++ {
		medium.Get(string(i))
	}

	for i := 0; i < 10000; i++ {
		large.Get(string(i))
	}
}

func BenchmarkMapGet(b *testing.B) {
	small := generateMap(100)
	medium := generateMap(1000)
	large := generateMap(10000)

	b.ResetTimer()

	for i := 0; i < 100; i++ {
		_ = small[string(i)]
	}

	for i := 0; i < 1000; i++ {
		_ = medium[string(i)]
	}

	for i := 0; i < 10000; i++ {
		_ = large[string(i)]
	}

}

func BenchmarkHashMapDelete(b *testing.B) {
	small := generateHashMap(100)
	medium := generateHashMap(1000)
	large := generateHashMap(10000)

	b.ResetTimer()

	for i := 0; i < 100; i++ {
		small.Delete(string(i))
	}

	for i := 0; i < 1000; i++ {
		medium.Delete(string(i))
	}

	for i := 0; i < 10000; i++ {
		large.Delete(string(i))
	}
}

func BenchmarkMapDelete(b *testing.B) {
	small := generateMap(100)
	medium := generateMap(1000)
	large := generateMap(10000)

	b.ResetTimer()

	for i := 0; i < 100; i++ {
		delete(small, string(i))
	}

	for i := 0; i < 1000; i++ {
		delete(medium, string(i))
	}

	for i := 0; i < 10000; i++ {
		delete(large, string(i))
	}

}
