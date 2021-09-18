package config

import (
	"os"
	"strings"
	"gopkg.in/ini.v1"
)

// Interface godoc
type Interface interface {
	GetString(section, key string) string
	GetInt(section, key string) int
	GetBool(section, key string) bool
}


func IsPalindrome(sentence string, isCaseSensitive bool) bool  {
	lengthCharacter := len(sentence)
	median := lengthCharacter / 2

	for i := 0; i < median; i++ {
		frontCharacter := string(sentence[i])
		backCharacter := string(sentence[lengthCharacter-(i+1)])

		if !isCaseSensitive {
			frontCharacter = strings.ToLower(frontCharacter)
			backCharacter = strings.ToLower(backCharacter)
		}

		if frontCharacter != backCharacter {
			return false
		}
	}

	return true
}

// NewConfig godoc
func NewConfig() Interface {
	root, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	conf, err := ini.Load(root + "/" + os.Getenv("CGF_FILE"))
	if err != nil {
		panic(err)
	}
	return &cfg{conf}
}

type cfg struct {
	*ini.File
}

func (c *cfg) GetString(section, key string) string {
	return c.Section(section).Key(key).String()
}

func (c *cfg) GetInt(section, key string) int {
	result, err := c.Section(section).Key(key).Int()
	if err != nil {
		panic(err)
	}
	return result
}

func (c *cfg) GetBool(section, key string) bool {
	result, err := c.Section(section).Key(key).Bool()
	if err != nil {
		panic(err)
	}
	return result
}
