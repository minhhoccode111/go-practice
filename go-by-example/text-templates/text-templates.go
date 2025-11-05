package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	t1 := template.New("t1")
	fmt.Println(t1)

	t1, err := t1.Parse("Value is {{.}}\n")
	if err != nil {
		panic(err)
	}
	fmt.Println(t1)

	t1 = template.Must(t1.Parse("Value: {{.}}\n"))
	fmt.Println(t1)

	t1.Execute(os.Stdout, "some text")
	t1.Execute(os.Stdout, 5)
	t1.Execute(os.Stdout, []string{
		"Go",
		"Rust",
		"C++",
		"C#",
	})

	Create := func(name, t string) *template.Template {
		return template.Must(template.New(name).Parse(t))
	}

	t2 := Create("t2", "Name: {{.Name}}\n")

	fmt.Println(t2) // or we can use t1 to print to stdout

	t2.Execute(os.Stdout, struct{ Name string }{"tai vi sao"})
	t2.Execute(os.Stdout, map[string]string{
		"Name": "tai vi sao",
	})

	t3 := Create("t3", "{{if . -}} yes {{else -}} no {{end}}\n")
	t3.Execute(os.Stdout, "not empty")
	t3.Execute(os.Stdout, "")

	t4 := Create("t4", "Range: {{range .}}{{.}} {{end}}\n")
	t4.Execute(os.Stdout, []string{
		"Go",
		"Rust",
		"C++",
		"C#",
	})

	/*
	   $ go run text-templates.go
	   &{t1 <nil> 0xc0000c0000  }
	   &{t1 0xc0000ce000 0xc0000c0000  }
	   &{t1 0xc0000ce120 0xc0000c0000  }
	   Value: some text
	   Value: 5
	   Value: [Go Rust C++ C#]
	   &{t2 0xc0000ce240 0xc0000c0140  }
	   Name: tai vi sao
	   Name: tai vi sao
	   yes
	   no
	   Range: Go Rust C++ C#
	*/
}
