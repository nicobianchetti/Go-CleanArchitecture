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
)

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

func TestGetAll(t *testing.T) {

}

func cleanUp(permiso *model.DTOPermisoResponse) error {
	//Eliminar registro creado
	id := permiso.ID
	err := permisoRepositoryTst.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
