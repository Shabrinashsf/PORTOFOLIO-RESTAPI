package utils

import (
	"bytes"
	"crypto/tls"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/Shabrinashsf/PORTOFOLIO-RESTAPI/entity"
	"github.com/k3a/html2text"
	"gopkg.in/gomail.v2"
)

type EmailData struct {
	URL     string
	Name    string
	Subject string
}

var (
	cachedTemplate *template.Template
	once           sync.Once
)

func LoadTemplates(dir string) {
	once.Do(func() {
		var paths []string
		err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				paths = append(paths, path)
			}
			return nil
		})

		if err != nil {
			log.Fatal("Failed to load templates:", err)
		}

		cachedTemplate, err = template.ParseFiles(paths...)
		if err != nil {
			log.Fatal("Failed to parse templates:", err)
		}

		log.Println("Templates loaded successfully.")
	})
}

func GetCachedTemplates() *template.Template {
	return cachedTemplate
}

/* tidak optimal, tiap ngirim akan ngeparsing
func ParseTemplateDir(dir string) (*template.Template, error) {
	var paths []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return template.ParseFiles(paths...)
}
*/

func SendEmail(user *entity.User, data *EmailData) {
	var body bytes.Buffer

	template := GetCachedTemplates()
	if template == nil {
		log.Println("Template not loaded properly.")
		return
	}

	if err := template.ExecuteTemplate(&body, "verificationCode.html", &data); err != nil {
		log.Fatal("Error executing template:", err)
		return
	}

	SENDER := os.Getenv("MAIL_SENDER")
	PASS := os.Getenv("APP_PASS")

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", SENDER)
	mailer.SetHeader("To", user.Email)
	mailer.SetHeader("Subject", data.Subject)
	mailer.SetBody("text/html", body.String())
	mailer.AddAlternative("text/plain", html2text.HTML2Text(body.String()))

	dialer := gomail.NewDialer("smtp.gmail.com", 587, SENDER, PASS)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := dialer.DialAndSend(mailer); err != nil {
		log.Println("Could not send email: ", err)
	}
}
