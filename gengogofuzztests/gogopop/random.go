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
