package controllers

import (
	"fmt"
	"week9/models"

	gomail "gopkg.in/mail.v2"
)

func NewEmailConfig(host string, port int, senderEmail, senderPassword string) *models.EmailConfig {
	return &models.EmailConfig{
		Host:           host,
		Port:           port,
		SenderEmail:    senderEmail,
		SenderPassword: senderPassword,
	}
}

func SendEmail(config models.EmailConfig, recipientEmail string, subject string, body string) error {
	email := gomail.NewMessage()
	email.SetHeader("From", config.SenderEmail)
	email.SetHeader("To", recipientEmail)
	email.SetHeader("Subject", subject)
	email.SetBody("text/html", body)
	dialer := gomail.NewDialer(config.Host, config.Port, config.SenderEmail, config.SenderPassword)

	if err := dialer.DialAndSend(email); err != nil {
		return err
	}

	return nil
}

func KirimPenambahanPoin(config models.EmailConfig, user *models.User, pointsAdded int) error {
	subject := "Penambahan Poin"
	body := fmt.Sprintf("Halo %s,\n\nAnda telah menerima penambahan %d poin. Total poin saat ini: %d\n\nTerima kasih.", user.Name, pointsAdded, user.Points)
	if err := SendEmail(config, user.Email, subject, body); err != nil {
		return err
	}
	return nil
}

func KirimPenguranganPoin(config models.EmailConfig, user *models.User, pointsUsed int) error {
	subject := "Penggunaan Poin"
	body := fmt.Sprintf("Halo %s,\n\nAnda telah menggunakan %d poin. Total poin saat ini: %d\n\nTerima kasih.", user.Name, pointsUsed, user.Points)
	if err := SendEmail(config, user.Email, subject, body); err != nil {
		return err
	}
	return nil
}
