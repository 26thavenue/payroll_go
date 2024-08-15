package utils

import (
	_"errors"
	"golang.org/x/crypto/bcrypt"
)

func ComparePasswords (hashedPassword, plainPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))

	if err != nil{
		return err
	}

	return  nil
}

func HashPassword(plainPassword string) (string,error){

	hashedPassword,err := bcrypt.GenerateFromPassword([]byte(plainPassword),bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	password := string(hashedPassword)

	return password, nil
}