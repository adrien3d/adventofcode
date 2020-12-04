package utils

import (
	"github.com/sirupsen/logrus"
	"github.com/snwfdhmp/errlog"
)

// CheckErr checks error and print it if it exists
func CheckErr(e error) {
	if e != nil {
		errlog.Debug(e)
	}
}

// Log logs if debug env var is set at true
func Log(verbose bool, level string, msg ...interface{}) {
	if verbose {
		switch level {
		case "info":
			logrus.Infoln(msg...)
		case "warn":
			logrus.Warnln(msg...)
		case "err":
			logrus.Errorln(msg...)
		default:
			logrus.Infoln(msg...)
		}
	}
}

// RemoveStringFromSlice allows to remove a string in a slice
func RemoveStringFromSlice(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

// FindStringInSlice allows to find a string in a slice
func FindStringInSlice( val string, slice []string) (isFound bool) {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}