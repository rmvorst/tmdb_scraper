package env

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type EnvConfig struct {
	ApiKey  string
	RootNFO string
}

func SetupEnv() error {
	envPath, err := GetEnvPath()
	if err != nil {
		return fmt.Errorf("Error: Cannot get env path")
	}

	err = os.MkdirAll(envPath, 0755)
	if err != nil {
		return fmt.Errorf("Error: Cannot make tmdb_scraper dir in ~/.config")
	}

	err = setupEnvFile()
	if err != nil {
		return err
	}

	return nil
}

func setupEnvFile() error {
	fpath, err := GetEnvPath()

	if err != nil {
		return fmt.Errorf("Error: Cannot get env path")
	}

	envFPath := filepath.Join(fpath, ".env")
	fmt.Println(envFPath)
	file, _ := os.OpenFile(envFPath, os.O_RDWR|os.O_CREATE, 0755)

	defer file.Close()

	env, _ := godotenv.Read(envFPath)
	env = checkEnvKeys(env)

	err = godotenv.Write(env, envFPath)
	if err != nil {
		return fmt.Errorf("Error: Updating .env file")
	}

	return nil
}

func checkEnvKeys(env map[string]string) map[string]string {
	value, ok := env["API_KEY"]
	if !ok || value == "" {
		fmt.Print("Provide your tmdb-api key>>> ")
		fmt.Scanln(&value)
		env["API_KEY"] = value
	}

	value, ok = env["NFO_ROOT"]
	if !ok || value == "" {
		fmt.Print("Provide directory to save NFO files>>> ")
		fmt.Scanln(&value)
		env["NFO_ROOT"] = value
	}

	return env
}

func GetEnvPath() (string, error) {
	// Get user's home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("Error: Cannot get users home directory")
	}

	targetDir := ".config/tmdb_scraper"

	// The .env file will be at ~/.config/tmdb_scraper
	fullPath := filepath.Join(homeDir, targetDir)

	return fullPath, nil
}
