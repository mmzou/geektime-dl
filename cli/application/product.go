package application

import (
	"github.com/mmzou/geektime-dl/service"
)

//ProductAll product
func ProductAll() (*service.ProductAll, error) {
	return getService().ProductAll()
}

//Columns all columns
func Columns() (*service.Product, error) {
	all, err := ProductAll()
	return all.Column, err
}

//Courses all columns
func Courses() (*service.Product, error) {
	all, err := ProductAll()
	return all.Course, err
}
