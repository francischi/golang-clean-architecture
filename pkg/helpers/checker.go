package helpers

import "net/mail"

func IsValidPassword(pwd string)bool{
    return len(pwd) > 6
}

func IsValidEmail(email string) bool {
    _, err := mail.ParseAddress(email)
    return err == nil
}