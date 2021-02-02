package sexpr

import (
	"bytes"
	"fmt"
	"reflect"
	"unicode/utf8"
)

func MarshalIndent(v interface{}) ([]byte, error) {
	p := newPrinter()
	if err := pretty(p, reflect.ValueOf(v)); err != nil {
		return nil, err
	}
	return p.Bytes(), nil
}

type printer struct {
	bytes.Buffer
	depth         int
	prefixLengths []int // length list of prefix strings after new line on current depth
}

func newPrinter() *printer {
	return &printer{prefixLengths: []int{0}}
}
func (p *printer) string(str string) {
	p.WriteString(str)
	p.prefixLengths[p.depth] += utf8.RuneCountInString(str)
}
func (p *printer) stringf(format string, args ...interface{}) {
	p.string(fmt.Sprintf(format, args...))
}
func (p *printer) nest() {
	p.depth++
	p.prefixLengths = append(p.prefixLengths, 0) // push
}
func (p *printer) unnest() {
	p.depth--
	p.prefixLengths = p.prefixLengths[:len(p.prefixLengths)-1] // pop
}
func (p *printer) newLine() {
	p.WriteString("\n")
	p.prefixLengths[p.depth] = 0
	numIndent := 0
	for _, l := range p.prefixLengths {
		numIndent += l
	}
	fmt.Fprintf(&p.Buffer, "%*s", numIndent, "")
}

func pretty(p *printer, v reflect.Value) error {
	switch v.Kind() {
	case reflect.Invalid:
		p.string("nil")

	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		p.stringf("%d", v.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		p.stringf("%d", v.Uint())

	case reflect.String:
		p.stringf("%q", v.String())

	case reflect.Array, reflect.Slice: // (value ...)
		p.string("(")
		p.nest()
		for i := 0; i < v.Len(); i++ {
			if err := pretty(p, v.Index(i)); err != nil {
				return err
			}
			if i < v.Len()-1 {
				p.newLine()
			}
		}
		p.string(")")
		p.unnest()

	case reflect.Struct: // ((name value ...)
		p.string("(")
		p.nest()
		for i := 0; i < v.NumField(); i++ {
			p.string("(")
			p.string(v.Type().Field(i).Name)
			p.string(" ")
			if err := pretty(p, v.Field(i)); err != nil {
				return err
			}
			p.string(")")
			if i < v.NumField()-1 {
				p.newLine()
			}
		}
		p.string(")")
		p.unnest()

	case reflect.Map: // ((key value ...)
		p.string("(")
		p.nest()
		for i, key := range v.MapKeys() {
			p.string("(")
			if err := pretty(p, key); err != nil {
				return err
			}
			p.string(" ")
			if err := pretty(p, v.MapIndex(key)); err != nil {
				return err
			}
			p.string(")")
			if i < len(v.MapKeys())-1 {
				p.newLine()
			}
		}
		p.string(")")
		p.unnest()

	case reflect.Ptr:
		return pretty(p, v.Elem())

	default: // float, complex, bool, chan, func, interface
		return fmt.Errorf("unsupported type: %s", v.Type())
	}
	return nil
}
