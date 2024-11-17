package mail

import (
	"bytes"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
	"gopkg.in/gomail.v2"
)

type EmailConfig struct {
	SMTPHost    string
	SMTPPort    int
	Username    string
	Password    string
	FromAddress string
}

type EmailService struct {
    Config EmailConfig
}

func NewEmailService(config EmailConfig) *EmailService {
    return &EmailService{ Config: config }
}

func NewEmailConfig() EmailConfig {
    return EmailConfig{
        SMTPHost:    os.Getenv("MAIL_HOST"),
        SMTPPort:    parseEnvInt(os.Getenv("MAIL_PORT"), 465),
        Username:    os.Getenv("MAIL_USER"),
        Password:    os.Getenv("MAIL_PASS"),
        FromAddress: os.Getenv("MAIL_USER"),
    }
}

func parseEnvInt(key string, defaultValue int) int {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	parsedValue, err := strconv.Atoi(value)
	if err != nil {
		log.Printf("Invalid integer for %s: %v. Using default value %d", key, err, defaultValue)
		return defaultValue
	}

	return parsedValue
}


func (s *EmailService) SendEmail(
    to []string,
	subject, templateName string,
	data interface{},
	cc []string,    // Optional
	bcc []string,   // Optional
	attachments []string, // Optional
) error {

    templatePath := filepath.Join("pkg", "mail", "templates", templateName)

    tmpl, err := template.ParseFiles(templatePath)
    if err != nil {
        return err
    }

    var body bytes.Buffer
    if err := tmpl.Execute(&body, data); err != nil {
        return err
    }

    mail := gomail.NewMessage()
    mail.SetHeader("From", s.Config.FromAddress)
    mail.SetHeader("To", to...)
    mail.SetHeader("Subject", subject)
    mail.SetBody("text/html", body.String())

    if len(cc) > 0 {
        mail.SetHeader("Cc", cc...)
    }

    if len(bcc) > 0 {
        mail.SetHeader("Bcc", bcc...)
    }

    for _, attachment :=range attachments {
        mail.Attach(attachment)
    }

    smtpDialer := gomail.NewDialer(s.Config.SMTPHost, s.Config.SMTPPort, s.Config.Username, s.Config.Password)

    return smtpDialer.DialAndSend(mail);
}