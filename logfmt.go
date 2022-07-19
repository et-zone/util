package util

import (
	"strings"
)


func NewLogFmt() *LogFmt {
	return &LogFmt{
		strings.Builder{},
	}
}

// split = ||
type LogFmt struct {
	strings.Builder
}

func (t *LogFmt) SetFormatString(key, val string) *LogFmt {
	if t.String()!=""{
		t.WriteString("||")
	}
	t.WriteString(key)
	t.WriteString("=")
	t.WriteString(val)

	return t
}
func (t *LogFmt) SetFormatByte(key string, b []byte) *LogFmt {
	if t.String()!=""{
		t.WriteString("||")
	}
	t.WriteString(key)
	t.WriteString("=")
	t.Write(b)
	return t
}

