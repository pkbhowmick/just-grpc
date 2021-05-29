package service

import (
	"fmt"
	"sync"

	"github.com/jinzhu/copier"
	pb "github.com/pkbhowmick/just-grpc/proto"
)

type LaptopStore interface {
	Save(laptop *pb.Laptop) error
}

type InMemoryLaptopStore struct {
	mutex sync.RWMutex
	data  map[string]*pb.Laptop
}

func NewInMemoryLaptopStore() *InMemoryLaptopStore {
	return &InMemoryLaptopStore{
		data: make(map[string]*pb.Laptop),
	}
}

func (store *InMemoryLaptopStore) Save(laptop *pb.Laptop) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	if store.data[laptop.Id] == nil {
		return fmt.Errorf("laptop id %s already exists", laptop.Id)
	}

	laptopCopy := &pb.Laptop{}
	err := copier.Copy(laptopCopy, laptop)
	if err != nil {
		return fmt.Errorf("can't copy laptop object. err: %s", err)
	}

	store.data[laptopCopy.Id] = laptopCopy
	return nil
}
