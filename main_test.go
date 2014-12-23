package configo

import "testing"

type Config struct {
	Title     string
	Port      string
	Staticdir string
}

func TestReadConfig(t *testing.T) {

	var cnf Config

	if err := ReadConfig("resources/test_config.toml", &cnf); err != nil {
		t.Error("Expected resources/test_config.toml, err: %g", err)
	}

	if cnf.Port != ":8123" {
		t.Error("Expected Port to be a string :8123, got ", cnf.Port)
	}
	if cnf.Staticdir != "assets/" {
		t.Error("Expected StaticDir to be a string assets/, got ", cnf.Staticdir)
	}
	if cnf.Title != "Test Configo" {
		t.Error("Expected Title to be a string 'Test Configo', got ", cnf.Title)
	}

}
