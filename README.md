# Go-mailer
Async and sync golang mailer

## Installing
```bash
go get github.com/gronpipmaster/go-mailer
```

## Usage
```go
//create mail server
mailServer := mailer.New("example.com:587", "from", "pass")
//listen and wait messages
go mailServer.Listen()
//create message
subject := "Hi"
body := "Hello <a href=\"#\">Some link</a>"
msg := mailer.NewHtmlMessage(
    []string{"to@example.com"},
    "from@example.com",
    "from username",
    subject,
    body,
)
//send
mailServer.SendAsync(msg)
```