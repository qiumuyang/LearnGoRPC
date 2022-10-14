package main

import (
	"bufio"
	"log"
	"os"
	"time"
)

// Interface for authentication
type Authentication interface {
	Authenticate(token string) bool
}

// RPC TimeService
type TimeService struct {
	Auth Authentication
}

type TimeServiceRequest struct {
	AuthToken string
}

type TimeServiceResponse struct {
	Status string
	Time   int64
}

func (t *TimeService) GetTime(request *TimeServiceRequest, response *TimeServiceResponse) error {
	if t.Auth == nil {
		response.Status = "error"
	} else if t.Auth.Authenticate(request.AuthToken) {
		response.Status = "ok"
		response.Time = time.Now().Unix()
		log.Printf("TimeService.GetTime: %d %s", response.Time, request.AuthToken)
	} else {
		response.Status = "unauthorized"
		log.Printf("request from unauthorized client: %s", request.AuthToken)
	}
	return nil
}

// An authentication implementation that reads tokens from a file
type FileBasedAuthentication struct {
	Tokens []string
}

func (f *FileBasedAuthentication) Authenticate(token string) bool {
	for _, t := range f.Tokens {
		if t == token {
			return true
		}
	}
	return false
}

func (f *FileBasedAuthentication) LoadTokens(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		f.Tokens = append(f.Tokens, scanner.Text())
	}
	return scanner.Err()
}
