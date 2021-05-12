package sample

import (
	pb "github.com/pkbhowmick/just-grpc/proto"

	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func NewLaptop() *pb.Laptop {
	brand := randomLaptopBrand()
	name := randomLaptopName(brand)
	laptop := &pb.Laptop{
		Id:           randomId(),
		Brand:        brand,
		Name :        name,
		Cpu:          NewCPU(),
		Gpu:          []*pb.GPU{NewGPU()},
		Memory:       NewMemory(),
		Storage:      []*pb.Storage{NewStorageHDD(), NewStorageSSD()},
		Screen:       NewScreen(),
		Keyboard:     NewKeyboard(),
		Weight:       &pb.Laptop_WeightKg{
			WeightKg: randomFloat(1.0, 2.5),
		},
		PriceBdt:     randomFloat(50000, 120000),
		ReleasedYear: uint32(randomInt(2015, 2021)),
		UpdatedAt: timestamppb.Now(),
	}
	return laptop
}

func NewCPU() *pb.CPU {
	brand := randomStringFromSet("Intel", "AMD")
	name := randomCPUName(brand)
	cores := randomInt(2,8)
	threads := randomInt(cores, 16)
	minGhz := randomFloat(2.0,3.5)
	maxGhz := randomFloat(minGhz, 5.0)
	cpu := &pb.CPU{
		Brand:         brand,
		Name:          name,
		NumberCores:   uint32(cores),
		NumberThreads: uint32(threads),
		MinGhz:        minGhz,
		MaxGhz:        maxGhz,
	}
	return cpu
}

func NewGPU() *pb.GPU {
	brand := randomStringFromSet("NVIDIA", "AMD")
	name := randomGPUName(brand)
	minGhz := randomFloat(1.0, 1.5)
	maxGhz := randomFloat(minGhz, 2.0)
	gpu := &pb.GPU{
		Brand:  brand,
		Name:   name,
		MinGhz: minGhz,
		MaxGhz: maxGhz,
		Memory: NewMemory(),
	}
	return gpu
}

func NewMemory() *pb.Memory {
	memory := &pb.Memory{
		Value: uint64(randomInt(4, 32)),
		Unit:  pb.Memory_UNIT_GIGABYTE,
	}
	return memory
}

func NewStorageHDD() *pb.Storage {
	storage := &pb.Storage{
		Driver: pb.Storage_DRIVER_HDD,
		Memory: &pb.Memory{
			Value: uint64(randomInt(1, 5)),
			Unit: pb.Memory_UNIT_TERABYTE,
		},
	}
	return storage
}

func NewStorageSSD() *pb.Storage {
	storage := &pb.Storage{
		Driver: pb.Storage_DRIVER_SSD,
		Memory: &pb.Memory{
			Value: uint64(randomInt(120, 1024)),
			Unit:  pb.Memory_UNIT_GIGABYTE,
		},
	}
	return storage
}

func NewScreen() *pb.Screen {
	w :=  uint32(randomInt(10, 25))
	h := uint32(w * 16 / 9)
	sc := &pb.Screen{
		SizeInch:   float32(randomFloat(21.5, 34.0)),
		Resolution: &pb.Screen_Resolution{
			Width: w,
			Height: h,
		},
		Panel:      pb.Screen_PANEL_IPS,
		MultiTouch: randomBool(),
	}
	return sc
}

func NewKeyboard() *pb.Keyboard {
	kb := &pb.Keyboard{
		Layout:  pb.Keyboard_LAYOUT_QWERTY,
		Backlit: randomBool(),
	}
	return kb
}