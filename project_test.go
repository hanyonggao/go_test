package stringinverted

import "testing"

func TestInvert(t *testing.T)  {
	array := "Hello!"
	tmp := "!olleH"
	r := invert(array)
	if tmp != r {
		t.Error("Error!")
		return
	}
}

func TestInvert_withZ(t *testing.T)  {
	array := "ZHello!"
	tmp := "Z!olleH"
	r := invert(array)
	if tmp != r {
		t.Error("Error!")
		return
	}
}

func TestInvert_withextend(t *testing.T)  {
    array := "Hello world, hello zoo!"
    tmp := "!oo olleh ,dlorw olzleH"
    r := invert(array)
    if tmp != r {
        t.Error("Error!")
        return
    }
}