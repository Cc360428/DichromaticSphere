/**
 * @Author: Cc
 * @Description: 工具方法
 * @File: utils
 * @Version: 1.0.0
 * @Date: 2023/3/1 10:51
 * @Software : GoLand
 */

package utils

import (
	"fmt"
	"math/rand"
	"sort"
)

func RandInt64s(min, max int64) int64 {
	if min > max {
		max = min
	}
	if max-min == 0 {
		return max
	}
	randNum := rand.Int63n(max - min)
	randNum = randNum + min
	return randNum
}

type ByValue []KeyValue

type KeyValue struct {
	Key   string
	Value int64
}

func (s ByValue) Len() int {
	return len(s)
}

func (s ByValue) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByValue) Less(i, j int) bool {
	return s[i].Value > s[j].Value
}

func SortMapByValue(m map[string]int64) []string {
	sorted := make([]KeyValue, 0, len(m))
	for k, v := range m {
		sorted = append(sorted, KeyValue{k, v})
	}

	sort.Sort(ByValue(sorted))

	sorted = sorted[0:5]
	keys := make([]string, len(sorted))
	for i, kv := range sorted {
		keys[i] = kv.Key
		fmt.Println("number:", kv.Value, " value:", kv.Key)
	}
	return keys
}
