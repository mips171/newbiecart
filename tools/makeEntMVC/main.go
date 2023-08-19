package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	// Define command-line arguments
	entityName := flag.String("name", "", "Name of the entity to create")
	flag.Parse()

	if *entityName == "" {
		fmt.Println("Please provide an entity name using the -name flag.")
		os.Exit(1)
	}

	// Create entity schema
	createEntitySchema(*entityName)

	// Create route
	createEntityRoute(*entityName)

	// Create view template
	createEntityView(*entityName)

	// Any other steps...

	fmt.Println("Entity and associated files successfully created!")
}

func createEntitySchema(name string) error {
	// Get the current directory
	currentDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("Failed to get the current directory: %v", err)
	}

	// Move up two directories to get the root directory
	rootDir := filepath.Join(currentDir, "../../")

	// Running: make ent-new $entityName
	cmd := exec.Command("make", "ent-new", "name="+name)
	cmd.Dir = rootDir
	if output, err := cmd.CombinedOutput(); err != nil {
		fmt.Printf("Failed to run 'make ent-new %s'. Output: %s\n", name, output)
		return err
	}

	// Running: make ent-gen
	cmd = exec.Command("make", "ent-gen")
	cmd.Dir = rootDir
	if output, err := cmd.CombinedOutput(); err != nil {
		fmt.Printf("Failed to run 'make ent-gen'. Output: %s\n", output)
		return err
	}

	return nil
}

func createEntityRoute(name string) {
	// Code to generate route based on templates...
}

func createEntityView(name string) {
	// Code to generate view template based on templates...
}
