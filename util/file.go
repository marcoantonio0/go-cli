package util

import (
	"fmt"
	"os"
	textTemplate "text/template"
)

func CreateFile(directoryPath, fileName, fileData string) (*os.File, int, error) {
	var err error
	var file *os.File
	if directoryPath != "" {
		fullPath := fmt.Sprintf("%s/%s", directoryPath, fileName)

		file, err = os.Create(fullPath)
		if err != nil {
			return nil, 0, err
		}
	} else {
		file, err = os.Create(fileData)
		if err != nil {
			return nil, 0, err
		}
	}

	n, err := file.WriteString(fileData)
	if err != nil {
		return nil, 0, err
	}

	return file, n, nil
}

func CreateGoFileByTemplate(directoryPath, fileName, template string, data interface{}) (*os.File, error) {
	var file *os.File
	tmpl, err := textTemplate.New(fileName).Parse(template)
	if err != nil {
		panic(err)
	}

	if directoryPath != "" {
		fullPath := fmt.Sprintf("%s/%s.go", directoryPath, fileName)

		file, err = os.Create(fullPath)
		if err != nil {
			return nil, err
		}
	} else {
		file, err = os.Create(fileName + ".go")
		if err != nil {
			return nil, err
		}
	}
	defer file.Close()

	err = tmpl.Execute(file, data)
	if err != nil {
		return nil, err
	}

	fmt.Printf("%s.go has been created successfully.\n", fileName)
	return file, nil
}
