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
	"github.com/gogo/protobuf/proto"
	"math/rand"
	"reflect"
	"time"
)

func Randoms() [][]byte {
	amount := 2 * len(popFuncs)
	randoms := make([][]byte, amount)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	easy := reflect.ValueOf(true)
	rr := reflect.ValueOf(r)
	for i := 0; i < amount; i++ {
		j := i
		if i >= len(popFuncs) {
			j = r.Intn(len(popFuncs))
		}
		f := reflect.ValueOf(popFuncs[j])
		v := f.Call([]reflect.Value{rr, easy})
		data, err := proto.Marshal(v[0].Interface().(proto.Message))
		if err != nil {
			panic(err)
		}
		randoms[i] = data
	}
	return randoms
}
