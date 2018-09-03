package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

const vietBaseMarks = `
 áàảãạ
ăắằẳẵặ
âấầẩẫậ
 éèẻẽẹ
êếềểễệ
 íìỉĩị
 óòỏõọ
ôốồổỗộ
ơờớởỡợ
 úùủũụ
ưứừửữự
 ýỳỷỹỵ
đ
`

var (
	vietWordRegex          *regexp.Regexp
	vietWordToneStartRegex *regexp.Regexp
)

func init() {
	vbm := strings.Replace(vietBaseMarks, " ", "", -1)
	vbm = strings.Replace(vbm, "\n", "", -1)
	VBM := strings.ToUpper(vbm)

	vietWordPattern := fmt.Sprintf(`[a-zA-Z]*[%[1]s%[2]s][a-zA-Z]*`, vbm, VBM)
	vietWordToneStartPattern := fmt.Sprintf(`[%[1]s%[2]s][a-zA-Z]*`, vbm, VBM)

	var err error
	vietWordRegex, err = regexp.Compile(vietWordPattern)
	if err != nil {
		log.Fatal(err)
	}

	vietWordToneStartRegex, err = regexp.Compile(vietWordToneStartPattern)
	if err != nil {
		log.Fatal(err)
	}
}

func extractVietWord(s string, toMap map[string]bool) {
	words := vietWordRegex.FindAllString(s, -1)
	for _, w := range words {
		if len(w) > 7 {
			continue
		}
		wlower := strings.ToLower(w)
		toMap[wlower] = true
	}
}
