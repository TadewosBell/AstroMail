package main

import (
	storage "AstroMail/config"
	emailparser "AstroMail/email-parser"
	smtpstack "AstroMail/smtp-stack"
	"fmt"
	"strings"

	// "AstroMail/storage"
	"context"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {

	storage.CreateConfig()
	a.ctx = ctx
}

// Launch SMTP Server
func (a *App) Launch_Smtp_Server(username, domain, aws_id, aws_secret string) {
	fmt.Println(domain, aws_id, aws_secret)
	storage.WriteKeyToFile("Username", username, "Config.Json")
	// storage.AddAWSProfile(aws_id, aws_secret)
	bucket, _ := smtpstack.CreateEmailBucket(domain)

	storage.WriteKeyToFile("Bucket", bucket, "Config.Json")
	storage.WriteKeyToFile("Bucket Status", "Created", "Config.Json")
	storage.WriteKeyToFile("Domain", domain, "Config.Json")
	storage.WriteKeyToFile("Domain Status", "Verifying", "Config.Json")
	fmt.Println("bucket created: ", bucket)
	verificationStatus, _ := smtpstack.IsDomainVerified(domain)
	storage.WriteKeyToFile("Domain Status", "Verified", "Config.Json")
	fmt.Println("Verification Status: ", verificationStatus)

	roleArn, err := smtpstack.CreateSESPolicyAndRole(domain, bucket)
	storage.WriteKeyToFile("RoleArn", roleArn, "Config.Json")
	fmt.Println("CreateSESPolicyAndRole failed: ", err)

	err = smtpstack.ConfigureSESReceiptRules(domain, username, roleArn, bucket)

	fmt.Println("Configure ses rec failed: ", err)
	if err != nil {
		return
	}
	storage.WriteKeyToFile("Status", "Working", "Config.Json")
}

// Launch SMTP Server
func (a *App) Send_Email(to string, subject string, body string) {
	username, err := storage.ReadKeyFromFile("Config.Json", "Username")
	err = smtpstack.SendEmail(username+"@astrocommits.com", subject, body, []string{to}, []string{})
	if err == nil {
		runtime.EventsEmit(a.ctx, "emailSent")
	}
	// else return a message with an error emit

}

// Greet returns a greeting for the given name
func (a *App) Get_Inbox() []string {
	bucket, _ := storage.ReadKeyFromFile("Config.Json", "Bucket")
	objectNames, err := smtpstack.ReadBucketFolderContent(bucket, "email", 1)
	for _, filePath := range objectNames {
		filepathSplit := strings.Split(filePath, "/")
		messageId := filepathSplit[1]
		if messageId != "AMAZON_SES_SETUP_NOTIFICATION" {
			content, err := smtpstack.GetObjectContentAsString(bucket, filePath)
			if err != nil {
				fmt.Printf("Error reading object %s: %v\n", filePath, err)
				continue
			}
			err = storage.SaveEmail(messageId, content, "inbox")
			if err != nil {
				fmt.Println(err)
				continue
			}
			_, err = emailparser.ParseEmail(content)
		}

	}

	emails, err := storage.RetrieveEmailsPaginated("inbox", 1, 50)
	var inbox []string
	for _, email := range emails {
		emailObj, _ := emailparser.ParseEmail(email)
		inbox = append(inbox, emailObj)
	}

	if err != nil {
		return inbox
	}

	return inbox
}

func (a *App) Get_Next_Page(folder string, page int) []string {
	fmt.Println("Page: ", page)
	emails, err := storage.RetrieveEmailsPaginated(folder, page, 50)
	var items []string
	for _, email := range emails {
		fmt.Println(email)
		emailObj, _ := emailparser.ParseEmail(email)
		items = append(items, emailObj)
	}

	if err != nil {
		return items
	}

	return items
}

// func (a *App) Get_Inbox() []string {
// 	emails, err := storage.RetrieveEmailsPaginated("inbox", 1, 50)
// 	fmt.Println(emails)
// 	var inbox []string
// 	for _, email := range emails {
// 		fmt.Println(email)
// 		emailObj, _ := emailparser.ParseEmail(email)
// 		inbox = append(inbox, emailObj)
// 	}

// 	if err != nil {
// 		return inbox
// 	}

// 	return inbox
// }

func (a *App) Get_Sent() []string {
	emails, err := storage.RetrieveEmailsPaginated("sent", 1, 50)
	var sent []string
	for _, email := range emails {
		emailObj, _ := emailparser.ParseEmail(email)
		sent = append(sent, emailObj)
	}

	if err != nil {
		return sent
	}

	return sent
}

func (a *App) Is_Setup() bool {
	status, _ := storage.ReadKeyFromFile("Config.Json", "Status")

	if status == "Working" {
		return true
	} else {
		return false
	}
}
