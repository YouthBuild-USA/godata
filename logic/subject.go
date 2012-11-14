package logic

import (
	"github.com/YouthBuild-USA/ybdata/datastore"
)

type Subject interface {
	Id() int
	Name() string
	SetName(name string)
	SortName() string
	SetSortName(string)

	Type() string

	Save() error
}

func LoadSubject(id int) (sub Subject, err error) {
	s, err := datastore.LoadSubject(id)
	if err != nil {
		return nil, err
	}
	creator := Registry.SubjectLoader(s.Type)
	return creator(s), nil
}

func CreateSubject(t string) (Subject, error) {
	creator := Registry.SubjectCreator(t)
	return creator(t), nil
}

type SubjectLoader func(*datastore.Subject) Subject
type SubjectCreator func(string) Subject
