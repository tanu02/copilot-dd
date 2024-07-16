//write a function to seed a random number generator

import (
	"math/rand"
	"time"
)

func seedRandomNumberGenerator() {
	rand.Seed(time.Now().UnixNano())
}

// Create and seed the generator.
	// Typically a non-fixed seed should be used, such as time.Now().UnixNano().
	// Using a fixed seed will produce the same output on every run.
	// r := rand.New(rand.NewSource(99))
	// r := rand.New(rand.NewSource(time.Now().UnixNano()))
	
	func main() {
		seedRandomNumberGenerator()
		// Generate random numbers using the rand package
		// ...
	}
	func generateRandomNumbers() {
		for i := 0; i < 10; i++ {
			fmt.Println(rand.Intn(100))
		}
	}