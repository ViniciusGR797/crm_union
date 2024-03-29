package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/go-gomail/gomail"
)

type ConfigMail struct {
	SendMail struct {
		SMTPServer   string `json:"smtpserver"`
		SMTPPort     int    `json:"smtpport"`
		SMTPUsername string `json:"smtpusername"`
		SMTPPassword string `json:"smtppassword"`
	} `json:"sendmail"`
}

func sendMail(to, subject, body string) {

	filePath := "/home/appuser/app/env.json" // caminho para o arquivo JSON
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Erro ao ler o arquivo JSON:", err)
		return
	}

	// Decodificar o arquivo JSON em uma estrutura Config
	var config ConfigMail
	err = json.Unmarshal(content, &config)
	if err != nil {
		fmt.Println("Erro ao decodificar o arquivo JSON:", err)
		return
	}

	// Recuperar as informações em string
	smtpServer := config.SendMail.SMTPServer
	smtpPort := config.SendMail.SMTPPort
	smtpUsername := config.SendMail.SMTPUsername
	smtpPassword := config.SendMail.SMTPPassword

	// Criar um objeto gomail.Message com as informações do e-mail
	message := gomail.NewMessage()
	message.SetHeader("From", smtpUsername)
	message.SetHeader("To", to)
	message.SetHeader("Subject", subject)
	message.SetBody("text/plain", body)

	// Configurar a autenticação OAuth2
	d := gomail.NewDialer(smtpServer, smtpPort, smtpUsername, smtpPassword)

	// Autenticar e enviar o e-mail
	if err := d.DialAndSend(message); err != nil {
		log.Fatal(err)
	}

	fmt.Println("E-mail enviado com sucesso!")
}
