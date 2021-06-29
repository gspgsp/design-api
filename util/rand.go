package util

import (
	"fmt"
	"bytes"
	"crypto/rand"
	"math/big"
	"log"
)

type randDom interface {
	Generate(len int)
}

type RandStr struct {
}

type RandInt struct {
}

/**
生成随机字符串
 */
func (rs *RandStr) Generate(lens int) (string) {

	var str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	var container string

	var buffer bytes.Buffer
	buffer.WriteString(str)

	//b := bytes.NewBufferString(str)
	//length := b.Len()
	//length := len(b.Bytes())
	//bigInt := big.NewInt(int64(length))

	//length := len(str)

	length := len(buffer.Bytes())
	bigInt := big.NewInt(int64(length))

	for i := 1; i <= lens; i++ {
		random, err := rand.Int(rand.Reader, bigInt)
		if err != nil {
			log.Println("随机字符串生成错误:" + err.Error())
		} else {
			container += string(str[random.Int64()])
		}
	}

	return container
}

/**
生成随机数字
 */
func (ri *RandInt) Generate(len int) (string) {

	var numbers = []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	var container string
	length := bytes.NewReader(numbers).Len()

	for i := 1; i <= len; i++ {
		random, err := rand.Int(rand.Reader, big.NewInt(int64(length)))
		if err != nil {
			log.Println("随机数字生成错误:" + err.Error())
		} else {
			container += fmt.Sprintf("%d", numbers[random.Int64()])
		}
	}

	return container
}
