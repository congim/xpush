package basic

import (
	"bytes"
	"compress/flate"
	"crypto/md5"
	"encoding/gob"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/url"
	"os"
	"runtime"
	"strconv"
	"time"
)

// Flate flate 压缩函数
func Flate(data []byte) ([]byte, error) {
	var res bytes.Buffer
	fn, err := flate.NewWriter(&res, 7)
	if err != nil {
		return nil, fmt.Errorf(" compress Flate NewWriter failed, err message is %s", err)
	}

	_, err = fn.Write(data)
	if err != nil {
		fn.Close()
		return nil, fmt.Errorf("[Error] compress Flate Write failed, err message is %s", err)
	}
	fn.Close()

	return res.Bytes(), nil
}

// UnFlate flate 解压缩函数
func UnFlate(data []byte) []byte {
	b := bytes.NewReader(data)
	var out bytes.Buffer
	r := flate.NewReader(b)
	io.Copy(&out, r)
	return out.Bytes()
}

// StringTobyte string切片转[][]byte
func StringTobyte(str []string) [][]byte {
	if str == nil {
		return nil
	}
	var b [][]byte
	for _, v := range str {
		b = append(b, []byte(v))
	}
	return b
}

// InterfaceToByte interface转[]byte
func InterfaceToByte(data interface{}) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buf)
	err := enc.Encode(data)
	if err != nil {
		fmt.Printf("[Error] InterfaceToByte err, err message is %s", err)
		return nil, err
	}
	return buf.Bytes(), nil
}

// ByteToInterface  []byte转 interface
func ByteToInterface(data []byte, to interface{}) error {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	return dec.Decode(to)
}

// Substr 截取字符串 start 起点下标 length 需要截取的长度
func Substr(str string, start int, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length
	if start > end {
		start, end = end, start
	}
	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}

	return string(rs[start:end])
}

// 二分法查找
// 切片s是升序的
// k为待查找的整数
// 如果查到有就返回对应下标,
// 没有就返回-1
func BinarySearch(s []string, k string) int {
	lo, hi := 0, len(s)-1
	for lo <= hi {
		m := (lo + hi) >> 1
		if s[m] < k {
			lo = m + 1
		} else if s[m] > k {
			hi = m - 1
		} else {
			return m
		}
	}
	return -1
}

var (
	dunno     = []byte("???")
	centerDot = []byte("·")
	dot       = []byte(".")
	slash     = []byte("/")
)

// Mystack stack returns a nicely formated stack frame, skipping skip frames
func Mystack(skip int) []byte {
	buf := new(bytes.Buffer) // the returned data
	// As we loop, we open files and read them. These variables record the currently
	// loaded file.
	var lines [][]byte
	var lastFile string
	for i := skip; ; i++ { // Skip the expected number of frames
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		// Print this much at least.  If we can't find the source, it won't show.
		fmt.Fprintf(buf, "%s:%d (0x%x)\n", file, line, pc)
		if file != lastFile {
			data, err := ioutil.ReadFile(file)
			if err != nil {
				continue
			}
			lines = bytes.Split(data, []byte{'\n'})
			lastFile = file
		}
		fmt.Fprintf(buf, "\t%s: %s\n", function(pc), source(lines, line))
	}
	return buf.Bytes()
}

// function returns, if possible, the name of the function containing the PC.
func function(pc uintptr) []byte {
	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return dunno
	}
	name := []byte(fn.Name())
	// The name includes the path name to the package, which is unnecessary
	// since the file name is already included.  Plus, it has center dots.
	// That is, we see
	//	runtime/debug.*T·ptrmethod
	// and want
	//	*T.ptrmethod
	// Also the package path might contains dot (e.g. code.google.com/...),
	// so first eliminate the path prefix
	if lastslash := bytes.LastIndex(name, slash); lastslash >= 0 {
		name = name[lastslash+1:]
	}
	if period := bytes.Index(name, dot); period >= 0 {
		name = name[period+1:]
	}
	name = bytes.Replace(name, centerDot, dot, -1)
	return name
}

// source returns a space-trimmed slice of the n'th line.
func source(lines [][]byte, n int) []byte {
	n-- // in stack trace, lines are 1-indexed but our array is 0-indexed
	if n < 0 || n >= len(lines) {
		return dunno
	}
	return bytes.TrimSpace(lines[n])
}

// URLDecoded url 解码
func URLDecoded(str string) (string, error) {
	return url.QueryUnescape(str)
}

// URLEncoded url 编码
func URLEncoded(str string) (string, error) {
	u, err := url.Parse(str)
	if err != nil {
		return "", err
	}
	return u.String(), nil
}

// Token 内部用户令牌
func Token(sid string) string {
	nano := time.Now().UnixNano()
	rand.Seed(nano)
	rndNum := rand.Int63()
	return Md5(Md5(sid) + Md5(strconv.FormatInt(nano, 10)) + Md5(strconv.FormatInt(rndNum, 10)))
}

// Md5 md5
func Md5(text string) string {
	hashMd5 := md5.New()
	io.WriteString(hashMd5, text)
	return fmt.Sprintf("%x", hashMd5.Sum(nil))
}

func Exist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}
