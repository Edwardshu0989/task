package dl

import (
	"awesomeProject/server"
	"github.com/gin-gonic/gin"
)

type App struct {
	h *gin.Engine
	s *server.Server
}

//
func NewApp(h *gin.Engine, s *server.Server) (app *App, cf func(), err error) {
	app = &App{
		h: h,
		s: s,
	}
	cf = Close
	err = nil
	return
}

func Close() {

}
