package bob

import (
    "strings"
    "unicode"
)

func letterIn(remark string) bool {
    return strings.IndexFunc(remark, unicode.IsLetter) != -1  
}

func Hey(remark string) string {
    remark = strings.TrimSpace(remark)

    isSilence := len(remark) == 0
    isQuestion := strings.HasSuffix(remark, "?")
    isShouted := letterIn(remark) && remark == strings.ToUpper(remark)

	if isSilence {
        return "Fine. Be that way!"
    } else if isQuestion && isShouted {
        return "Calm down, I know what I'm doing!"
    } else if isQuestion {
        return "Sure."
    } else if isShouted {
        return "Whoa, chill out!"
    } else {
        return "Whatever."
    }
}
