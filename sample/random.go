package sample

import (
	"github.com/google/uuid"
	"math/rand"
)

func randomId() string {
	id := uuid.New()
	return id.String()
}

func randomStringFromSet(str ...string) string {
	size := len(str)
	if size == 0 {
		return ""
	}
	id := rand.Intn(size)
	return str[id]
}

func randomLaptopBrand() string {
	return randomStringFromSet("Apple", "HP", "Dell")
}

func randomLaptopName(brand string) string {
	switch brand {
	case "Apple":
		return randomStringFromSet("Macbook air", "Macbook pro", "Macbook ultra")
	case "HP":
		return randomStringFromSet("Probook", "Pavilion", "Notebook")
	default:
		return randomStringFromSet("Latitude", "Ultrabook")
	}
}

func randomCPUName(brand string) string {
	switch brand {
	case "Intel":
		return randomStringFromSet("Core i3 10100u", "Core i5 11400u", "Core i7 11700H")
	case "AMD":
		return randomStringFromSet("Ryzen 4001u", "Ryzen 5001H", "Ryzen 3500H")
	default:
		return "unknown"
	}
}

func randomInt(min int, max int) int {
	return min + rand.Intn(max - min)
}

func randomFloat(min float64, max float64) float64 {
	return max - min * rand.Float64()
}

func randomGPUName(brand string) string {
	switch brand {
	case "NVIDIA":
		return randomStringFromSet("GTX 1650 SUPER", "RTX 3050 TI", "GT 1030")
	case "AMD":
		return randomStringFromSet("RX 570", "RX 1050", "RX 2030")
	default:
		return "unknown"
	}
}

func randomBool() bool {
	n := rand.Intn(2)
	return n == 1
}



