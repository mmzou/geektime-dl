package application

import (
	"github.com/mmzou/geektime-dl/service"
)

//BuyProductAll product
func BuyProductAll() (*service.ProductAll, error) {
	return getService().BuyProductAll()
}

//BuyColumns all columns
func BuyColumns() (*service.Product, error) {
	all, err := BuyProductAll()
	return all.Columns, err
}

//BuyVideos all columns
func BuyVideos() (*service.Product, error) {
	all, err := BuyProductAll()
	return all.Videos, err
}
