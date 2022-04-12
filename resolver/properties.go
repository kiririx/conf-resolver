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
	var content []byte
	for {
		n, err := file.Read(buf[:])
		if err == io.EOF {
			// 读取结束
			break
		}
		if err != nil {
			return nil, err
		}
		content = append(content, buf[:n]...)
	}
	props := string(content)
	propArr := strings.Split(props, "\n")
	conf := make(map[string]string)
	for _, prop := range propArr {
		prop = strings.TrimSpace(prop)
		if len(prop) > 2 && !strings.HasPrefix(prop, "#") {
			key := prop[:strings.Index(prop, "=")]
			val := prop[strings.Index(prop, "=")+1:]
			conf[key] = val
		}
	}
	return conf, nil
}
