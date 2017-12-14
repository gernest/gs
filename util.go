package goss

import (
	"sort"
	"strings"

	"github.com/gernest/classnames"
)

func IndentStr(src string, indent int) string {
	r := ""
	for i := 0; i < indent; i++ {
		r += "  "
	}
	return r + src
}

type Options struct {
	Indent     int
	ClassNamer func(string) string
	ClassMap   ClassMap
}

// NewOpts returns a new *Options instance with non nil CLassMap
func NewOpts() *Options {
	return &Options{
		ClassMap: make(ClassMap),
	}
}

// ClassMap is a map of selectors to generated classname
type ClassMap map[string]string

// Classes returns a string representation of css classes stored in this map.
func (c ClassMap) Classes() string {
	var v []interface{}
	for _, i := range c {
		v = append(v, i)
	}
	return classnames.Join(v...)
}

func (c ClassMap) Merge(cm ClassMap) {
	for k, v := range cm {
		c[k] = v
	}
}

type CSSTree struct {
	Selector string
	Parent   *CSSTree
	Children TreeList
	Text     string
}

type TreeList []*CSSTree

func (t TreeList) Len() int {
	return len(t)
}
func (t TreeList) Less(i, j int) bool {
	return t[i].Text < t[j].Text
}

func (t TreeList) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

// ToCSS returns css string representation for style
func ToCSS(style *Style, opts *Options) string {
	r := ""
	if style == nil {
		return r
	}
	nested := ""
	indent := opts.Indent
	indent++
	for k, v := range style.Fallbacks {
		if k == 0 {
			r = IndentStr(v.ToString(opts), indent)
		} else {
			r += "\n" + IndentStr(v.ToString(opts), indent)
		}
	}
	for _, v := range style.Rules {
		if vt, ok := v.(*Style); ok {
			if style.Selector == "root" || style.Selector == "" {
				if opts.ClassNamer != nil && opts.ClassMap != nil {
					n := opts.ClassNamer(vt.Selector)
					opts.ClassMap[vt.Selector] = n
					vt.Selector = n
				}
			}
			if nested == "" {
				nested = ToCSS(vt, opts)
			} else {
				nested += "\n" + ToCSS(vt, opts)
			}
		} else {
			if style.Selector == "" {
				if r == "" {
					r = v.ToString(opts)
				} else {
					r += "\n" + v.ToString(opts)
				}
			} else {
				if r == "" {
					r = IndentStr(v.ToString(opts), indent)
				} else {
					r += "\n" + IndentStr(v.ToString(opts), indent)
				}
			}
		}
	}
	indent--
	result := r
	if style.Selector != "" {
		result = IndentStr(style.Selector+" {\n"+r, indent) + "\n" + IndentStr("}", indent)
	}
	if nested != "" {
		return result + "\n" + nested
	}
	return result
}

func FormatCSS(style *Style, parent *CSSTree, opts *Options) *CSSTree {
	var fallback TreeList
	for _, v := range style.Fallbacks {
		fallback = append(fallback, &CSSTree{
			Parent: parent,
			Text:   v.ToString(opts),
		})
	}
	current := &CSSTree{
		Parent:   parent,
		Selector: style.Selector,
	}
	for _, v := range style.Rules {
		switch e := v.(type) {
		case *Style:
			current.Children = append(current.Children, FormatCSS(e, current, opts))
		default:
			current.Children = append(current.Children, &CSSTree{
				Parent: current,
				Text:   v.ToString(opts),
			})
		}

	}
	sort.Sort(fallback)
	current.Children = append(current.Children, fallback...)
	return current
}

func hasPrefix(str string, prefix string) bool {
	return strings.HasPrefix(str, prefix)
}

func replace(str string, old, new string) string {
	return strings.Replace(str, old, new, -1)
}

func (c *CSSTree) Print(depth int) string {
	var values []string
	if c.Selector != "" {
		if len(c.Children) > 0 {
			o := c.Selector + "{"
			for _, v := range c.Children {
				if v.Selector != "" {
					values = append(values, v.Print(depth))
				} else {
					o += "\n" + v.Print(depth+2)
				}
			}
			o += "\n}"
			values = append(values, o)
		}
	} else if c.Text != "" {
		values = append(values, IndentStr(c.Text, depth))
	} else {
		for _, v := range c.Children {
			values = append(values, v.Print(depth))
		}
	}
	sort.Strings(values)
	return strings.Join(values, "\n")
}
