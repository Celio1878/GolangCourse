package greeting

import (
	"os"
	"testing"
)

const filename string = "local_greetings.txt"

func TestSavingFileLocally(t *testing.T) {
	// Arrange
	greets := greetings{
		"Hello %v",
		"Greeting to see you, %v",
		"Hail, %v! Well met !!!",
	}

	// Act
	err := greets.savingLocalFile(filename)

	// Assert
	if err != nil {
		t.Errorf("File can't be saved %v", err)
		return
	}

	t.Log("File saved successfully")
}

func TestReadingFile(t *testing.T) {
	// Arrange

	// Act
	res, err := greetingsFromFile(filename)

	// Assert
	if err != nil {
		t.Errorf("Error reading file %v", err)
		return
	}

	if len(res) < 1 {
		t.Errorf("Expected 3 greetings, got %v", len(res))
		return
	}

	t.Log("File read successfully")
	t.Log(res)
}

func TestRemovingLocalFile(t *testing.T) {
	err := os.Remove(filename)

	if err != nil {
		t.Errorf("Error removing file %v", err)
		return
	}

	t.Log("File removed successfully")
}
