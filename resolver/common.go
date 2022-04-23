package resolver

import "strings"

func filterStr(str *string) {
	strings.TrimSpace(*str)
}
