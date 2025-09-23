package sheet

import "strings"


func sheetNameOnly(a1 string) string {
    // "'Expected value'!A161:E" → Expected value; "Sheet1!A:E" → Sheet1; "Sheet1" → Sheet1
    s := a1
    if i := strings.Index(s, "!"); i >= 0 { s = s[:i] }
    s = strings.Trim(s, "'")
    return s
}