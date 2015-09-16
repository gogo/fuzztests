// Copyright (c) 2015, Vastech SA (PTY) LTD. All rights reserved.
// http://github.com/gogo/fuzztests
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//     * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package fuzztests

import (
	"bytes"
	"fmt"
	"github.com/gogo/fieldpath"
	gogopop "github.com/gogo/fuzztests/gengogofuzztests/gogopop"
	gofast "github.com/gogo/fuzztests/gofast"
	gogo "github.com/gogo/fuzztests/gogo"
	gogofast "github.com/gogo/fuzztests/gogofast"
	golang "github.com/gogo/fuzztests/golang"
	gogoproto "github.com/gogo/protobuf/proto"
	descriptor "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	goproto "github.com/golang/protobuf/proto"
	"reflect"
	"strings"
)

func debug(s string, i int, data []byte, err error) {
	gostringer := gogopop.NewFuncs[i]()
	if err := gogoproto.Unmarshal(data, gostringer); err == nil {
		s += ":" + fmt.Sprintf("%#v", gostringer)
	}
	err = fmt.Errorf("[%d](%T):%s:%v", i, gostringer, s, err)
	panic(err)
}

//proto3 does not keep the XXX_unrecognized bytes
func hasUnrecognized(msg interface{}) bool {
	v := reflect.TypeOf(msg).Elem()
	_, ok := v.FieldByName("XXX_unrecognized")
	return ok
}

//packed fields are not idempotent
//TODO make this function recursive
func hasPacked(msg interface{}) bool {
	v := reflect.TypeOf(msg).Elem()
	num := v.NumField()
	for i := 0; i < num; i++ {
		f := v.Field(i)
		if strings.Contains(string(f.Tag), "packed") {
			return true
		}
	}
	return false
}

type message interface {
	Reset()
	String() string
	ProtoMessage()
}

type marshalFunc func(message) ([]byte, error)
type unmarshalFunc func([]byte, message) error

func assert(s string, i int, input []byte, msg message, marshal marshalFunc, unmarshal unmarshalFunc) {
	if err := unmarshal(input, msg); err != nil {
		debug(s+" unmarshal1", i, input, err)
	}
	output, err := marshal(msg)
	if err != nil {
		panic(err)
	}
	//check the length whenever it is possible, since I found most of the bugs this way.
	if hasUnrecognized(msg) && !hasPacked(msg) {
		s := reflect.TypeOf(msg).Elem()
		pkgPath := s.PkgPath()
		pkgPaths := strings.Split(pkgPath, "/")
		pkgName := pkgPaths[len(pkgPaths)-1]
		msgName := s.Name()
		descriptorSet := gogopop.NewFuncs[i]().(interface {
			Description() *descriptor.FileDescriptorSet
		}).Description()
		//if a field was merged then the length would have changed.
		if err := fieldpath.NoMerge(input, descriptorSet, pkgName, msgName); err == nil {
			//only length is checked since field orders can change
			if len(input) != len(output) {
				panic(fmt.Sprintf("[%d](%T):%s length has changed input %#v output %#v", i, msg, s, input, output))
			}
		}
	}
	msg.Reset()
	if err := unmarshal(output, msg); err != nil {
		debug(s+" unmarshal2", i, output, err)
	}
	output2, err := marshal(msg)
	if err != nil {
		panic(err)
	}
	if !bytes.Equal(output, output2) {
		panic(fmt.Sprintf("[%d](%T):%s is not idempotent input %#v output %#v", i, msg, s, output, output2))
	}
}

func gogomarshal(msg message) ([]byte, error) {
	return gogoproto.Marshal(msg)
}

func gomarshal(msg message) ([]byte, error) {
	return goproto.Marshal(msg)
}

func gogounmarshal(buf []byte, msg message) error {
	return gogoproto.Unmarshal(buf, msg)
}

func gounmarshal(buf []byte, msg message) error {
	return goproto.Unmarshal(buf, msg)
}

func Fuzz(data []byte) int {
	score := 0
	for i := range golang.NewFuncs {
		testpb := gogofast.NewFuncs[i]()
		if err := gogoproto.Unmarshal(data, testpb); err != nil {
			continue
		}
		score = 1
		assert("gogofast", i, data, gogofast.NewFuncs[i](), gogomarshal, gogounmarshal)
		assert("gofast", i, data, gofast.NewFuncs[i](), gomarshal, gounmarshal)
		assert("golang", i, data, golang.NewFuncs[i](), gomarshal, gounmarshal)
		assert("gogo", i, data, gogo.NewFuncs[i](), gogomarshal, gogounmarshal)
	}
	return score
}
