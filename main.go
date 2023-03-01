package main

import (
	"github.com/Cc360428/DichromaticSphere/ball"
	"github.com/Cc360428/DichromaticSphere/utils"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // 初始化随机种子
	ballMap := make(map[string]int64)
	for i := 0; i < ball.MaxCount; i++ {
		ballMap[ball.GenerateBall()]++
	}
	utils.SortMapByValue(ballMap)
}

/**
1000000000

number: 104  value:  01  03  06  08  21  32 10
number: 101  value:  03  06  07  08  12  29 04
number: 97   value:  10  15  16  19  24  28 06
number: 97   value:  01  06  10  18  27  33 02
number: 97   value:  18  19  26  29  30  33 03

177210880

number: 36  value:  05  15  18  26  28  30 05
number: 30  value:  01  04  08  26  30  32 03
number: 29  value:  08  09  16  21  22  29 07
number: 29  value:  03  11  12  14  27  33 01
number: 29  value:  07  10  20  25  26  27 04

number: 30  value:  08  11  18  20  22  24 10
number: 30  value:  01  07  08  10  22  33 03
number: 29  value:  04  13  17  20  28  29 16
number: 29  value:  01  02  15  20  28  31 10
number: 29  value:  01  12  13  24  28  31 04

*/
