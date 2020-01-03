// 模板配置
// author: baoqiang
// time: 2019-08-27 14:45
package network

import (
	"fmt"
	"os"
	"strings"

	//"text/template"
	"html/template"
)

type Person struct {
	Name   string
	Age    int
	Emails []string
	Jobs   []*Job
}

type Job struct {
	Employer string
	Role     string
}

const templ = `The name is {{.Name}}.
The age is {{.Age}}.
{{range .Emails}}
	An email is {{. | emailExpand}}
{{end}}

{{with .Jobs}}
	{{range .}}
	An employer is {{.Employer}}
	and the role is {{.Role}}
	{{end}}
{{end}}
`

// 自定义变量
const templ2 = `{{$name := .Name}}.
{{range .Emails}}
	name is {{$name}}, email is {{. | emailExpand}}
{{end}}
`

// 条件判断打印一下json
const templ3 = `{"Name": "{{.Name}}",
"Emails": [
	{{range $index,$ele := .Emails}}
		{{if $index}}
			, "{{$ele}}"
		{{else}}
			"{{$ele}}"
		{{end}}
	{{end}}
]
}`

// html escape
func RunTempl() {
	person := Person{
		Name:   "jan<hello",
		Age:    50,
		Emails: []string{"jan@newmarch.name", "jan.newmarch@gmail.com"},
		Jobs: []*Job{
			{Employer: "Monash", Role: "Honorary"},
			{Employer: "Box Hill", Role: "Head of HE"},
		},
	}

	t := template.New("Person template")

	// 扩展一个自定义函数
	t = t.Funcs(template.FuncMap{"emailExpand": EmailExpander})

	t, err := t.Parse(templ3)
	checkError(err)

	err = t.Execute(os.Stdout, person)
	checkError(err)

}

func EmailExpander(args ...interface{}) string {
	ok := false
	var s string
	if len(args) == 1 {
		s, ok = args[0].(string)
	}
	if !ok {
		s = fmt.Sprint(args...)
	}

	substrs := strings.Split(s, "@")
	if len(substrs) != 2 {
		return s
	}

	return substrs[0] + " at " + substrs[1]
}
