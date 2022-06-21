package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"mime"
	"mime/multipart"
	"net/mail"
	"net/textproto"
	"os"
	"strings"
)

func main() {
	email := getEmailFromStdin()
	reader := strings.NewReader(email)
	readMsg(reader)
}

func getEmailFromStdin() (email string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		email += scanner.Text() + "\n"
	}
	return
}

func getMediaType(h textproto.MIMEHeader) (mediaType string, params map[string]string) {
	cth := h.Get("Content-Type")
	mediaType, params, err := mime.ParseMediaType(cth)
	if err != nil {
		log.Fatalf("Failed to parse content type header %q: %v", cth, err)
	}
	return mediaType, params
}

func isMultipart(mediaType string) bool {
	return strings.HasPrefix(mediaType, "multipart/")
}

func readMultipart(r *multipart.Reader) {
	for {
		p, err := r.NextPart()
		switch err {
		case nil:
			// Carry on
		case io.EOF:
			return
		default:
			log.Fatalf("%+v", err.Error())
		}

		readPart(p)
	}
}

func readPart(p *multipart.Part) {
	mediaType, params := getMediaType(p.Header)

	if isMultipart(mediaType) {
		boundary := params["boundary"]

		fmt.Printf("Reading multipart with boundary %q\n", boundary)

		pr := multipart.NewReader(p, boundary)
		readMultipart(pr)
		return
	} else if mediaType == "message/rfc822" {
		fmt.Println("Part is attached email")
		readMsg(p)
	}

	fmt.Printf("Not multipart: %s\n", p.Header.Get("Content-Type"))
}

func readMsg(r io.Reader) {
	m, err := mail.ReadMessage(r)
	if err != nil {
		log.Fatalf("%+v", err.Error())
	}

	mediaType, params := getMediaType(textproto.MIMEHeader(m.Header))

	if !isMultipart(mediaType) {
		fmt.Printf("Not multipart: %s\n", mediaType)
		return
	}

	boundary := params["boundary"]

	fmt.Printf("Reading multipart with boundary %q\n", boundary)

	mr := multipart.NewReader(m.Body, boundary)
	readMultipart(mr)
}
