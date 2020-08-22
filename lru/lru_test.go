package lru

import (
	"testing"
)

func TestLRUCache(t *testing.T) {

	cache := Constructor(2)

	cache.Put(1, 1)
	cache.Put(2, 2)

	val := cache.Get(1)

	if val != 1 {
		t.Errorf("expected %d got %d", 1, val)
	}
	cache.Put(3, 3)
	val = cache.Get(2)

	if val != -1 {
		t.Errorf("expected %d got %d", -1, val)
	}

}

func BenchmarkCache_Set1(b *testing.B) {benchmarkSet(1,b)}
func BenchmarkCache_Set10(b *testing.B) {benchmarkSet(10,b)}
func BenchmarkCache_Set100(b *testing.B) {benchmarkSet(100,b)}
func BenchmarkCache_Set1000(b *testing.B) {benchmarkSet(1000,b)}
func BenchmarkCache_Set10000(b *testing.B) {benchmarkSet(10000,b)}


func BenchmarkCache_Get1(b *testing.B) {benchmarkGet(1,b)}
func BenchmarkCache_Get10(b *testing.B) {benchmarkGet(10,b)}
func BenchmarkCache_Get100(b *testing.B) {benchmarkGet(100,b)}
func BenchmarkCache_Get1000(b *testing.B) {benchmarkGet(1000,b)}
func BenchmarkCache_Get10000(b *testing.B) {benchmarkGet(10000,b)}



func benchmarkSet(count int, b *testing.B) {
	c := Constructor(10)


	b.ResetTimer()
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		for i := 0; i < count; i++ {
			c.Put(1,2)
		}
	}
}
func benchmarkGet(count int, b *testing.B) {
	c := Constructor(10)
	c.Put(1,2)

	b.ResetTimer()
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		for i := 0; i < count; i++ {
			c.Get(1)
		}
	}
}
