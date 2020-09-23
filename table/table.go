package table

import (
	"fmt"
	"reflect"
)

type table struct {
	elem    interface{} // origin data
	style   style       // border style enum
	border  bool        // whether display border-bottom
	headers []string    // header list
	widths  []int       // column width list
	rows    [][]string  // data list
}

// New support type: []struct, []slice...
func New(elem interface{}) *table {
	return &table{
		elem:  elem,
		style: Solid,
	}
}

// Style ...
func (t *table) Style(s style) *table {
	if s.Valid() {
		t.style = s
	}
	return t
}

// Border ...
func (t *table) Border(border bool) *table {
	t.border = border
	return t
}

// Header ...
func (t *table) Header(v []string) *table {
	if len(v) > 0 {
		t.headers = v
	}
	return t
}

// Content ...
func (t *table) Content() (content string) {
	err := t.parse()
	if err != nil {
		return err.Error()
	}

	if len(t.rows) == 0 {
		return content
	}
	b := styles[t.style]
	headerT, headerM, headerB, footer := []rune{b.DR}, []rune{b.V}, []rune{b.VR}, []rune{b.UR}
	for i, width := range t.widths {
		// width + 2, two spaces left and right
		headerT = append(headerT, []rune(t.repeat(b.H, width+2)+string(b.HD))...)
		headerB = append(headerB, []rune(t.repeat(b.H, width+2)+string(b.VH))...)
		footer = append(footer, []rune(t.repeat(b.H, width+2)+string(b.HU))...)

		if len(t.headers) > 0 {
			// with header, get actual display length
			l := width - t.getLentgh([]rune(t.headers[i])) + 1
			headerM = append(headerM, []rune(" "+t.headers[i]+t.repeat(' ', l)+string(b.V))...)
		}
	}
	headerT[len(headerT)-1], headerB[len(headerB)-1], footer[len(footer)-1] = b.DL, b.VL, b.UL

	// header area
	content += string(headerT) + "\n"
	if len(t.headers) > 0 {
		content += string(headerM) + "\n"
		content += string(headerB) + "\n"
	}

	// body area
	for i, row := range t.rows {
		body := []rune{b.V}
		for i, width := range t.widths {
			l := width - t.getLentgh([]rune(row[i])) + 1
			body = append(body, []rune(" "+row[i]+t.repeat(' ', l)+string(b.V))...)
		}
		content += string(body) + "\n"
		if t.border && i != len(t.rows)-1 {
			content += string(headerB) + "\n"
		}
	}

	// footer area
	content += string(footer)
	return content
}

// Render ...
func (t *table) Render() {
	fmt.Println(t.Content() + "\n")
}

// parse ...
func (t *table) parse() (err error) {
	value := reflect.ValueOf(t.elem)
	if value.Kind() != reflect.Slice && value.Kind() != reflect.Array {
		return errType
	}

	list := make([]interface{}, value.Len())
	for i := 0; i < value.Len(); i++ {
		list[i] = value.Index(i).Interface()
	}

	for index, item := range list {
		iv, it := reflect.ValueOf(item), reflect.TypeOf(item)
		if iv.Kind() == reflect.Ptr {
			iv = iv.Elem()
			it = it.Elem()
		}
		// check header & filed length
		headerLen := len(t.headers)
		if iv.Kind() == reflect.Struct {
			if headerLen > 0 && headerLen != iv.NumField() {
				return errHeader
			}
			row := []string{}
			for n := 0; n < iv.NumField(); n++ {
				cn := it.Field(n).Name
				cv := fmt.Sprintf("%+v", iv.FieldByName(cn).Interface())
				row = append(row, cv)
				// parse struct tag "table" to header
				if index == 0 {
					ct := it.Field(n).Tag.Get("table")
					if ct == "" {
						ct = cn
					}
					if headerLen == 0 {
						t.headers = append(t.headers, ct)
					}
					t.widths = append(t.widths, len(ct))
				}
				if t.widths[n] < len(cv) {
					t.widths[n] = len(cv)
				}
				// cal column max width
				if headerLen > 0 && len(t.headers[n]) > t.widths[n] {
					t.widths[n] = len(t.headers[n])
				}
			}
			t.rows = append(t.rows, row)
		} else if iv.Kind() == reflect.Slice || iv.Kind() == reflect.Array {
			if headerLen > 0 && headerLen != iv.Len() {
				return errHeader
			}
			row := []string{}
			for n := 0; n < iv.Len(); n++ {
				cv := fmt.Sprintf("%+v", iv.Index(n).Interface())
				row = append(row, cv)
				if index == 0 {
					t.widths = append(t.widths, len(cv))
				}
				if len(cv) > t.widths[n] {
					t.widths[n] = len(cv)
				}
				// cal column max width
				if headerLen > 0 && len(t.headers[n]) > t.widths[n] {
					t.widths[n] = len(t.headers[n])
				}
			}
			t.rows = append(t.rows, row)
		} else {
			return errType
		}
	}
	return err
}

// repeat ...
func (t *table) repeat(char rune, num int) string {
	var s = make([]rune, num)
	for i := range s {
		s[i] = char
	}
	return string(s)
}

// getLentgh get actual display length
func (t *table) getLentgh(r []rune) int {
	length := len(r)
	for _, v := range r {
		if v >= chineseCharset[0] && v <= chineseCharset[1] {
			length++
		}
	}
	return length
}
