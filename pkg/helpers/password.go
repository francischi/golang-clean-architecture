package helpers

import (
	"crypto/sha256"
	"encoding/hex"
)

func GetSHA256HashCode(stringMessage string) string {

    message := []byte(stringMessage+"SECRET_KEY")

    hash := sha256.New() 
    
    hash.Write(message)
    
	bytes := hash.Sum(nil)
    
    hashCode := hex.EncodeToString(bytes)
    
    return hashCode
}