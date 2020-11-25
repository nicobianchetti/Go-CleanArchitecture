package controller

import (
	"encoding/json"
	"net/http"

	"github.com/nicobianchetti/Go-CleanArchitecture/model"
	"github.com/nicobianchetti/Go-CleanArchitecture/service"
)

//IPermisoController interac with IPermisoService
type IPermisoController interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	GetByID(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type permisoController struct{}

var (
	permisoService service.IPermisoService
)

//NewPermisoController create new instance of controller
func NewPermisoController(service service.IPermisoService) IPermisoController {
	permisoService = service
	return &permisoController{}
}

func (c *permisoController) GetAll(w http.ResponseWriter, r *http.Request) {
	// pr, err := c.controller.GetAll()

	// if err != nil {
	// 	responsePermisos(w, http.StatusNotFound, nil)
	// }

	// var dtoPermiso []*DTOPermisoResponse

	// for _, permiso := range *pr {
	// 	dtoItem := NewPermisoDTOWFromPermiso(&permiso)
	// 	dtoPermiso = append(dtoPermiso, dtoItem)
	// }

	// responsePermisos(w, http.StatusOK, dtoPermiso)

	// pr, err := c.controller.GetAll()

	// if err != nil {
	// 	responsePermisos(w, http.StatusNotFound, nil)
	// }

	// var dtoPermiso []*DTOPermisoResponse

	// for _, permiso := range *pr {
	// 	dtoItem := NewPermisoDTOWFromPermiso(&permiso)
	// 	dtoPermiso = append(dtoPermiso, dtoItem)
	// }

	// responsePermisos(w, http.StatusOK, dtoPermiso)

	pr, err := permisoService.GetAll()

	if err != nil {
		responsePermisos(w, http.StatusNotFound, nil)
	}

	var dtoPermiso []*model.DTOPermisoResponse

	for _, permiso := range *pr {
		dtoItem := model.NewPermisoDTOWFromPermiso(&permiso)
		dtoPermiso = append(dtoPermiso, dtoItem)
	}

	responsePermisos(w, http.StatusOK, dtoPermiso)
}

func (c *permisoController) GetByID(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)

	// permiso, err := c.controller.GetByID(vars["id"])

	// if err != nil {
	// 	// responsePermiso(w, http.StatusNotFound, nil)
	// 	http.Error(w, "Permiso Not found", http.StatusNotFound)
	// 	return
	// }

	// var dtoPermiso *DTOPermisoResponse

	// dtoPermiso = NewPermisoDTOWFromPermiso(permiso)

	// responsePermiso(w, http.StatusOK, dtoPermiso)

}

func (c *permisoController) Create(w http.ResponseWriter, r *http.Request) {
	// decoder := json.NewDecoder(r.Body)

	// var permisoDTO DTOPermisoRequest
	// err := decoder.Decode(&permisoDTO)

	// defer r.Body.Close()

	// if err != nil {
	// 	responsePermiso(w, http.StatusBadRequest, nil)
	// 	return
	// }

	// permiso := NewPermisoFromDTO(&permisoDTO)

	// err = c.controller.Create(permiso)

	// if err != nil {
	// 	responsePermiso(w, http.StatusBadRequest, nil)
	// 	return
	// }

	// responsePermiso(w, http.StatusCreated, nil)

}

func (c *permisoController) Update(w http.ResponseWriter, r *http.Request) {
	// fmt.Print("\n Entra al update")

	// vars := mux.Vars(r)
	// id := vars["id"]

	// fmt.Print("\n ID ", id)

	// decoder := json.NewDecoder(r.Body)

	// var permisoDTO DTOPermisoRequest
	// err := decoder.Decode(&permisoDTO)

	// fmt.Print("\n Permiso DTO ", permisoDTO)

	// defer r.Body.Close()

	// if err != nil {
	// 	responsePermiso(w, http.StatusInternalServerError, nil)
	// 	return
	// }

	// permiso := NewPermisoFromDTO(&permisoDTO)

	// spew.Dump(permiso)

	// err = c.controller.Update(id, permiso)

	// if err != nil {
	// 	responsePermiso(w, http.StatusInternalServerError, nil)
	// 	fmt.Print(err)
	// 	return
	// }

	// responsePermiso(w, http.StatusOK, nil)
}

func (c *permisoController) Delete(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)

	// err := c.controller.Delete(vars["id"])

	// if err != nil {
	// 	responsePermiso(w, http.StatusNotFound, nil)
	// 	return
	// }

	// responsePermiso(w, http.StatusNoContent, nil)

}

// func responsePermiso(w http.ResponseWriter, status int, permiso *DTOPermisoResponse) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(status)
// 	json.NewEncoder(w).Encode(permiso)
// }

func responsePermisos(w http.ResponseWriter, status int, permisos []*model.DTOPermisoResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(permisos)
}
