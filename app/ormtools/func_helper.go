/* ====================================================================
 * Author           : tianyh(mknight)
 * Email            : 824338670@qq.com
 * Last modified    : 2022-03-30 11:33
 * Filename         : func_helper.go
 * Description      :
 * ====================================================================*/
package ormtools

import (
	"bufio"
	"io"
	"os"
	"strings"
)

// string -> AaBb,默认下横线分割

func isLowwerChar(c byte) bool {
	return c >= 'a' && c <= 'z'
}

func isUpperChar(c byte) bool {
	return c >= 'A' && c <= 'Z'
}

func charToUpper(c byte) byte {
	if isLowwerChar(c) {
		return c - 32
	}
	return c
}

func charToLower(c byte) byte {
	if isUpperChar(c) {
		return c + 32
	}
	return c
}

func StringToAaBb(src string) string {
	if len(src) < 1 {
		panic("param error")
	}
	r := ""
	words := strings.Split(src, "_")
	for _, w := range words {
		if w == "id" {
			r = r + "ID"
		} else {
			s := []byte(w)
			if isLowwerChar(s[0]) {
				s[0] = (charToUpper(w[0]))
				r = r + string(s)
			}
		}
	}
	return r
}

func StringToaaBb(src string) string {
	if len(src) < 1 {
		panic("param error")
	}
	r := ""
	words := strings.Split(src, "_")
	for i, w := range words {
		if i == 0 {
			if w == "ID" || w == "iD" || w == "Id" {
				r = r + "id"
				continue
			}
			s := []byte(w)
			if isUpperChar(s[0]) {
				s[0] = (charToLower(w[0]))
				r = r + string(s)
			} else {
				r = r + w
			}
		} else {
			r = r + StringToAaBb(w)
		}
	}
	return r
}

// 读取文件内容为行列表
func readFileForLine(fileName string) [][]byte {
	f, err := os.Open(fileName)
	if err != nil {
		return nil
	}
	lineList := make([][]byte, 0)
	fread := bufio.NewReader(f)
	for {
		line, _, err := fread.ReadLine()
		if err == io.EOF {
			break
		}
		lineList = append(lineList, line)
	}
	return lineList
}
