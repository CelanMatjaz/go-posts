package utils

import (
	"log"
	"os"
	"time"
)

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}
}

func FormatDate(t time.Time) string {
	return t.Format(time.DateTime)
}
