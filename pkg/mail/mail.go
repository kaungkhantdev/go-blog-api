package mail

import (
	"bytes"
	"html/template"

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

func (s *EmailService) SendEmail(
    to []string,
	subject, templateFile string,
	data interface{},
	cc []string,    // Optional
	bcc []string,   // Optional
	attachments []string, // Optional
) error {

    tmpl, err := template.ParseFiles(templateFile)
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