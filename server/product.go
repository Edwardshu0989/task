package server

import (
	"awesomeProject/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func (s *Server) AddProduct(c *gin.Context) string {
	req := &model.AddProduct{}
	if err := c.ShouldBindJSON(&req); err != nil {
		return fmt.Sprintf("添加商品时，商品请求体解析错误")
	}
	product := &model.Product{
		Model:       gorm.Model{},
		ProductName: req.ProductName,
	}
	if err := s.dao.AddProduct(product); err != nil {
		return fmt.Sprintf("添加商品时，商品入库出错")
	}
	return fmt.Sprintf("商品入库成功")
}

func (s *Server) GetProduct(c *gin.Context) string {
	req := &model.AddProduct{}
	if err := c.ShouldBindJSON(&req); err != nil {
		return fmt.Sprintf("查询商品时，商品请求体解析错误")
	}
	product := s.dao.GetProduct(req.ProductName)
	if product == nil {
		return fmt.Sprintf("查询商品时，商品查询出错")
	}
	return fmt.Sprintf("商品信息 %v,id %v", product.CreatedAt, product.ID)
}
