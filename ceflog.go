package ceflog

import (
	"fmt"
	"io"
	"strings"
)

const (
	cefVersion = 0
)

type Logger struct {
	vendor  string
	product string
	version string

	w io.Writer
}

func New(w io.Writer, vendor, product, version string) *Logger {
	return &Logger{
		vendor:  prefixEscape(vendor),
		product: prefixEscape(product),
		version: prefixEscape(version),

		w: w,
	}
}

func (l *Logger) Event(name, signature string, sev Severity, ext Extension) {
	fmt.Fprintf(l.w, "CEF:%d|%s|%s|%s|%s|%s|%d|%s\n",
		cefVersion,
		l.vendor,
		l.product,
		l.version,
		prefixEscape(signature),
		prefixEscape(name),
		sev,
		ext,
	)
}

var prefixEscaper = strings.NewReplacer(
	"\n", `\n`,
	`\`, `\\`,
	`|`, `\|`,
)

var extensionEscaper = strings.NewReplacer(
	`\`, `\\`,
	"\n", `\n`,
	`=`, `\=`,
)

func prefixEscape(input string) string {
	return prefixEscaper.Replace(input)
}

func extensionEscape(input string) string {
	return extensionEscaper.Replace(input)
}

type Severity int

func Sev(s int) Severity {
	if s < 0 {
		s = 0
	} else if s > 10 {
		s = 10
	}

	return Severity(s)
}

type Pair struct {
	Key   string
	Value string
}

type Extension []Pair

func (e Extension) String() string {
	var pairs []string

	for _, p := range e {
		key := p.Key
		value := extensionEscape(p.Value)

		pairs = append(pairs, fmt.Sprintf("%s=%s", key, value))
	}

	return strings.Join(pairs, " ")
}

func Ext(pairs ...string) Extension {
	if len(pairs)%2 != 0 {
		panic("pairs length must be even!")
	}

	var e Extension

	for i := 0; i < len(pairs); i = i + 2 {
		e = append(e, Pair{
			Key:   pairs[i],
			Value: pairs[i+1],
		})
	}

	return e
}
