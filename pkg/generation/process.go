package generation

import (
	"math/rand"
	"time"
)

//Code генерация защитного кода... указывая длину кода
func CodeChan(ln int, ch chan string) {
	defer close(ch)

	str := "0123456789" + "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ" + "!@#&?%"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < ln; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}

	ch <- string(result)
}

//DateTimeAdd генерация времяни (Установка интервала)
func DateTimeAddChan(addMinute int, ch chan string) {
	defer close(ch)

	ch <- time.Now().Add(time.Minute * time.Duration(addMinute)).Format("2006-01-02 15:04:05")
}
