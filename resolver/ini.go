package resolver

import (
	"conf-resolver/format"
	"io"
	"os"
	"regexp"
	"strings"
)

func ResolveIni(file *os.File) (format.Ini, error) {
	NilIni := format.Ini{}
	_, err := file.Stat()
	if err != nil {
		return NilIni, err
	}
	var buf [128]byte
	sections := make(map[string]*format.Properties)
	currentSection := ""
	for {
		n, err := file.Read(buf[:])
		if err == io.EOF {
			// 读取结束
			break
		}
		if err != nil {
			return NilIni, err
		}
		lineContent := string(buf[:n])
		filterStr(&lineContent)
		if lineContent != "" {
			if ok, _ := regexp.MatchString("(\\[).*(\\])", lineContent); ok {
				// sections
				lineContent = strings.TrimPrefix(lineContent, "[")
				lineContent = strings.TrimSuffix(lineContent, "]")
				sections[lineContent] = &format.Properties{}
				currentSection = lineContent
			} else if ok, _ := regexp.MatchString(".*=.*", lineContent); ok {
				// k v
				contentArr := strings.Split(lineContent, "=")
				sections[currentSection].Set(contentArr[0], contentArr[1])
			}
		}
	}
	iniFile := format.Ini{Sections: sections}
	return iniFile, nil
}
