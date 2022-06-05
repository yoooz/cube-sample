package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/samber/lo"
)

type Record struct {
	A int
	B int
	C int
	D int
	E int
	F int
	G int
	H int
	I int
	J int
}

func getRandom() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(20)
}

func getDataPath() string {
	p, _ := os.Getwd()

	pwd := strings.Split(p, "/")
	pwd = pwd[0:len(pwd)-3]
	pwd = append(pwd, []string{"db", "init", "sample.csv"}...)
	return strings.Join(pwd, "/")
}

func writeRecord(record []Record) {
	body := ""
	for _, r := range record {
		body += describeRecord(r)
	}
	ioutil.WriteFile(getDataPath(), []byte(body), os.ModePerm)
}

func describeRecord(r Record) string {
	return fmt.Sprintf("\"%d\",\"%d\",\"%d\",\"%d\",\"%d\",\"%d\",\"%d\",\"%d\",\"%d\",\"%d\"\n", r.A, r.B, r.C, r.D, r.E, r.F, r.G, r.H, r.I, r.J)
}

func createRecord() Record {
	return Record{
		A: getRandom(),
		B: getRandom(),
		C: getRandom(),
		D: getRandom(),
		E: getRandom(),
		F: getRandom(),
		G: getRandom(),
		H: getRandom(),
		I: getRandom(),
		J: getRandom(),
	}
}

func main() {
	recordList := lo.Map(make([]bool, 1000), func(t bool, i int) Record { return createRecord() })
	writeRecord(recordList)
}
