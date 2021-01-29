package cache

import "github.com/nicobianchetti/Go-CleanArchitecture/model"

//PermisoCache .
type PermisoCache interface {
	Set(key string, value *model.Permiso)
	Get(key string) *model.Permiso
}
