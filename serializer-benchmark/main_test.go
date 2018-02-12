/**
 * @author liangbo
 * @email  liangbo@codoon.com
 * @date   2018/2/9 下午4:45
 */
package main

import (
	goproto "github.com/gogo/protobuf/proto"
	"github.com/golang/protobuf/proto"

	"encoding/json"
	"encoding/xml"
	msgpackv2 "github.com/vmihailenco/msgpack"
	"log"
	"testing"
)

var group = ColorGroup{
	Id:     1,
	Name:   "Reds",
	Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
}

var protobufGroup = ProtoColorGroup{
	Id:     proto.Int32(1),
	Name:   proto.String("Reds"),
	Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
}

var gogoProtobufGroup = GogoProtoColorGroup{
	Id:     1,
	Name:   "Reds",
	Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
}

func TestMarshaledDataLen(t *testing.T) {
	log.SetFlags(log.LstdFlags)

	buf, _ := json.Marshal(group)
	t.Logf("json:\t\t\t\t %d bytes", len(buf))

	buf, _ = xml.Marshal(group)
	t.Logf("xml:\t\t\t\t %d bytes", len(buf))

	buf, _ = proto.Marshal(&protobufGroup)
	t.Logf("protobuf:\t\t\t\t %d bytes", len(buf))

	buf, _ = goproto.Marshal(&gogoProtobufGroup)
	t.Logf("gogoprotobuf:\t\t\t %d bytes", len(buf))

	buf, _ = group.MarshalMsg(nil)
	t.Logf("msgp:\t\t\t\t %d bytes", len(buf))

	buf, _ = msgpackv2.Marshal(&group)
	t.Logf("msgpack:\t\t\t %d bytes", len(buf))

}

func BenchmarkMarshalByJson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		json.Marshal(group)
	}
}
func BenchmarkUnmarshalByJson(b *testing.B) {
	bytes, _ := json.Marshal(group)
	result := ColorGroup{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		json.Unmarshal(bytes, &result)
	}
}

func BenchmarkMarshalByXml(b *testing.B) {
	for i := 0; i < b.N; i++ {
		xml.Marshal(group)
	}
}
func BenchmarkUnmarshalByXml(b *testing.B) {
	bytes, _ := xml.Marshal(group)
	result := ColorGroup{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		xml.Unmarshal(bytes, &result)
	}
}

func BenchmarkMarshalByProtoBuf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		proto.Marshal(&protobufGroup)
	}
}
func BenchmarkUnmarshalByProtoBuf(b *testing.B) {
	bytes, _ := proto.Marshal(&protobufGroup)
	result := ProtoColorGroup{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		proto.Unmarshal(bytes, &result)
	}
}

func BenchmarkMarshalByGogoProtoBuf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goproto.Marshal(&gogoProtobufGroup)
	}
}
func BenchmarkUnmarshalByGogoProtoBuf(b *testing.B) {
	bytes, _ := proto.Marshal(&gogoProtobufGroup)
	result := GogoProtoColorGroup{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		goproto.Unmarshal(bytes, &result)
	}
}

func BenchmarkMarshalByMsgp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		group.MarshalMsg(nil)
	}
}
func BenchmarkUnmarshalByMsgp(b *testing.B) {
	bytes, _ := group.MarshalMsg(nil)
	result := ColorGroup{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result.UnmarshalMsg(bytes)
	}
}

func BenchmarkMarshalByMsgpackV2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		msgpackv2.Marshal(&group)
	}
}
func BenchmarkUnmarshalByMsgpackv2(b *testing.B) {
	bytes, _ := msgpackv2.Marshal(&group)
	v := &ColorGroup{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		msgpackv2.Unmarshal(bytes, v)
	}
}
