package main

import (
	"fmt"
	"net/http"

	"github.com/nicobianchetti/Go-CleanArchitecture/controller"
	"github.com/nicobianchetti/Go-CleanArchitecture/router"
)

var (
	permisoController controller.IPermisoController = controller.NewPermisoController()
	httpRouter        router.IRouter                = router.NewMuxRouter()
)

func main() {
	const port string = ":8080"

	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Up and running...")
	})

	httpRouter.GET("/permisos", permisoController.GetAll)

	httpRouter.SERVE(port)
}
