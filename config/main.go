package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"gopkg.in/ini.v1"
)

// GetAWSCredentials reads the AWS credentials file and returns the access key id and secret access key for the specified profile.
func GetAWSCredentials() (accessKeyID, secretAccessKey string, err error) {
	profileName := "AstroMailApp"
	// Find the user's home directory.
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", "", fmt.Errorf("unable to find home directory: %w", err)
	}

	// Construct the path to the AWS credentials file.
	credsFilePath := filepath.Join(homeDir, ".aws", "credentials")

	// Load the credentials file.
	cfg, err := ini.Load(credsFilePath)
	if err != nil {
		return "", "", fmt.Errorf("unable to load AWS credentials file: %w", err)
	}

	// Ensure the specified profile exists.
	if !cfg.HasSection(profileName) {
		return "", "", fmt.Errorf("profile %s does not exist", profileName)
	}

	// Get the AWS Access Key ID and Secret Access Key from the profile.
	section, err := cfg.GetSection(profileName)
	if err != nil {
		return "", "", fmt.Errorf("unable to get section for profile %s: %w", profileName, err)
	}

	accessKeyID = section.Key("aws_access_key_id").String()
	secretAccessKey = section.Key("aws_secret_access_key").String()

	if accessKeyID == "" || secretAccessKey == "" {
		return "", "", fmt.Errorf("credentials for profile %s are incomplete or missing", profileName)
	}

	return accessKeyID, secretAccessKey, nil
}

// WriteKeyToFile writes a key to a JSON file, replacing the existing key if it already exists.
func WriteKeyToFile(key, value, filename string) error {
	// Check if the file exists
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		// If the file does not exist, create it
		return createNewFile(filename)
	}

	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Read existing content
	existingData, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	// Unmarshal existing JSON data
	data := make(map[string]string)
	if err := json.Unmarshal(existingData, &data); err != nil {
		return err
	}

	// Update the key
	data[key] = value

	// Marshal the updated data
	updatedData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// Write the updated data back to the file using os.WriteFile
	err = os.WriteFile(filename, updatedData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func CreateConfig() {
	filename := "Config.Json"
	// Check if the file exists
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		// If the file does not exist, create it
		createNewFile(filename)
	}

}

// createNewFile creates a new JSON file with the given key.
func createNewFile(filename string) error {
	data := map[string]string{}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		return err
	}

	return nil
}

// / ReadKeyFromFile reads a key from a JSON file.
func ReadKeyFromFile(filename, key string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	jsonData, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	var data map[string]string
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		return "", err
	}

	value, exists := data[key]
	if !exists {
		return "", fmt.Errorf("key '%s' not found in JSON file", key)
	}

	return value, nil
}

func AddAWSProfile(awsID, awsSecret string) error {
	profileName := "AstroMailApp"
	// Determine the credentials file path.
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("finding home directory: %w", err)
	}
	credsPath := filepath.Join(homeDir, ".aws", "credentials")

	// Load or initialize the credentials file.
	cfg, err := ini.LoadSources(ini.LoadOptions{IgnoreInlineComment: true}, credsPath)
	if err != nil {
		// If the file doesn't exist, ini.LoadSources will not return an error.
		// It returns an empty object ready to be populated.
		cfg = ini.Empty()
	}

	// Get the section (profile) or create it if it doesn't exist.
	section, err := cfg.GetSection(profileName)
	if err != nil {
		if section, err = cfg.NewSection(profileName); err != nil {
			return fmt.Errorf("creating profile section: %w", err)
		}
	}

	// Set the AWS credentials in the profile.
	section.Key("aws_access_key_id").SetValue(awsID)
	section.Key("aws_secret_access_key").SetValue(awsSecret)

	// Save the file.
	if err := cfg.SaveTo(credsPath); err != nil {
		return fmt.Errorf("saving credentials file: %w", err)
	}

	return nil
}
