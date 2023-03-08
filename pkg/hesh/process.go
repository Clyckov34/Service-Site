package hesh

import (
	"crypto/sha512"
	"fmt"
)

//соль
const (
	salt_left  = "k_LYw73^5Td7eVd=e:P:dB5Gwab^@sjeDy0HRr]h.!kbHsZd3v"
	salt_right = "]:X+T3ubac1.R5QYq4H~hJmf*6KQi}i_Y5u7X0uz+v)fB9oRL-"
)

//CheckHesh проверка хеша
func Check(data, hesh string) bool {
	return NewHesh512(data) == hesh
}

//newHesh256 хеширует
func NewHesh512(data string) string {
	res := salt_left + data + salt_right

	sh := sha512.New()
	sh.Write([]byte(res))
	return fmt.Sprintf("%x", sh.Sum(nil))
}
