package repository

import (
	"testing"

	"../../models"
)

func TestPortsMap(t *testing.T) {
	db := NewPortsMap()
	data1, data2 := dummyData()
	key1, key2 := "TEST1", "TEST2"

	result, err := db.Add(key1, data1)
	if err != nil {
		t.Fatal(err)
	}
	if result != key1 {
		t.Errorf("DB returned unexpected key: got %s want %s",
			result, key1)
	}
	result, err = db.Add(key1, data2)
	if err == nil {
		t.Errorf("DB Added existent: %s", key1)
	}
	_, errGet := db.Get(key1)
	if errGet != nil {
		t.Fatal(errGet)
	}
	err = db.Update(key1, data2)
	if err != nil {
		t.Fatal(errGet)
	}
	err = db.Delete(key2)
	if err == nil {
		t.Errorf("DB deleted non-existent key: %s", key2)
	}
	result, err = db.Add(key2, data2)
	if err != nil {
		t.Fatal(err)
	}
	err = db.Update(key1, data1)
	if err != nil {
		t.Fatal(errGet)
	}
	p := models.Page{
		Start: 0,
		Size:  10,
	}
	ports, err := db.GetPage(p)
	if err != nil {
		t.Fatal(errGet)
	}
	if len(ports) != 2 {
		t.Errorf("DB listed invalid page size: %d", len(ports))
	}
}

func dummyData() (models.Data, models.Data) {
	data1 := models.Data{
		Name:        "Data 1",
		City:        "City 1",
		Country:     "Country 1",
		Coordinates: []float32{12.0, 32.9},
		Province:    "Prov 1",
		Timezone:    "Timezone 1",
		Code:        "Code 1",
	}
	data2 := models.Data{
		Name:        "Data 2",
		City:        "City 2",
		Country:     "Country 3",
		Coordinates: []float32{22.0, 67.34},
		Province:    "Prov 2",
		Timezone:    "Timezone 2",
		Code:        "Code 2",
	}
	return data1, data2
}
