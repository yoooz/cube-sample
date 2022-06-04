package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/samber/lo"
)

type TaskStatusType int

const (
	Preparation TaskStatusType = 0
	Start       TaskStatusType = 1
	Done        TaskStatusType = 2
)

type User struct {
	Name string `faker:"first_name"`
}

type StatusTime struct {
	Time string `faker:"time"`
}

type Task struct {
	Id string `faker:"word"`
}

type TaskStatus struct {
	Date           string
	User           string
	Task           string
	TaskStatusType TaskStatusType
	Time           string
}

func createDate(count int) (result []string) {
	base := time.Now()
	return lo.Map(make([]bool, count), func(t bool, i int) string {
		return base.AddDate(0, 0, i).Format("2006-01-02")
	})
}

func createFakerUser() string {
	user := User{}
	faker.FakeData(&user)
	return user.Name
}

func createFakeTime() string {
	statusTime := StatusTime{}
	faker.FakeData(&statusTime)
	return statusTime.Time
}

func createTask() string {
	task := Task{}
	faker.FakeData(&task)
	return task.Id
}

func createSortedFakeTime(count int) (result []string) {
	for i := 0; i < count; i++ {
		fake := createFakeTime()
		result = append(result, fake)
	}

	sort.Slice(result, func(i, j int) bool {
		format := "15:04:05"
		t1, _ := time.Parse(format, result[i])
		t2, _ := time.Parse(format, result[j])
		return t1.Before(t2)
	})

	return result
}

func createHistory(date, user string, count int) (result []TaskStatus) {
	timeList := createSortedFakeTime(count * 3)
	statusTypeList := []TaskStatusType{Preparation, Start, Done}

	task := createTask()
	for i, timestamp := range timeList {
		result = append(result, TaskStatus{
			Date:           date,
			User:           user,
			Task:           task,
			TaskStatusType: statusTypeList[i%3],
			Time:           fmt.Sprintf("%s %s", date, timestamp),
		})
		if i%3 == 2 {
			task = createTask()
		}
	}
	return result
}

func describeHistory(history TaskStatus) string {
	return fmt.Sprintf("\"%s\",\"%s\",\"%s\",%d,\"%s\"\n", history.Date, history.User, history.Task, history.TaskStatusType, history.Time)
}

func printHistory(history TaskStatus) {
	fmt.Printf(describeHistory(history))
}

func getDataPath() string {
	p, _ := os.Getwd()

	pwd := strings.Split(p, "/")
	pwd[len(pwd)-1] = "db"
	pwd = append(pwd, []string{"init", "task.csv"}...)
	return strings.Join(pwd, "/")
}

func writeHistory(history []TaskStatus) {
	body := ""
	for _, h := range history {
		body += describeHistory(h)
	}
	ioutil.WriteFile(getDataPath(), []byte(body), os.ModePerm)
}

func createData() {
	userList := lo.Map(make([]bool, 10), func(t bool, i int) string { return createFakerUser() })
	historyList := []TaskStatus{}
	for _, date := range createDate(10) {
		for _, user := range userList {
			for _, history := range createHistory(date, user, 10) {
				historyList = append(historyList, history)
			}
		}
	}
	writeHistory(historyList)
}

func main() {
	createData()
}
