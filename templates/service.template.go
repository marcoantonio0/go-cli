package templates

var ServiceTemplate string = `
package services

import (
	"{{.ProjectName}}/app/repositories"
)

type {{.Name}}Service struct {
	{{.NameUnderline}}Repository *repositories.{{.Name}}Repository
}

func New{{.Name}}Service(userRepository *repository.{{.Name}}Repository) *UserService {
	return &{{.Name}}Service{
		{{.NameUnderline}}Repository: {{.NameUnderline}}Repository,
	}
}

func (s *{{.Name}}Service) Create{{.Name}}({{.NameUnderline}} repository.{{.Name}}) (any, error) {
	result, err := s.{{.NameUnderline}}Repository.Create{{.Name}}({{.NameUnderline}})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *{{.Name}}Service) ListAll{{.Name}}s() ([]repository.{{.Name}}, error) {
	result, err := s.{{.NameUnderline}}Repository.List()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *{{.Name}}Service) FindById(id string) (*repository.{{.Name}}, error) {
	result, err := s.{{.NameUnderline}}Repository.FindById(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

}
`
