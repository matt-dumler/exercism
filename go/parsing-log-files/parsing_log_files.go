package parsinglogfiles

import (
    "regexp"
    "strings"
)

func IsValidLine(text string) bool {
    re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\].*$`)
    return re.MatchString(text)
}

func SplitLogLine(text string) []string {
    re := regexp.MustCompile(`<([~*=-])*>`)
    return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
    re := regexp.MustCompile(`".*(?i)password.*"`)

    count := 0
    for _, text := range lines {
        count += len(re.FindAllString(text, -1))
    }

    return count
}

func RemoveEndOfLineText(text string) string {
    re := regexp.MustCompile(`end-of-line\d+`)
    return strings.Join(re.Split(text, -1), "")
}

func TagWithUserName(lines []string) []string {
    re := regexp.MustCompile(`User\s+\w+`)

    for i, line := range lines {
        matches := re.FindString(line)
        if len(matches) > 0 {
            strs := strings.Split(matches, " ")
            lines[i] = "[USR] " + strs[len(strs)-1] + " " + line
        }
    }

    return lines
}
