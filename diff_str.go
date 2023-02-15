package main

import (
	"fmt"

	"github.com/sergi/go-diff/diffmatchpatch"
)

const (
	text1 = "患者头晕乏力，胸痛胸闷无明显好转，睡眠佳，PE:T36.5°℃，R18次/分、P65次/分、BP140/80mmHg，神清，精神欠佳，瞳孔等大等圆，颈软无抵抗，心率65次/分，律齐，心浊音界正常范围，两肺未闻及干湿性罗音及哮鸣音，腹软无压痛，科主任对医嘱仔细查对后结合病史认为诊断明确，"
	text2 = "患者不晕，胸痛胸闷无明显好转，睡眠佳，PE:T36.5°℃，R18次/分、P65次/分、BP140/80mmHg，神清，精神欠佳，瞳孔等大等圆，颈软无抵抗，心率65次/分，律齐，心浊音界正常范围，两肺未闻及干湿性罗音及哮鸣音，腹软无压痛，科主任对医嘱仔细查对后结合病史认为诊断明确，"
)

func main() {
	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(text1, text2, true)
	fmt.Println(dmp.DiffPrettyHtml(diffs))
	for _, v := range diffs {
		fmt.Println(v)
	}

}
