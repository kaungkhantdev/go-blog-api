package mail

// import (
// 	"bytes"
// 	"html/template"
// 	"log"

// 	"gopkg.in/gomail.v2"
// )

// // EmailConfig holds the SMTP server configuration
// type EmailConfig struct {
// 	SMTPHost    string
// 	SMTPPort    int
// 	Username    string
// 	Password    string
// 	FromAddress string
// }

// // EmailService handles email sending
// type EmailService struct {
// 	Config EmailConfig
// }

// // NewEmailService creates and returns a new EmailService
// func NewEmailService(config EmailConfig) *EmailService {
// 	return &EmailService{Config: config}
// }

// // SendEmail sends an email with optional CC, BCC, and attachments
// func (s *EmailService) SendEmail(
// 	to []string,
// 	cc []string,
// 	bcc []string,
// 	subject string,
// 	templateFile string,
// 	data interface{},
// 	attachments []string,
// ) error {
// 	// Parse the HTML template
// 	tmpl, err := template.ParseFiles(templateFile)
// 	if err != nil {
// 		return err
// 	}

// 	// Execute the template with dynamic data
// 	var body bytes.Buffer
// 	if err := tmpl.Execute(&body, data); err != nil {
// 		return err
// 	}

// 	// Create a new email message
// 	m := gomail.NewMessage()
// 	m.SetHeader("From", s.Config.FromAddress)
// 	m.SetHeader("To", to...)
// 	if len(cc) > 0 {
// 		m.SetHeader("Cc", cc...)
// 	}
// 	if len(bcc) > 0 {
// 		m.SetHeader("Bcc", bcc...)
// 	}
// 	m.SetHeader("Subject", subject)
// 	m.SetBody("text/html", body.String())

// 	// Add attachments if any
// 	for _, attachment := range attachments {
// 		m.Attach(attachment)
// 	}

// 	// Set up the SMTP dialer
// 	d := gomail.NewDialer(s.Config.SMTPHost, s.Config.SMTPPort, s.Config.Username, s.Config.Password)

// 	// Send the email
// 	return d.DialAndSend(m)
// }

// func main() {
// 	// Initialize email configuration
// 	config := EmailConfig{
// 		SMTPHost:    "smtp.example.com",
// 		SMTPPort:    587,
// 		Username:    "your-email@example.com",
// 		Password:    "your-email-password",
// 		FromAddress: "your-email@example.com",
// 	}

// 	// Create email service
// 	emailService := NewEmailService(config)

// 	// Define recipients, CC, BCC, and email data
// 	to := []string{"recipient@example.com"}
// 	cc := []string{"cc-recipient@example.com"}
// 	bcc := []string{"bcc-recipient@example.com"}
// 	subject := "Welcome Email with Attachments"
// 	data := map[string]string{
// 		"Name":    "John Doe",
// 		"Message": "Here's an email with attachments!",
// 	}
// 	attachments := []string{"path/to/file1.pdf", "path/to/file2.jpg"}

// 	// Send the email
// 	err := emailService.SendEmail(to, cc, bcc, subject, "email_template.html", data, attachments)
// 	if err != nil {
// 		log.Fatalf("Failed to send email: %v", err)
// 	}

// 	log.Println("Email sent successfully with CC, BCC, and attachments!")
// }
