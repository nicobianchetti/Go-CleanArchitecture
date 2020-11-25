package service

import "testing"

//ejemplo del video :https://www.youtube.com/watch?v=uB_45bSIyik&list=PL3eAkoh7fypqUQUQPn-bXtfiYT_ZSVKmB&index=4

func TestValidateEmptyPost(t *testing.T) {
	testService := NewPermisoService(nil)

	err := testService.Validate(nil)
	_ = err
}
