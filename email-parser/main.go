package emailparser

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/jhillyerd/enmime"
)

type Email struct {
	From        string       `json:"headers"`
	Text        string       `json:"text"`
	HTML        string       `json:"html"`
	Attachments []Attachment `json:"attachments,omitempty"`
}

type Attachment struct {
	Filename    string `json:"filename"`
	ContentType string `json:"contentType"`
	Content     []byte `json:"content"`
}

func ParseEmail(filename string) string {
	// Read the MIME file
	fileData, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	// Parse the MIME data
	email, err := enmime.ReadEnvelope(bytes.NewReader(fileData))
	if err != nil {
		fmt.Printf("Error parsing MIME data: %v\n", err)
		os.Exit(1)
	}

	// Prepare the JSON structure
	emailJSON := Email{
		From: email.GetHeader("From"),
		Text: email.Text,
		HTML: email.HTML,
	}

	// Attachments
	for _, att := range email.Attachments {
		emailJSON.Attachments = append(emailJSON.Attachments, Attachment{
			Filename:    att.FileName,
			ContentType: att.ContentType,
		})
	}

	// Convert to JSON
	jsonEmail, err := json.MarshalIndent(emailJSON, "", "  ")
	if err != nil {
		fmt.Printf("Error converting to JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(jsonEmail))
	return string(jsonEmail)
}
