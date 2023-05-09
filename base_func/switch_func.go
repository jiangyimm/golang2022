package main

import "fmt"

func main() {
	num := 99
	switch num {
	case 99:
		fallthrough
	case 98:
		fmt.Println("num is 99 or 98")
	default:
		fmt.Println("num is not 99 or 98")
	}

	switch {
	case num < 100:
	case num > 100:
	}

	season := Season(13)
	fmt.Println(season)
}

//获取季节
func Season(month int8) (season string) {
	switch {
	case month <= 3:
		season = "春季"
	case month <= 6:
		season = "夏季"
	case month <= 9:
		season = "秋季"
	case month <= 12:
		season = "冬季"
	default:
		season = "未知"
	}
	return season
}
