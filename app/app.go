package app

import (
	"admin-panel/v1/pkg/config"
	"admin-panel/v1/pkg/workers"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

type App struct {
	Engine *gin.Engine
}

type middleware func(*gin.Engine)
type route func(*gin.Engine)

func New() *App {
	version := "v1.0.0"
	banner := `
   db       888b.    8b   d8    888    8b  8          888b.       db       8b  8    8888    8    
  dPYb      8   8    8YbmdP8     8     8Ybm8          8  .8      dPYb      8Ybm8    8www    8    
 dPwwYb     8   8    8  "  8     8     8  "8          8wwP'     dPwwYb     8  "8    8       8    
dP    Yb    888P'    8     8    888    8   8          8        dP    Yb    8   8    8888    8888 
                                                                                                 
                                       db       888b.    888                                     
                                      dPYb      8  .8     8                                      
                                     dPwwYb     8wwP'     8                                      
                                    dP    Yb    8        888    - Core APIs %s                                 
                                                                                                 ` + "\n"
	fmt.Printf(
		banner,
		version,
	)
	return &App{
		Engine: gin.New(),
	}
}

func (app *App) Middleware(middlewares ...middleware) *App {
	for _, middleware := range middlewares {
		middleware(app.Engine)
	}
	return app
}

func (app *App) Route(routes ...route) *App {
	for _, route := range routes {
		route(app.Engine)
	}
	return app
}

func (app *App) Run(port string) {
	err := app.Engine.Run(port)
	if err != nil {
		log.Fatal(err)
	}
}

func (app *App) Worker(wcf *config.WorkerConfig) {
	workers.WorkerConfig = wcf
}
