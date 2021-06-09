## Key: string

### Insertion
* int      -> 442 ns/op              82 B/op          0 allocs/op
* struct{} -> 397 ns/op              56 B/op          0 allocs/op

### Selection
* int      -> 147 ns/op               0 B/op          0 allocs/op
* struct{} -> 146 ns/op               0 B/op          0 allocs/op

## Key: int

### Insertion
* int      -> 269 ns/op              59 B/op          0 allocs/op
* struct{} -> 231 ns/op              34 B/op          0 allocs/op

### Selection
* int      -> 108 ns/op               0 B/op          0 allocs/op
* struct{} -> 116 ns/op               0 B/op          0 allocs/op
