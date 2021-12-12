package hw10programoptimization

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	jsoniter "github.com/json-iterator/go"
)

type User struct {
	ID       int
	Name     string
	Username string
	Email    string
	Phone    string
	Password string
	Address  string
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
	result := make(users, 10_000)

	bufR := bufio.NewReader(r)
	for {
		l, err := bufR.ReadBytes('\n')
		if err != nil {
			if err != io.EOF {
				return result, err
			} else if err == io.EOF && len(l) == 0 {
				break
			}
		}
		var user User
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
	result := make(DomainStat)

	for _, user := range u {
		matched := strings.Contains(user.Email, "."+domain)

		if matched {
			pos := strings.LastIndexByte(user.Email, '@')
			domainEmail := strings.ToLower(user.Email[pos+1:])
			num := result[domainEmail]
			num++
			result[domainEmail] = num
		}
	}

	return result, nil
}
