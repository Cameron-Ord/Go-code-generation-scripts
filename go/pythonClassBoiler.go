package main

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
)

func main() {
	fmt.Print("Enter a value for InputVariable: ")
	var inputVariable string
	_, err := fmt.Scanln(&inputVariable)
	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}

	codeTemplate := `
func main() {
    // Generated code using input variable: {{ .InputVariable }}
}`

	data := struct {
		InputVariable string
	}{
		InputVariable: inputVariable,
	}

	tmpl, _ := template.New("code").Parse(codeTemplate)

	var generatedCode bytes.Buffer
	tmpl.Execute(&generatedCode, data)
	generatedCodeStr := generatedCode.String()
	fmt.Println(generatedCodeStr)

}
