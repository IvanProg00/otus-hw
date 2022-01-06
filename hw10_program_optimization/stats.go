package hw10programoptimization

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	jsoniter "github.com/json-iterator/go"
)

type User struct {
	Email string
}

type DomainStat map[string]int

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	u, err := getUsers(r)
	if err != nil {
		return nil, fmt.Errorf("get users error: %w", err)
	}
	return countDomains(u, domain)
}

type users []User

func getUsers(r io.Reader) (users, error) {
	result := make(users, 0, 14_000)
	bufR := bufio.NewReader(r)
	var user User
	jsoniter := jsoniter.ConfigFastest

	for {
		l, _, err := bufR.ReadLine()
		// l, err := bufR.ReadBytes('\n')
		if err != nil {
			if err != io.EOF {
				return result, err
			} else if err == io.EOF && len(l) == 0 {
				break
			}
		}

		if err = jsoniter.Unmarshal(l, &user); err != nil {
			return result, err
		}
		result = append(result, user)
		if err == io.EOF {
			break
		}
	}

	return result, nil
}

func countDomains(u users, domain string) (DomainStat, error) {
	result := make(DomainStat, 1_000)
	domain = "." + domain

	for _, user := range u {
		if strings.HasSuffix(user.Email, domain) {
			pos := strings.IndexByte(user.Email, '@')
			result[strings.ToLower(user.Email[pos+1:])]++
		}
	}

	return result, nil
}
