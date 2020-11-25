package service

import (
	"github.com/nicobianchetti/Go-CleanArchitecture/model"
	"github.com/nicobianchetti/Go-CleanArchitecture/repository"
)

//IPermisoService interact with IPermisoRepository
type IPermisoService interface {
	Migrate() error
	Create(*model.Permiso) error
	Update(string, *model.Permiso) error
	GetAll() (*model.Permisos, error)
	GetByID(string) (*model.Permiso, error)
	Delete(string) error
}

type permisoService struct{}

var (
	repo repository.IPermisoRepository
)

//NewPermisoService create new instance of service
func NewPermisoService(repos repository.IPermisoRepository) IPermisoService {
	repo = repos
	return &permisoService{}
}

// Migrate is used for migrate permiso
func (s *permisoService) Migrate() error {
	return nil
	// return s.service.Migrate()
}

// Create is used for create a permiso
func (s *permisoService) Create(p *model.Permiso) error {
	// p.ID = uuid.New().String()
	// p.Status = true
	// p.CreatedAt = time.Now()
	// p.UpdatedAt = time.Now()

	// return s.service.Create(p)

	return nil
}

// GetAll is used for get all the permisos
func (s *permisoService) GetAll() (*model.Permisos, error) {
	// return s.service.GetAll()

	return repo.GetAll()

	// return nil, nil
}

// GetByID is used for get a permiso
func (s *permisoService) GetByID(id string) (*model.Permiso, error) {
	// return s.service.GetByID(id)
	return nil, nil
}

// Update is used for update a permiso
func (s *permisoService) Update(id string, p *model.Permiso) error {

	// p.UpdatedAt = time.Now()

	// return s.service.Update(id, p)

	return nil
}

// Delete is used for delete a permiso
func (s *permisoService) Delete(id string) error {
	// return s.service.Delete(id)

	return nil
}
