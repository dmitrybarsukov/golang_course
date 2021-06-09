package map_bench_test

import (
    "testing"
)

func BenchmarkMapIntToIntInsertion(b *testing.B) {
    keys := make([]int, b.N, b.N)
    for i := 0; i < b.N; i++ {
        keys[i] = i
    }
    someMap := make(map[int]int)

    b.ResetTimer()

    for _, key := range keys {
        someMap[key] = 0
    }
}

func BenchmarkMapIntToIntSelection(b *testing.B) {
    someMap := make(map[int]int)
    keys := make([]int, b.N, b.N)
    for i := 0; i < b.N; i++ {
        key := i
        keys[i] = key
        someMap[key] = 0
    }

    b.ResetTimer()
    for _, key := range keys {
        val := someMap[key]
        _ = val
    }
}

func BenchmarkMapIntToStructInsertion(b *testing.B) {
    keys := make([]int, b.N, b.N)
    for i := 0; i < b.N; i++ {
        keys[i] = i
    }
    someMap := make(map[int]struct{})

    b.ResetTimer()

    for _, key := range keys {
        someMap[key] = struct{}{}
    }
}

func BenchmarkMapIntToStructSelection(b *testing.B) {
    someMap := make(map[int]struct{})
    keys := make([]int, b.N, b.N)
    for i := 0; i < b.N; i++ {
        key := i
        keys[i] = key
        someMap[key] = struct{}{}
    }

    b.ResetTimer()
    for _, key := range keys {
        val := someMap[key]
        _ = val
    }
}