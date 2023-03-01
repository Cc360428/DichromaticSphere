/**
 * @Author: Cc
 * @Description: 球
 * @File: ball
 * @Version: 1.0.0
 * @Date: 2023/3/1 10:43
 * @Software : GoLand
 */

package ball

import (
	"fmt"
	"github.com/Cc360428/DichromaticSphere/utils"
	"sort"
	"strings"
)

const (
	redCount  = 6
	blueCount = 1
)

var (
	_redBall = []Ball{
		{value: 1, name: " 01 "},
		{value: 2, name: " 02 "},
		{value: 3, name: " 03 "},
		{value: 4, name: " 04 "},
		{value: 5, name: " 05 "},
		{value: 6, name: " 06 "},
		{value: 7, name: " 07 "},
		{value: 8, name: " 08 "},
		{value: 9, name: " 09 "},
		{value: 10, name: " 10 "},
		{value: 11, name: " 11 "},
		{value: 12, name: " 12 "},
		{value: 13, name: " 13 "},
		{value: 14, name: " 14 "},
		{value: 15, name: " 15 "},
		{value: 16, name: " 16 "},
		{value: 17, name: " 17 "},
		{value: 18, name: " 18 "},
		{value: 19, name: " 19 "},
		{value: 20, name: " 20 "},
		{value: 21, name: " 21 "},
		{value: 22, name: " 22 "},
		{value: 23, name: " 23 "},
		{value: 24, name: " 24 "},
		{value: 25, name: " 25 "},
		{value: 26, name: " 26 "},
		{value: 27, name: " 27 "},
		{value: 28, name: " 28 "},
		{value: 29, name: " 29 "},
		{value: 30, name: " 30 "},
		{value: 31, name: " 31 "},
		{value: 32, name: " 32 "},
		{value: 33, name: " 33 "},
	}

	_blueBall = []Ball{
		{value: 1, name: "01"},
		{value: 2, name: "02"},
		{value: 3, name: "03"},
		{value: 4, name: "04"},
		{value: 5, name: "05"},
		{value: 6, name: "06"},
		{value: 7, name: "07"},
		{value: 8, name: "08"},
		{value: 9, name: "09"},
		{value: 10, name: "10"},
		{value: 11, name: "11"},
		{value: 12, name: "12"},
		{value: 13, name: "13"},
		{value: 14, name: "14"},
		{value: 15, name: "15"},
		{value: 16, name: "16"},
	}
)

type Ball struct {
	value int
	name  string
}

type List struct {
	Red  []Ball
	Blue []Ball
}

func NewBallList() *List {

	b := &List{
		Red:  make([]Ball, 0),
		Blue: make([]Ball, 0),
	}

	_red := _redBall
	for i := len(_red) - 1; i >= 0; i-- {
		randCardIndex := utils.RandInt64s(0, int64(i))
		a := _red[i]
		_red[i] = _red[randCardIndex]
		_red[randCardIndex] = a
	}
	b.Red = _red

	_blue := _blueBall
	for i := len(_blue) - 1; i >= 0; i-- {
		randCardIndex := utils.RandInt64s(0, int64(i))
		a := _blue[i]
		_blue[i] = _blue[randCardIndex]
		_blue[randCardIndex] = a
	}
	b.Blue = _blue

	return b
}

const MaxCount = 177210880

/*
number: 4  value:  28  11  02  21  23  29 08
number: 4  value:  23  20  29  19  25  12 10
number: 4  value:  19  03  32  07  29  12 07
number: 3  value:  10  16  32  11  13  02 07
number: 3  value:  19  15  25  21  29  23 08
number: 3  value:  24  31  29  17  19  12 01
number: 3  value:  06  05  29  21  02  26 02
number: 3  value:  05  17  29  07  04  15 07
number: 3  value:  10  01  04  05  32  15 02
number: 3  value:  26  30  25  16  27  23 07
*/

func GenerateBall() string {
	list := NewBallList()
	var redName []string
	red := list.Red[0:redCount]
	for _, ball := range red {
		redName = append(redName, ball.name)
	}
	sort.Strings(redName)
	return fmt.Sprintf("%s%s", strings.Join(redName, ""), list.Blue[0:blueCount][0].name)
}
