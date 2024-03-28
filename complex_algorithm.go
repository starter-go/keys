package keys

import "strings"

// ComplexAlgorithm 是一组由多个单词构成的字符串，表示一组由多种算法构成的密码学套件
// 它在形式上，可以是：
// 驼峰命名，如： "SHA256withRSA"；
// 下划线分隔的小写字符，如： "sha256_with_rsa"；
// 下划线分隔的大写字符，如： "SHA256_WITH_RSA"；
// 空格分隔的小写（或大写）字符，如： "SHA256 WITH RSA"；
// 等。。。
// 默认形式为："sha256_with_rsa"
type ComplexAlgorithm string

// NewComplexAlgorithm 新建算法字符串
func NewComplexAlgorithm(words []string) ComplexAlgorithm {
	b := new(strings.Builder)
	for _, item := range words {
		item = strings.TrimSpace(item)
		if item == "" {
			continue
		}
		if b.Len() > 0 {
			b.WriteRune('_')
		}
		b.WriteString(item)
	}
	str := b.String()
	str = strings.ToLower(str)
	return ComplexAlgorithm(str)
}

func (a ComplexAlgorithm) String() string {
	return string(a)
}

// Normalize 标准化算法字符串
func (a ComplexAlgorithm) Normalize() ComplexAlgorithm {
	items := a.ToArray()
	return NewComplexAlgorithm(items)
}

// ToArray 把这个 ComplexAlgorithm 解析为数组
func (a ComplexAlgorithm) ToArray() []string {
	str := a.String()
	if strings.Contains(str, "with") {
		str = strings.ReplaceAll(str, "with", "_")

	}
	if strings.Contains(str, "and") {
		str = strings.ReplaceAll(str, "and", "_")
	}
	chs := []rune(str)
	parser := new(complexAlgorithmParser)
	for _, ch := range chs {
		parser.add(ch)
	}
	return parser.complete()
}

////////////////////////////////////////////////////////////////////////////////

type complexAlgorithmParser struct {
	items  []string
	buffer strings.Builder
}

func (inst *complexAlgorithmParser) add(ch rune) {
	isMark := false
	if ch == ' ' {
		isMark = true
	} else if '0' <= ch && ch <= '9' {
		inst.buffer.WriteRune(ch)
	} else if 'a' <= ch && ch <= 'z' {
		inst.buffer.WriteRune(ch)
	} else if 'A' <= ch && ch <= 'Z' {
		inst.buffer.WriteRune(ch)
	} else {
		isMark = true
	}
	if isMark {
		inst.flush()
	}
}

func (inst *complexAlgorithmParser) flush() {
	if inst.buffer.Len() > 0 {
		str := inst.buffer.String()
		inst.buffer.Reset()
		inst.items = append(inst.items, str)
	}
}

func (inst *complexAlgorithmParser) complete() []string {
	inst.flush()
	return inst.items
}
