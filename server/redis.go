package server

import (
	"fmt"
)

func (s *Server) AddRedis() string {
	resp := s.dao.AddData()
	if resp == false {
		return fmt.Sprintf("%s", "redis调用出错")
	}
	return fmt.Sprintf("%s", "redis调用成功")
}

func (s *Server) GetRedisData() string {
	resp, err := s.dao.GetData()
	if err != nil {
		return fmt.Sprintf("%s", "redis调用出错")
	}
	return fmt.Sprintf("%s", resp)
}
