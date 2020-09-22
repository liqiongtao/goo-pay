package gooUtils

import "strings"

var (
	idcode_words = []string{
		"A", "3", "B", "5", "C", "D", "E",
		"F", "6", "G", "7", "H", "J", "K",
		"L", "8", "M", "9", "N", "P", "Q",
		"R", "S", "T", "U", "V", "W", "X",
	}
)

type idcode struct {
	BaseNum int64
	StepNum uint
	Wrods   []string
}

func (this *idcode) Code(id int64) string {
	str := ""
	words_len := len(this.Wrods)
	id = (id + this.BaseNum) << this.StepNum
	for ; id > 0; id /= int64(words_len) {
		str = this.Wrods[id%int64(words_len)] + str
	}
	return str
}

func (this *idcode) Id(code string) int64 {
	var id int64
	words_len := len(this.Wrods)
	arr := strings.Split(code, "")
	for _, word := range arr {
		for j, __word := range this.Wrods {
			if word == __word {
				id = id*int64(words_len) + int64(j)
				break
			}
		}
	}
	return (id >> this.StepNum) - this.BaseNum
}

func Id2Code(id int64) string {
	return (&idcode{BaseNum: 100000, StepNum: 9, Wrods: idcode_words}).Code(id)
}

func Code2Id(code string) int64 {
	return (&idcode{BaseNum: 100000, StepNum: 9, Wrods: idcode_words}).Id(strings.ToUpper(code))
}
