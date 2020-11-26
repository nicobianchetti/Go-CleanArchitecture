package service

import (
	"errors"

	"github.com/nicobianchetti/Go-CleanArchitecture/model"
	"github.com/nicobianchetti/Go-CleanArchitecture/repository"
)

//IPermisoService interact with IPermisoRepository
type IPermisoService interface {
	Migrate() error
	Create(*model.Permiso) (*model.Permiso, error)
	Update(string, *model.Permiso) error
	GetAll() ([]model.Permiso, error)
	GetByID(string) (*model.Permiso, error)
	Delete(string) error
	Validate(permiso *model.Permiso) error
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
func (s *permisoService) Validate(permiso *model.Permiso) error {
	if permiso == nil {
		err := errors.New("Permiso is empty")
		return err
	}

	if permiso.Name == "" {
		err := errors.New("The name permiso es empty")
		return err
	}
	return nil
}

// Migrate is used for migrate permiso
func (s *permisoService) Migrate() error {
	return nil
	// return s.service.Migrate()
}

// Create is used for create a permiso
func (s *permisoService) Create(p *model.Permiso) (*model.Permiso, error) {
	// p.ID = uuid.New().String()
	// p.Status = true
	// p.CreatedAt = time.Now()
	// p.UpdatedAt = time.Now()

	// return s.service.Create(p)

	return nil, nil
}

// GetAll is used for get all the permisos
func (s *permisoService) GetAll() ([]model.Permiso, error) {
	// return s.service.GetAll()

	// return repo.GetAll()

	permisos, _ := repo.GetAll()

	for _, v := range permisos {
		v.ID = "55"
	}

	return permisos, nil

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
