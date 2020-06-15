package utils

import "log"

// CheckError verifica erros
func CheckError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
