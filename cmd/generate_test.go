package cmd

import (
	"bufio"

	"os"
	"strings"
	"testing"
)

func TestGetSecret(t *testing.T) {
	envFile := "test.env"
	f, err := os.Create(envFile)
	if err != nil {
		t.Fatalf("Could not create file: %s", err)
	}
	defer os.Remove(envFile)
	defer f.Close()

	_, err = f.WriteString("KEY1=Value1\nKEY2=Value2\n")
	if err != nil {
		t.Fatalf("Could not write to file: %s", err)
	}

	secret := getSecret(envFile)
	expected := map[string]string{
		"KEY1": "Value1",
		"KEY2": "Value2",
	}

	if len(secret) != len(expected) {
		t.Fatalf("Incorrect number of environment variables, expected %d but got %d", len(expected), len(secret))
	}

	for k, v := range expected {
		value, exists := secret[k]
		if !exists {
			t.Fatalf("Expected environment variable %s not found", k)
		}
		if value != v {
			t.Fatalf("Incorrect value for environment variable %s, expected %s but got %s", k, v, value)
		}
	}
}

func TestWriteYaml(t *testing.T) {
	secret := map[string]string{
		"KEY1": "Value1",
		"KEY2": "Value2",
	}
	writeYaml(secret)

	yaml, err := os.ReadFile("secret.yaml")
	if err != nil {
		t.Fatalf("Could not read file: %s", err)
	}

	scanner := bufio.NewScanner(strings.NewReader(string(yaml)))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "  KEY1:") {
			if line != "  KEY1: VmFsdWUx" {
				t.Fatalf("Incorrect value for environment variable KEY1, expected 'Value1' but got '%s'", strings.Split(line, ": ")[1])
			}
		} else if strings.HasPrefix(line, "  KEY2:") {
			if line != "  KEY2: VmFsdWUy" {
				t.Fatalf("Incorrect value for environment variable KEY2, expected 'Value2' but got '%s'", strings.Split(line, ": ")[1])
			}
		}
	}
}
