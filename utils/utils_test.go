/**
 * @Author: Cc
 * @Description: 描述
 * @File: utils_test.go
 * @Version: 1.0.0
 * @Date: 2023/3/1 14:35
 * @Software : GoLand
 */

package utils

import (
	"fmt"
	"testing"
)

func TestSortMapByValue(t *testing.T) {
	ss := SortMapByValue(map[string]int64{
		"Cc": 1,
		"1":  111,
		"2":  222,
		"3":  333,
		"4":  444,
		"5":  555,
	})
	fmt.Println(ss)
}
