package email

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
	"html/template"
	"log"
	"os"
	"time"
)

func SendEmail(location string) error {
	config := oauth2.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		Endpoint:     google.Endpoint,
		RedirectURL:  "http://localhost",
	}

	token := oauth2.Token{
		AccessToken:  os.Getenv("ACCESS_TOKEN"),
		RefreshToken: os.Getenv("REFRESH_TOKEN"),
		TokenType:    "Bearer",
		Expiry:       time.Now(),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client := config.Client(ctx, &token)
	sendService, err := gmail.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve gmail Client %v", err)
	}

	header := make(map[string]string)
	header["To"] = os.Getenv("EMAIL")
	header["Subject"] = "Invoice for your order"
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = `text/html; charset="utf-8"`
	header["Content-Transfer-Encoding"] = "base64"

	t, err := template.ParseFiles("email/template.html")
	if err != nil {
		return err
	}

	var body bytes.Buffer
	err = t.Execute(&body, struct {
		URL string
	}{
		URL: location,
	})
	if err != nil {
		return err
	}

	var msg string
	for k, v := range header {
		msg += fmt.Sprintf("%s: %s\n", k, v)
	}
	msg += "\n" + body.String()

	_, err = sendService.Users.Messages.Send("me", &gmail.Message{
		Raw: encodeWeb64String([]byte(msg)),
	}).Do()
	return err
}

func encodeWeb64String(b []byte) string {
	s := base64.URLEncoding.EncodeToString(b)

	var i = len(s) - 1
	for s[i] == '=' {
		i--
	}

	return s[0 : i+1]
}
