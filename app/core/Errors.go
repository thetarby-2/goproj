package core

import "reflect"

type ParameterError struct {
	ParameterName string
	Err           error
}

func (m *ParameterError) Error() string {
	return m.ParameterName
}

type ObjectNotFound struct {
	ObjectType reflect.Type
	Err        error
	Pk         interface{}
}

func (m *ObjectNotFound) Error() string {
	return m.ObjectType.String()
}

type BindError struct {
	Err error
}

func (m *BindError) Error() string {
	return m.Err.Error()
}

type DbConstraintCheckFailed struct {
	Name string
	Err  error
}

func (m *DbConstraintCheckFailed) Error() string {
	return "Db constraint " + m.Name + " failed"
}
