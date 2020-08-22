package server

import (
	dao2 "awesomeProject/dao"
)

type Server struct {
	dao *dao2.Dao
}

func New(db *dao2.Dao) (s *Server) {
	s = &Server{dao: db}
	return
}
