package server

import dao2 "awesomeProject/dao"

type Server struct {
	dao *dao2.Db
}

func New(db *dao2.Db) (s *Server) {
	s = &Server{
		dao: db,
	}
	// 检查创建数据库表
	db.CreateTable()

	return
}
