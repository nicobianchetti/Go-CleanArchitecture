package service

import (
	"testing"

	"github.com/nicobianchetti/Go-CleanArchitecture/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

//ejemplo del video :https://www.youtube.com/watch?v=uB_45bSIyik&list=PL3eAkoh7fypqUQUQPn-bXtfiYT_ZSVKmB&index=4

var (
	testService = NewPermisoService(nil)
)

type MockPermisoRepository struct {
	mock.Mock
}

func (m *MockPermisoRepository) Migrate() error {
	return nil
}

func (m *MockPermisoRepository) Create(pr *model.Permiso) (*model.Permiso, error) {
	//Vamos a hacer stub para devolver los argumentos
	args := m.Called()
	result := args.Get(0) //Devolver el primer argumento que recibe --> pr *model.Permiso

	return result.(*model.Permiso), args.Error(1)
}

func (m *MockPermisoRepository) GetAll() ([]model.Permiso, error) {

	//Vamos a hacer stub para devolver los argumentos
	args := m.Called()
	result := args.Get(0) //Devolver el primer argumento que recibe --> pr *model.Permiso

	return result.([]model.Permiso), args.Error(1)

}

func (m *MockPermisoRepository) GetByID(ID string) (*model.Permiso, error) {
	return nil, nil
}

func (m *MockPermisoRepository) Update(id string, p *model.Permiso) error {
	return nil
}

func (m *MockPermisoRepository) Delete(id string) error {
	return nil
}

func TestGetAllPermiso(t *testing.T) {
	mockRepo := new(MockPermisoRepository)

	permiso := model.Permiso{ID: "123", Name: "Permiso 1", Description: "Descripcion de permiso 1", Status: true, Owner: "Nicolas"}

	//Setup expected
	//Cuando el method GetAll se invoca en éste mock , va a devolver un array incluyendo el elemento que se le pasa(permiso)
	mockRepo.On("GetAll").Return([]model.Permiso{permiso}, nil)

	testServiceM := NewPermisoService(mockRepo)

	result, _ := testServiceM.GetAll()

	//Mock Assertion: Behavioral
	mockRepo.AssertExpectations(t)

	//Data assertion
	assert.Equal(t, "123", result[0].ID)
	assert.Equal(t, "Permiso 1", result[0].Name)
	assert.Equal(t, "Descripcion de permiso 1", result[0].Description)
	assert.Equal(t, true, result[0].Status)
	assert.Equal(t, "Nicolas", result[0].Owner)

}

func TestCreatePermiso(t *testing.T) {
	mockRepo := new(MockPermisoRepository)

	permiso := model.Permiso{Name: "Permiso 1", Description: "Descripcion de permiso 1", Status: true, Owner: "Nicolas"}

	//Setup expectations , uso método Create DEL REPOSITORY
	mockRepo.On("Create").Return(&permiso, nil)

	//Creo el servicio de prueba que es el servicio que estoy probando y le paso el repositorio simulado
	testServiceM := NewPermisoService(mockRepo)

	//Lamo al create implementado en el servicio
	result, err := testServiceM.Create(&permiso)

	//Mock Assertion: Behavioral, le paso la prueba (t)
	mockRepo.AssertExpectations(t)

	//Data assertion
	assert.NotNil(t, result.ID)
	assert.Equal(t, "Permiso 1", result.Name)
	assert.Equal(t, "Descripcion de permiso 1", result.Description)
	assert.Equal(t, true, result.Status)
	assert.Equal(t, "Nicolas", result.Owner)

	assert.Nil(t, err)

}

func TestValidateEmptyPermiso(t *testing.T) {
	// testService := NewPermisoService(nil)

	err := testService.Validate(nil)

	assert.NotNil(t, err)

	assert.Equal(t, "Permiso is empty", err.Error())
}

func TestValidateEmptyPermisoTitle(t *testing.T) {
	// testService := NewPermisoService(nil)

	permiso := model.Permiso{ID: "123", Name: ""}

	err := testService.Validate(&permiso)
	assert.NotNil(t, err)

	assert.Equal(t, "The name permiso es empty", err.Error())
}
