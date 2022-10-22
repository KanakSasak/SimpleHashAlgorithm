package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	input  = "input plain text to be hash"
	newbit = 125
)

func main() {
	stringbit := binary(input)
	//log.Println(len(stringbit))

	digest := SHF(stringbit)

	fmt.Println("Digest :", digest)
}

func binary(s string) string {
	res := ""
	for _, c := range s {
		res = fmt.Sprintf("%s%.8b", res, c)
	}
	return res
}

func SHF(stringbit string) string {

	arr := [len(input)][newbit]int{}
	digest := [newbit]int{}
	var block = 0
	var lastflag = 0
	for i := 1; i <= len(input); i++ {
		value := newbit + (i-1)*newbit //aritmathic sequence
		//log.Println(value)
		chars := []rune(stringbit)
		//log.Println(lastflag, value)
		idx := 0
		for j := lastflag; j < value; j++ {
			if j >= len(stringbit) {
				arr[block][idx] = 0

				//log.Println(block, idx, 0)
			} else {
				char := string(chars[j])
				intVar, _ := strconv.Atoi(char)
				arr[block][idx] = intVar

				//log.Println(block, idx, char)
			}

			idx++
		}

		lastflag = value

		block++
	}

	//log.Println(len(arr))

	arrt := transpose(arr)

	//fmt.Println(arr)
	//fmt.Println(arrt)

	for x := 0; x < newbit; x++ {
		for i := 0; i < len(input); i++ {
			if len(digest) == 0 {
				digest[x] = arrt[x][i]
			} else {
				digest[x] = digest[x] ^ arrt[x][i]
			}

		}

	}

	//log.Println(digest)

	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(digest)), ""), "[]")
}

func transpose(slice [len(input)][newbit]int) [][]int {
	xl := len(slice[0])
	yl := len(slice)
	result := make([][]int, xl)
	for i := range result {
		result[i] = make([]int, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = slice[j][i]
		}
	}
	return result
}
