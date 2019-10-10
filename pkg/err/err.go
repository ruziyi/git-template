package err

import "strings"

type ErrorList []error

func (e ErrorList) Error() string {
	var buf strings.Builder
	for _, v := range e {
		buf.WriteString(v.Error())
		buf.WriteString("\n")
	}
	return buf.String()
}
