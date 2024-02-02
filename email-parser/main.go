package emailparser

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime"
	"mime/multipart"
	"net/mail"
	"os"
	"strings"
)

type Email struct {
	From        string       `json:"from"`
	Subject     string       `json:"subject"`
	To          []string     `json:"to"`
	Text        string       `json:"text"`
	HTML        string       `json:"html"`
	Attachments []Attachment `json:"attachments,omitempty"`
}

type Attachment struct {
	Filename    string `json:"filename"`
	ContentType string `json:"contentType"`
	Content     string `json:"content"`
}

func ParseEmail(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	msg, err := mail.ReadMessage(file)
	if err != nil {
		return "", fmt.Errorf("error reading message: %v", err)
	}

	// Parse subject and from
	subject := msg.Header.Get("Subject")
	from := msg.Header.Get("From")

	// Decode RFC 2047 encoded strings if necessary
	dec := new(mime.WordDecoder)
	subject, err = dec.DecodeHeader(subject)
	if err != nil {
		log.Printf("Failed to decode subject: %v", err)
	}
	from, err = dec.DecodeHeader(from)
	if err != nil {
		log.Printf("Failed to decode from: %v", err)
	}

	// Parse recipients
	var to []string
	for _, addr := range msg.Header["To"] {
		to = append(to, addr)
	}

	email := Email{
		From:    from,
		Subject: subject,
		To:      to,
	}

	mediaType, params, err := mime.ParseMediaType(msg.Header.Get("Content-Type"))
	if err != nil {
		return "", fmt.Errorf("error parsing media type: %v", err)
	}

	if strings.HasPrefix(mediaType, "multipart/") {
		mr := multipart.NewReader(msg.Body, params["boundary"])
		if err := parseParts(mr, &email); err != nil {
			return "", err
		}
	} else {
		// Handle non-multipart emails here if needed
	}

	jsonEmail, err := json.MarshalIndent(email, "", "  ")
	if err != nil {
		return "", fmt.Errorf("error marshalling to JSON: %v", err)
	}

	return string(jsonEmail), nil
}

func parseParts(reader *multipart.Reader, email *Email) error {
	for {
		part, err := reader.NextPart()
		if err == io.EOF {
			break
		} else if err != nil {
			return fmt.Errorf("error reading part: %v", err)
		}

		contentType := part.Header.Get("Content-Type")
		contentTransferEncoding := part.Header.Get("Content-Transfer-Encoding")

		body, err := io.ReadAll(part)
		if err != nil {
			return fmt.Errorf("error reading part body: %v", err)
		}

		switch {
		case strings.HasPrefix(contentType, "text/plain"):
			email.Text += string(body)

		case strings.HasPrefix(contentType, "text/html"):
			email.HTML += string(body)

		case strings.Contains(contentType, "attachment") || strings.Contains(part.Header.Get("Content-Disposition"), "attachment"):
			filename := part.FileName()
			var content string
			switch contentTransferEncoding {
			case "base64":
				content = base64.StdEncoding.EncodeToString(body)
			default:
				content = string(body)
			}

			email.Attachments = append(email.Attachments, Attachment{
				Filename:    filename,
				ContentType: contentType,
				Content:     content,
			})
		}

		// Check for nested multipart parts
		// if strings.HasPrefix(contentType, "multipart/") {
		// 	nestedMr := multipart.NewReader(bytes.NewReader(body), params["boundary"])
		// 	if err := parseParts(nestedMr, email); err != nil {
		// 		return err
		// 	}
		// }
	}
	return nil
}
