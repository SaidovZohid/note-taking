package utils_test

import (
	"log"
	"testing"

	"github.com/SaidovZohid/note-taking/pkg/utils"
)

func TestPassword(t *testing.T) {
	TestCases := []string{"123456789", "987654321", "1234", "4321", "5445", "4545", "zohid2004"}
	for _, password := range TestCases {
		hashedPassword, err := utils.HashPassword(password)
		if err != nil {
			log.Printf("Test Failed: %v", password)
			continue
		}
		err = utils.CheckPassword(password, hashedPassword)
		if err != nil {
			log.Printf("Test Failed: %v", password)
			continue
		}
		log.Print("Test Passed")
	}
}
