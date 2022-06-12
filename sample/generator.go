package sample

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/namesjc/pcbook/pb"
)

// NewKeyboard returns a new sample keyboard
func NewKeyboard() *pb.Keyboard {
	keyboard := &pb.Keyboard{
		Layout:  randomKeyboardLayout(),
		Backlit: randomBool(),
	}
	return keyboard
}

// NewCPU returns a new simple CPU
func NewCPU() *pb.CPU {
	brand := randomCPUBrand()
	name := randomCPUName(brand)

	numberOfCores := randomInt(2, 8)
	numberOfThreds := randomInt(numberOfCores, 12)

	minGhz := randomFloat64(2.0, 3.5)
	maxGhz := randomFloat64(minGhz, 5.0)

	cpu := &pb.CPU{
		Brand:         brand,
		Name:          name,
		NumberCores:   uint32(numberOfCores),
		NumberThreads: uint32(numberOfThreds),
		MinGhz:        minGhz,
		MaxGhz:        maxGhz,
	}

	return cpu
}

// NewGPU returns a new simple GPU
func NewGPU() *pb.GPU {
	brand := randomGPUBrand()
	name := randomGPUName(brand)

	minGhz := randomFloat64(1.0, 1.5)
	maxGhz := randomFloat64(minGhz, 2.0)

	memory := &pb.Memory{
		Value: uint64(randomInt(2, 6)),
		Uint:  pb.Memory_GIGABYTE,
	}

	gpu := &pb.GPU{
		Brand:  brand,
		Name:   name,
		MinGhz: minGhz,
		MaxGhz: maxGhz,
		Memory: memory,
	}

	return gpu
}

// NewRAM returns a new simple RAM
func NewRAM() *pb.Memory {
	ram := &pb.Memory{
		Value: uint64(randomInt(2, 64)),
		Uint:  pb.Memory_GIGABYTE,
	}

	return ram
}

// NewSSD returns a new simple SSD
func NEWSSD() *pb.Storage {
	ssd := &pb.Storage{
		Driver: pb.Storage_SSD,
		Memory: &pb.Memory{
			Value: uint64(randomInt(128, 1024)),
			Uint:  pb.Memory_GIGABYTE,
		},
	}

	return ssd
}

// NewHDD returns a new simple HDD
func NEWHDD() *pb.Storage {
	hdd := &pb.Storage{
		Driver: pb.Storage_HDD,
		Memory: &pb.Memory{
			Value: uint64(randomInt(1, 6)),
			Uint:  pb.Memory_TERABYTE,
		},
	}

	return hdd
}

// NewScreen returns a new simple Screen
func NewScreen() *pb.Screen {
	screen := &pb.Screen{
		SizeInch:   randomFloat32(13, 17),
		Resolution: randomScreenResolution(),
		Panel:      randomScreenPanel(),
		Multitouch: randomBool(),
	}

	return screen
}

// NewLaptop returns a new simple Laptop
func NewLaptop() *pb.Laptop {

	brand := randomLaptopBrand()
	name := randomLaptopName(brand)

	laptop := &pb.Laptop{
		Id:    randomID(),
		Brand: brand,
		Name:  name,
		Cpu:   NewCPU(),
		Ram:   NewRAM(),
		Gpus: []*pb.GPU{
			NewGPU(),
		},
		Storages: []*pb.Storage{
			NEWSSD(),
			NEWHDD(),
		},
		Screen:   NewScreen(),
		Keyboard: NewKeyboard(),
		Weight: &pb.Laptop_WeightKg{
			WeightKg: randomFloat64(1.0, 3.0),
		},
		PriceUsd:    randomFloat64(1500, 3000),
		ReleaseYear: uint32(randomInt(2015, 2020)),
		UpdatedAt:   ptypes.TimestampNow(),
	}

	return laptop
}
