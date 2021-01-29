package main

import (
	"fmt"
	"net/http"

	"github.com/nicobianchetti/Go-CleanArchitecture/cache"
	"github.com/nicobianchetti/Go-CleanArchitecture/controller"
	"github.com/nicobianchetti/Go-CleanArchitecture/repository"
	"github.com/nicobianchetti/Go-CleanArchitecture/router"
	"github.com/nicobianchetti/Go-CleanArchitecture/service"
)

var (
	permisoRepository repository.IPermisoRepository = repository.NewPermisoRepository()
	permisoService    service.IPermisoService       = service.NewPermisoService(permisoRepository)
	permisoCache      cache.PermisoCache            = cache.NewRedisCache("localhost:6379", 1, 200)
	permisoController controller.IPermisoController = controller.NewPermisoController(permisoService, permisoCache)
	httpRouter        router.IRouter                = router.NewMuxRouter()
)

func main() {
	const port string = ":8080"

	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Up and running...")
	})

	httpRouter.GET("/v1/permiso/permisos", permisoController.GetAll)
	httpRouter.GET("/v1/permiso/permiso/{id}", permisoController.GetByID)
	httpRouter.POST("/v1/permiso/permiso", permisoController.Create)

	httpRouter.SERVE(port)
}
