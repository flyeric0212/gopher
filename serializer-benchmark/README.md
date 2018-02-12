性能对比：go test -bench=. -benchmem

BenchmarkMarshalByJson-4                 1000000              1229 ns/op             376 B/op          4 allocs/op
BenchmarkUnmarshalByJson-4                500000              3065 ns/op             344 B/op          9 allocs/op

BenchmarkMarshalByXml-4                   300000              4698 ns/op            4801 B/op         12 allocs/op
BenchmarkUnmarshalByXml-4                 100000             24595 ns/op            3139 B/op         75 allocs/op

BenchmarkMarshalByProtoBuf-4             2000000               634 ns/op             328 B/op          5 allocs/op
BenchmarkUnmarshalByProtoBuf-4           1000000              1009 ns/op             400 B/op         11 allocs/op

BenchmarkMarshalByGogoProtoBuf-4        10000000               129 ns/op              48 B/op          1 allocs/op
BenchmarkUnmarshalByGogoProtoBuf-4       3000000               507 ns/op             144 B/op          8 allocs/op

BenchmarkMarshalByMsgp-4                10000000               129 ns/op              80 B/op          1 allocs/op
BenchmarkUnmarshalByMsgp-4               5000000               323 ns/op              32 B/op          5 allocs/op

BenchmarkMarshalByMsgpackV2-4            1000000              1907 ns/op             192 B/op          4 allocs/op
BenchmarkUnmarshalByMsgpackv2-4          1000000              1668 ns/op             264 B/op         11 allocs/op


数据大小对比：
json:                                   65 bytes
xml:                                    137 bytes
protobuf:                               36 bytes
gogoprotobuf:                           36 bytes
msgp:                                   47 bytes
msgpack:                                47 bytes
