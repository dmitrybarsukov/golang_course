package map_bench_test

import (
    "strconv"
    "testing"
)

func BenchmarkMapStringToIntInsertion(b *testing.B) {
    keys := make([]string, b.N, b.N)
    for i := 0; i < b.N; i++ {
        keys[i] = strconv.Itoa(i)
    }
    someMap := make(map[string]int)

    b.ResetTimer()

    for _, key := range keys {
        someMap[key] = 0
    }
}

func BenchmarkMapStringToIntSelection(b *testing.B) {
    someMap := make(map[string]int)
    keys := make([]string, b.N, b.N)
    for i := 0; i < b.N; i++ {
        key := strconv.Itoa(i)
        keys[i] = key
        someMap[key] = 0
    }

    b.ResetTimer()
    for _, key := range keys {
        val := someMap[key]
        _ = val
    }
}

func BenchmarkMapStringToStructInsertion(b *testing.B) {
    keys := make([]string, b.N, b.N)
    for i := 0; i < b.N; i++ {
        keys[i] = strconv.Itoa(i)
    }
    someMap := make(map[string]struct{})

    b.ResetTimer()

    for _, key := range keys {
        someMap[key] = struct{}{}
    }
}

func BenchmarkMapStringToStructSelection(b *testing.B) {
    someMap := make(map[string]struct{})
    keys := make([]string, b.N, b.N)
    for i := 0; i < b.N; i++ {
        key := strconv.Itoa(i)
        keys[i] = key
        someMap[key] = struct{}{}
    }

    b.ResetTimer()
    for _, key := range keys {
        val := someMap[key]
        _ = val
    }
}