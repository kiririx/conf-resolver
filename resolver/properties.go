package resolver

import (
	"io"
	"os"
	"strings"
)

func ResolveProperties(file *os.File) (map[string]string, error) {
	_, err := file.Stat()
	if err != nil {
		return nil, err
	}
	var buf [128]byte
	conf := make(map[string]string)
	for {
		n, err := file.Read(buf[:])
		if err == io.EOF {
			// 读取结束
			break
		}
		if err != nil {
			return nil, err
		}
		lineContent := string(buf[:n])
		prop := strings.TrimSpace(lineContent)
		key := prop[:strings.Index(prop, "=")]
		val := prop[strings.Index(prop, "=")+1:]
		conf[key] = val
	}
	return conf, nil
}
