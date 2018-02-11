/**
 * @author liangbo
 * @email  liangbo@codoon.com
 * @date   2018/2/9 下午4:53
 */
package main

//go:generate msgp -o msgp_gen.go -io=false -tests=false
//go:generate protoc --go_out=. protobuf.proto
//go:generate  protoc --gogofaster_out=.  -I. -I$GOPATH/src  mygogo.proto
//go:generate flatc -g -o .. flatbuffers.fbs
//go:generate thrift -r -out ./.. --gen go colorgroup.thrift
//go:generate gencode go -schema=gencode.schema -package gosercomp
//go:generate codecgen -o data_codec.go data.go
type ColorGroup struct {
	Id     int      `json:"id" xml:"id,attr" msg:"id"`
	Name   string   `json:"name" xml:"name" msg:"name"`
	Colors []string `json:"colors" xml:"colors" msg:"colors"`
}
