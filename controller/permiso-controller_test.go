package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nicobianchetti/Go-CleanArchitecture/model"
	"github.com/nicobianchetti/Go-CleanArchitecture/repository"
	"github.com/nicobianchetti/Go-CleanArchitecture/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	NAME        string = "Permiso Post Test"
	DESCRIPTION string = "Creación de permiso desde test"
	OWNER       string = "Joe Satriani"
)

var (
	permisoRepositoryTst repository.IPermisoRepository = repository.NewPermisoRepository()
	permisoServiceTst    service.IPermisoService       = service.NewPermisoService(permisoRepositoryTst)
	permisoControllerTst IPermisoController            = NewPermisoController(permisoServiceTst)

	// testConteoller = NewPermisoController(nil)
)

type MockPermisoService struct {
	mock.Mock
}

func (m *MockPermisoService) Migrate() error {
	return nil
}

func (m *MockPermisoService) Create(pr *model.Permiso) (*model.Permiso, error) {
	//Vamos a hacer stub para devolver los argumentos
	args := m.Called()
	result := args.Get(0) //Devolver el primer argumento que recibe --> pr *model.Permiso

	return result.(*model.Permiso), args.Error(1)
}

func (m *MockPermisoService) GetAll() ([]model.Permiso, error) {

	//Vamos a hacer stub para devolver los argumentos
	args := m.Called()
	result := args.Get(0) //Devolver el primer argumento que recibe --> pr *model.Permiso

	return result.([]model.Permiso), args.Error(1)

}

func (m *MockPermisoService) GetByID(ID string) (*model.Permiso, error) {
	return nil, nil
}

func (m *MockPermisoService) Update(id string, p *model.Permiso) error {
	return nil
}

func (m *MockPermisoService) Delete(id string) error {
	return nil
}

func (m *MockPermisoService) Validate(permiso *model.Permiso) error {
	return nil
}

func TestGetAllPermiso(t *testing.T) {
	mockService := new(MockPermisoService)

	permiso := model.Permiso{ID: "123", Name: "Permiso 1", Description: "Descripcion de permiso 1", Status: true, Owner: "Nicolas"}

	mockService.On("GetAll").Return([]model.Permiso{permiso}, nil)

	testControllerM := NewPermisoController(mockService)

	//------------------------------------------------------------------------------------------

	//Create a GET HTTP Request
	req, err := http.NewRequest(http.MethodGet, "/v1/permiso/permisos", nil)

	if err != nil {
		t.Error("Eror in request test: ", err)
	}

	//Asing HTTP Handler function(controller GetAll function)
	handler := http.HandlerFunc(testControllerM.GetAll)

	//Record HTTP Response (with httptest library)
	response := httptest.NewRecorder()

	//Dispach the HTTP Request
	handler.ServeHTTP(response, req)

	//Add Assertions on the HTTP Status code ant the response
	status := response.Code

	if status != http.StatusOK {
		t.Errorf("Handler returned a wrong status code: got %v want %v", status, http.StatusOK)
	}

	//Si el status correcto -> Assertions
	//Decode the HTTP Response
	var permisoRes *[]model.DTOPermisoResponse

	json.NewDecoder(io.Reader(response.Body)).Decode(&permisoRes)

	fmt.Println("Permiso devuelto debug ", permiso)

	//Assert HTTP response
	assert.NotNil(t, (*permisoRes)[0].ID)
	assert.Equal(t, "Permiso 1", (*permisoRes)[0].Name)
	assert.Equal(t, "Descripcion de permiso 1", (*permisoRes)[0].Description)
	assert.Equal(t, "Nicolas", (*permisoRes)[0].Owner)
}

func TestCreate(t *testing.T) {
	//Create a new HTTP POST request -- Paso una matriz de bytes
	var jsonReq = []byte(`{"name": "Permiso Post Test","description": "Creación de permiso desde test","owner": "Joe Satriani"}`)
	req, err := http.NewRequest(http.MethodPost, "/v1/permiso/permiso", bytes.NewBuffer(jsonReq))

	if err != nil {
		t.Error("Eror in request test: ", err)
	}

	//Asing HTTP Handler function(controller Create function)
	handler := http.HandlerFunc(permisoControllerTst.Create)

	//Record HTTP Response (with httptest library)
	response := httptest.NewRecorder()

	//Dispach the HTTP Request
	handler.ServeHTTP(response, req)

	//Add Assertions on the HTTP Status code ant the response
	status := response.Code

	if status != http.StatusCreated {
		t.Errorf("Handler returned a wrong status code: got %v want %v", status, http.StatusCreated)
	}

	//Si el status correcto -> Assertions
	//Decode the HTTP Response
	var permiso model.DTOPermisoResponse
	// json.NewDecoder(response.Body).Decode(&permiso)
	json.NewDecoder(io.Reader(response.Body)).Decode(&permiso)

	fmt.Println("Permiso devuelto debug ", permiso)

	//Assert HTTP response
	assert.NotNil(t, permiso.ID)
	assert.Equal(t, NAME, permiso.Name)
	assert.Equal(t, DESCRIPTION, permiso.Description)
	assert.Equal(t, OWNER, permiso.Owner)

	err = cleanUp(&permiso)
	if err != nil {
		fmt.Println("Error al eliminar registro insertado para test ", err)
	}
}

/*
func TestGetAll(t *testing.T) {
	//Create a GET HTTP Request
	req, err := http.NewRequest(http.MethodGet, "/v1/permiso/permisos", nil)

	if err != nil {
		t.Error("Eror in request test: ", err)
	}

	//Asing HTTP Handler function(controller GetAll function)
	handler := http.HandlerFunc(permisoControllerTst.GetAll)

	//Record HTTP Response (with httptest library)
	response := httptest.NewRecorder()

	//Dispach the HTTP Request
	handler.ServeHTTP(response, req)

	//Add Assertions on the HTTP Status code ant the response
	status := response.Code

	if status != http.StatusCreated {
		t.Errorf("Handler returned a wrong status code: got %v want %v", status, http.StatusOK)
	}

	//Si el status correcto -> Assertions
	//Decode the HTTP Response
	var permiso *[]model.DTOPermisoResponse

	json.NewDecoder(io.Reader(response.Body)).Decode(&permiso)

	fmt.Println("Permiso devuelto debug ", permiso)

	//Assert HTTP response
	assert.NotNil(t, (*permiso)[0].ID)
	assert.Equal(t, NAME, (*permiso)[0].Name)
	assert.Equal(t, DESCRIPTION, (*permiso)[0].Description)
	assert.Equal(t, OWNER, (*permiso)[0].Owner)
}
*/

func cleanUp(permiso *model.DTOPermisoResponse) error {
	//Eliminar registro creado
	id := permiso.ID
	err := permisoRepositoryTst.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
