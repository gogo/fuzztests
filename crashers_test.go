package fuzztests

import (
	"testing"
)

func catchPanic(t *testing.T) {
	if r := recover(); r != nil {
		t.Fatalf("%v", r)
	}
}

func Test_WrongWireType_4d5745313695bb32a7be264930411f75ce40ef69(t *testing.T) {
	defer catchPanic(t)
	Fuzz([]byte("j\x00"))
}

func Test_WrongWireType_698bbc06c245ca5787c9628b4ac15909ce3d08cd(t *testing.T) {
	defer catchPanic(t)
	Fuzz([]byte("\xed\x000000\x15\xe7000"))
}

func Test_UnexpectedEOF_a1384fff023991f5e9f0184ba9697cc0ffdf26cb(t *testing.T) {
	defer catchPanic(t)
	Fuzz([]byte("\xed\x0000ܿ\x15\"000"))
}

func Test_UnexpectedEOF_da6c94d3b1f2b0035e00fff878c1b4227fc6753a(t *testing.T) {
	defer catchPanic(t)
	Fuzz([]byte("\xed\x000000"))
}

func Test_WrongWireType_00b4f30b0c1127a29525786214c45b99cde2bd73(t *testing.T) {
	defer catchPanic(t)
	Fuzz([]byte("k0b^0000000000000000" +
		"00000000000000000000" +
		"00000000000000000000" +
		"00000000000000000000" +
		"000000000000000000"))
}

func Test_WrongWireType_21e44783af4a173d22718037577cde36ff826934(t *testing.T) {
	defer catchPanic(t)
	Fuzz([]byte("\xed\x000\b00"))
}
