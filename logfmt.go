package util

import (
	"strconv"
	"strings"
)

func NewLogBuilder() LogBuilder {
	return LogBuilder{
		strings.Builder{},
	}
}

// split = ||
type LogBuilder struct {
	builder strings.Builder
}

func (t *LogBuilder)SetString(key string,val string)*LogBuilder{
	if t.builder.String()!=""{
		t.builder.WriteString("||")
	}
	t.builder.WriteString(key)
	t.builder.WriteString("=")
	t.builder.WriteString(val)
	return t
}
func (t *LogBuilder)SetInt(key string,val int)*LogBuilder{
	if t.builder.String()!=""{
		t.builder.WriteString("||")
	}
	t.builder.WriteString(key)
	t.builder.WriteString("=")
	t.builder.WriteString(strconv.Itoa(val))
	return t
}

func (t *LogBuilder)SetInt8(key string,val int8)*LogBuilder{
	if t.builder.String()!=""{
		t.builder.WriteString("||")
	}
	t.builder.WriteString(key)
	t.builder.WriteString("=")
	t.builder.WriteString(strconv.FormatInt(int64(val),10))

	return t
}

func (t *LogBuilder)SetInt32(key string,val int32)*LogBuilder{
	if t.builder.String()!=""{
		t.builder.WriteString("||")
	}
	t.builder.WriteString(key)
	t.builder.WriteString("=")
	t.builder.WriteString(strconv.FormatInt(int64(val),10))

	return t
}

func (t *LogBuilder)SetInt64(key string,val int64)*LogBuilder{
	if t.builder.String()!=""{
		t.builder.WriteString("||")
	}
	t.builder.WriteString(key)
	t.builder.WriteString("=")
	t.builder.WriteString(strconv.FormatInt(val,10))

	return t
}

func (t *LogBuilder)SetUint(key string,val uint)*LogBuilder{
	if t.builder.String()!=""{
		t.builder.WriteString("||")
	}
	t.builder.WriteString(key)
	t.builder.WriteString("=")
	t.builder.WriteString(strconv.FormatUint(uint64(val),10))
	return t
}

func (t *LogBuilder)SetUint8(key string,val uint8)*LogBuilder{
	if t.builder.String()!=""{
		t.builder.WriteString("||")
	}
	t.builder.WriteString(key)
	t.builder.WriteString("=")
	t.builder.WriteString(strconv.FormatUint(uint64(val),10))
	return t
}

func (t *LogBuilder)SetUint32(key string,val uint32)*LogBuilder{
	if t.builder.String()!=""{
		t.builder.WriteString("||")
	}
	t.builder.WriteString(key)
	t.builder.WriteString("=")
	t.builder.WriteString(strconv.FormatUint(uint64(val),10))
	return t
}

func (t *LogBuilder)SetUint64(key string,val uint64)*LogBuilder{
	if t.builder.String()!=""{
		t.builder.WriteString("||")
	}
	t.builder.WriteString(key)
	t.builder.WriteString("=")
	t.builder.WriteString(strconv.FormatUint(val,10))
	return t
}

func (t *LogBuilder)SetFloat32(key string,val float32)*LogBuilder{
	if t.builder.String()!=""{
		t.builder.WriteString("||")
	}
	t.builder.WriteString(key)
	t.builder.WriteString("=")
	t.builder.WriteString(strconv.FormatFloat(float64(val), 'f', -1, 32))
	return t
}

func (t *LogBuilder)SetFloat64(key string,val float64)*LogBuilder{
	if t.builder.String()!=""{
		t.builder.WriteString("||")
	}
	t.builder.WriteString(key)
	t.builder.WriteString("=")
	t.builder.WriteString(strconv.FormatFloat(val, 'f', -1, 32))
	return t
}

func (t *LogBuilder)SetBool(key string,val bool)*LogBuilder{
	if t.builder.String()!=""{
		t.builder.WriteString("||")
	}
	t.builder.WriteString(key)
	t.builder.WriteString("=")
	t.builder.WriteString(strconv.FormatBool(val))
	return t
}

func (t *LogBuilder) SetBytes(key string, b []byte) *LogBuilder {
	if t.builder.String()!=""{
		t.builder.WriteString("||")
	}
	t.builder.WriteString(key)
	t.builder.WriteString("=")
	t.builder.Write(b)
	return t
}

func (t *LogBuilder) Close(){
	t.builder.Reset()
}

func (t *LogBuilder) String()string{
	return t.builder.String()
}