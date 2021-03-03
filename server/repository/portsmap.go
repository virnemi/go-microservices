package repository

import (
	"errors"
	"fmt"
	"log"
	"sync"

	"../../models"
)

// InMemory database to persists data
type PortsMap struct {
	Ports     map[string]models.Data
	Positions []string
	Lock      sync.RWMutex
}

// Database constructor
func NewPortsMap() *PortsMap {
	db := &PortsMap{
		Ports:     make(map[string]models.Data),
		Positions: make([]string, 0),
		Lock:      sync.RWMutex{},
	}
	return db
}

// Returns a port from the database if the key exists in it.
// Returns an error otherwise
func (db *PortsMap) Get(key string) (*models.Port, error) {
	db.Lock.RLock()
	defer db.Lock.RUnlock()
	value, exists := db.Ports[key]
	if !exists {
		erMsg := fmt.Sprintf("Port not found with key %s", key)
		log.Println(erMsg)
		return nil, errors.New(erMsg)
	}
	port := &models.Port{key: value}
	return port, nil
}

// Returns a list of ports from the database filtering by the page information.
// If the start position is bigger than the persisted list's size, it returns an error
// The size of the returned list may be smaller than the requested page size,
// if the end of the list is reached within the start position requested.
func (db *PortsMap) GetPage(page models.Page) ([]models.Port, error) {
	db.Lock.RLock()
	defer db.Lock.RUnlock()
	if page.Start >= int32(len(db.Ports)) {
		return nil, errors.New("Invalid page information.")
	}
	p := make([]string, page.Size)
	copy(p, db.Positions[page.Start:])
	list := make([]models.Port, 0)
	for _, key := range p {
		data, exists := db.Ports[key]
		if exists {
			port := models.Port{key: data}
			list = append(list, port)
		}
	}
	return list, nil
}

// Adds a port with the provided key in the database if the key doesn't exists yet.
// Returns an error if the key already exists
func (db *PortsMap) Add(key string, m models.Data) (string, error) {
	db.Lock.Lock()
	defer db.Lock.Unlock()
	if _, exists := db.Ports[key]; exists {
		erMsg := fmt.Sprintf("Port with key '%s' already exists.", key)
		return "", errors.New(erMsg)
	}
	db.Ports[key] = m
	db.Positions = append(db.Positions, key)
	return key, nil
}

// Updates a port with the provided key in the database if the key exists.
// Returns an error if the key doesn't exists
func (db *PortsMap) Update(key string, m models.Data) error {
	db.Lock.Lock()
	defer db.Lock.Unlock()
	if _, exists := db.Ports[key]; !exists {
		erMsg := fmt.Sprintf("Port not found with key %s", key)
		log.Println(erMsg)
		return errors.New(erMsg)
	}
	db.Ports[key] = m
	return nil
}

// Deletes a port with the provided key from the database if the key exists.
// Returns an error if the key doesn't exists
func (db *PortsMap) Delete(key string) error {
	db.Lock.Lock()
	defer db.Lock.Unlock()
	if _, exists := db.Ports[key]; !exists {
		erMsg := fmt.Sprintf("Port not found with key %s", key)
		log.Println(erMsg)
		return errors.New(erMsg)
	}
	delete(db.Ports, key)
	for idx, value := range db.Positions {
		if value == key {
			db.Positions = append(db.Positions[:idx], db.Positions[idx+1:]...)
			break
		}
	}
	return nil
}
