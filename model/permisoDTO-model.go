package model

import (
	"time"
)

// //PermisoDTOR is struc of request permiso
// type PermisoDTOR struct {
// 	pagination int
// 	cantreg    int
// 	mode       string
// }

// DTOPermisoResponse contiene datos de Permiso para mostrar en api
type DTOPermisoResponse struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Status      bool      `json:"status"`
	Owner       string    `json:"owner"`
	CreatedAt   time.Time `json:"created"`
	UpdatedAt   time.Time `json:"updated"`
}

// DTOPermisoRequest trae datos para crear nuevo Permiso en base de datos
type DTOPermisoRequest struct {
	// ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
	Owner       string `json:"owner"`
}

//NewPermisoDTOWFromPermiso created dto from entity
func NewPermisoDTOWFromPermiso(p *Permiso) *DTOPermisoResponse {
	return &DTOPermisoResponse{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Status:      p.Status,
		Owner:       p.Owner,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}
}

//NewPermisoFromDTO created struc Permiso from dto input
func NewPermisoFromDTO(d *DTOPermisoRequest) *Permiso {
	return &Permiso{
		// ID:          d.ID,
		Name:        d.Name,
		Description: d.Description,
		Status:      d.Status,
		Owner:       d.Owner,
	}
}
