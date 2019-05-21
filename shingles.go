package golang_shingles

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"github.com/kljensen/snowball"
	"strings"
	"unicode"
)

type Lang string

const (
	English   Lang = "english"
	Spanish   Lang = "spanish"
	French    Lang = "french"
	Russian   Lang = "russian"
	Swedish   Lang = "swedish"
	Norwegian Lang = "norwegian"
)

var (
	length = 6
	lang   = English
)

func SetShinglesLength(l int) {
	length = l
}

func SetLanguage(l Lang) {
	lang = l
}

func Compare(line1, line2 string) (percent int, err error) {
	line1, err = snowball.Stem(line1, string(lang), true)
	if err != nil {
		return
	}
	shingles1 := splitShingles(line1)

	line2, err = snowball.Stem(line2, string(lang), true)
	if err != nil {
		return
	}
	shingles2 := splitShingles(line2)

	shinglesLen := len(shingles1)
	diffCount := different(shingles1, shingles2)

	percent = 100 * (shinglesLen - diffCount) / shinglesLen

	return
}

func splitShingles(line string) (shingles []string) {
	var filtered bytes.Buffer

	for _, r := range []rune(line) {
		if unicode.IsLetter(r) || r == ' ' {
			filtered.WriteRune(r)
		}
	}

	var words = strings.Split(filtered.String(), " ")

	max := len(words) - length
	for i := 0; i < max; i++ {
		shingles = append(shingles, getShinglesMD5(words[i:i+length]))
	}

	return
}

func getShinglesMD5(slice []string) string {
	hash := md5.New()
	hash.Write([]byte(strings.Join(slice, "")))

	return hex.EncodeToString(hash.Sum(nil))
}

func different(slices ...[]string) (count int) {
	var different = make(map[string]int)

	for _, slice := range slices {
		for _, v := range slice {
			different[v]++
		}
	}

	for _, c := range different {
		if c != 2 {
			count++
		}
	}

	return
}
