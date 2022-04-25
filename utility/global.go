package utility

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time"
)

func AddZeroAndToString(number int) string {
	if number < 10 {
		return fmt.Sprintf("0%d", number)
	}
	return fmt.Sprintf("%d", number)

}

func ThisDay() string {
	year, month, day := time.Now().Date()

	stringThisDay := fmt.Sprintf("%v-%s-%s", year, AddZeroAndToString(int(month)), AddZeroAndToString(day))
	return stringThisDay
}

func ThisTimeStamp() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func ThisTimeStampCode() string {
	time := time.Now().Format("2006-01-02 15:04:05")
	stringReplace := strings.Replace(time, "-", "", -1)
	stringReplace = strings.Replace(stringReplace, ":", "", -1)
	stringReplace = strings.Replace(stringReplace, " ", "", -1)
	return stringReplace
}

func CreateNameFile(name string) string {
	extension := strings.Split(name, ".")
	fileName := fmt.Sprintf("%s.%s", ThisTimeStampCode(), extension[len(extension)-1])
	return fileName
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
