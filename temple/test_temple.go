package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
	"text/template"
)

const testTemple = `
<html>
<head>
<title>Personal information</title>
</head>

<body style="text-align:center">
<h3>Person general infor</h3>
<hr>
<ul>
<li>Name: {{.Name}}:<p>
<li>Id: {{.Id}} <p>
<li>Country: {{.Country}}
</ul>
<hr>
</body>
</html>
`

type Person struct {
	Id      int
	Name    string
	Country string
}

func (p *Person) execute() ([]byte, error)  {
	buf := new(bytes.Buffer)
	tmpl, err := template.New("html").Parse(strings.TrimSpace(testTemple))
	if err != nil {
		return nil, err
	}
	if err := tmpl.Execute(buf, p); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func main() {
	lookup()
	p:= Person{
		Id:1,
		Name: "dandy",
		Country: "dandyking",
	}
	body, _:=p.execute()
	ioutil.WriteFile("./temple/temple.html", body, 0644)
}

func lookup() {
	f, err := exec.LookPath("lss")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(f) //  /bin/ls
}