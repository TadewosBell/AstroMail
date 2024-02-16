package emailparser

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"mime"
	"mime/multipart"
	"net/mail"
	"strings"
	"time"
)

type Email struct {
	From        string       `json:"from"`
	Subject     string       `json:"subject"`
	To          []string     `json:"to"`
	Text        string       `json:"text"`
	HTML        string       `json:"html"`
	Date        string       `json:"date"`
	Attachments []Attachment `json:"attachments,omitempty"`
}

type Attachment struct {
	Filename    string `json:"filename"`
	ContentType string `json:"contentType"`
	Content     string `json:"content"`
}

func ParseEmail(emailStr string) (string, error) {
	msg, err := mail.ReadMessage(strings.NewReader(emailStr))
	if err != nil {
		return "", fmt.Errorf("error reading message: %v", err)
	}

	// Parse subject and from
	subject := msg.Header.Get("Subject")
	from := msg.Header.Get("From")
	date := msg.Header.Get("Date")
	// Decode RFC 2047 encoded strings if necessary
	dec := new(mime.WordDecoder)
	subject, err = dec.DecodeHeader(subject)
	if err != nil {
		fmt.Printf("Failed to decode subject: %v\n", err)
	}
	from, err = dec.DecodeHeader(from)
	if err != nil {
		fmt.Printf("Failed to decode from: %v\n", err)
	}

	// Parse recipients
	var to []string
	for _, addr := range msg.Header["To"] {
		to = append(to, addr)
	}

	var text, html string

	mediaType, params, err := mime.ParseMediaType(msg.Header.Get("Content-Type"))
	if err != nil {
		return "", fmt.Errorf("error parsing media type: %v", err)
	}

	if strings.HasPrefix(mediaType, "multipart/") {
		mr := multipart.NewReader(msg.Body, params["boundary"])
		for {
			part, err := mr.NextPart()
			if err != nil {
				break
			}

			partMediaType, _, err := mime.ParseMediaType(part.Header.Get("Content-Type"))
			if err != nil {
				fmt.Printf("Error parsing part media type: %v\n", err)
				continue
			}

			partBody, err := io.ReadAll(part)
			if err != nil {
				fmt.Printf("Error reading part body: %v\n", err)
				continue
			}

			switch partMediaType {
			case "text/plain":
				text = string(partBody)
			case "text/html":
				html = string(partBody)
			}
		}
	} else {
		// Handle non-multipart emails here if needed
	}

	email := Email{
		From:    from,
		Subject: subject,
		To:      to,
		Text:    text,
		HTML:    html,
		Date:    date,
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

// CreateEMLString creates an email EML string with the given parameters.
func CreateEMLString(sender, subject, body string, recipient, ccAddresses []string, messageId string) string {
	var b strings.Builder

	// Get current date and time in the desired format
	now := time.Now().Format(time.RFC1123Z)

	// Write headers
	fmt.Fprintf(&b, "MIME-Version: 1.0\n")
	fmt.Fprintf(&b, "From: %s <%s>\n", sender, sender)
	fmt.Fprintf(&b, "Date: %s\n", now)
	fmt.Fprintf(&b, "Message-ID: <%s@mail.gmail.com>\n", messageId)
	fmt.Fprintf(&b, "Subject: %s\n", subject)
	fmt.Fprintf(&b, "To: %s\n", strings.Join(recipient, ", "))
	if len(ccAddresses) > 0 {
		fmt.Fprintf(&b, "Cc: %s\n", strings.Join(ccAddresses, ", "))
	}
	fmt.Fprintf(&b, "Content-Type: multipart/alternative; boundary=\"000000000000f0dd9b06111072ad\"\n\n")

	// Write body parts
	fmt.Fprintf(&b, "--000000000000f0dd9b06111072ad\n")
	fmt.Fprintf(&b, "Content-Type: text/plain; charset=\"UTF-8\"\n\n")
	fmt.Fprintf(&b, "%s\n", body)

	fmt.Fprintf(&b, "--000000000000f0dd9b06111072ad\n")
	fmt.Fprintf(&b, "Content-Type: text/html; charset=\"UTF-8\"\n\n")
	fmt.Fprintf(&b, "%s\n", body)

	fmt.Fprintf(&b, "--000000000000f0dd9b06111072ad--\n")

	return b.String()
}
