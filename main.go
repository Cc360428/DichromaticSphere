package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 初始化随机数生成器
	rand.Seed(time.Now().UnixNano())

	// 生成6个不重复的红色球
	var redBalls []int
	for len(redBalls) < 6 {
		num := rand.Intn(33) + 1
		if !contains(redBalls, num) {
			redBalls = append(redBalls, num)
		}
	}

	// 生成1个蓝色球
	blueBall := rand.Intn(16) + 1

	// 打印结果
	fmt.Printf("红球：%v，蓝球：%d\n", redBalls, blueBall)

	// 计算概率最大的中奖号码
	maxProbability(redBalls, blueBall)
}

// 判断一个int类型的切片中是否包含某个元素
func contains(slice []int, item int) bool {
	for _, val := range slice {
		if val == item {
			return true
		}
	}
	return false
}

// 计算概率最大的中奖号码
func maxProbability(redBalls []int, blueBall int) {
	var maxRedBalls []int
	var maxBlueBall int
	var maxProbability float64

	// 遍历所有可能的红球和蓝球组合，计算中奖概率
	for i := 1; i <= 33; i++ {
		for j := i + 1; j <= 33; j++ {
			for k := j + 1; k <= 33; k++ {
				for l := k + 1; l <= 33; l++ {
					for m := l + 1; m <= 33; m++ {
						for n := m + 1; n <= 33; n++ {
							for b := 1; b <= 16; b++ {
								r := []int{i, j, k, l, m, n}
								p := probability(r, b, redBalls, blueBall)
								if p > maxProbability {
									maxRedBalls = r
									maxBlueBall = b
									maxProbability = p
								}
							}
						}
					}
				}
			}
		}
	}

	// 打印结果
	fmt.Printf("概率最大的中奖号码：%v，蓝球：%d，概率：%f\n", maxRedBalls, maxBlueBall, maxProbability)
}

// 计算某个红球和蓝球组合的中奖概率
func probability(redBalls []int, blueBall int, winRedBalls []int, winBlueBall int) float64 {
	redMatches := 0
	for _, val := range redBalls {
		if contains(winRedBalls, val) {
			redMatches++
		}
	}
	blueMatch := blueBall == winBlueBall
	if redMatches == 0 && blueMatch {
		return 1.0 / 17721088.0
	} else if redMatches == 6 && !blueMatch {
		return 1.0 / 296017.0
	} else if redMatches == 6 && blueMatch {
		return 1.0 / 177210880.0
	} else if redMatches == 5 && blueMatch {
		return 1.0 / 35508216.0
	} else if redMatches == 5 && !blueMatch {
		return 1.0 / 591803.6
	} else if redMatches == 4 && blueMatch {
		return 1.0 / 36252.0
	} else if redMatches == 4 && !blueMatch {
		return 1.0 / 723.0
	} else if redMatches == 3 && blueMatch {
		return 1.0 / 1726.0
	} else if redMatches == 2 && blueMatch {
		return 1.0 / 96.0
	} else if redMatches == 1 && blueMatch {
		return 1.0 / 48.0
	} else {
		return 0.0
	}
}
