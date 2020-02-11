package utilities_test

import (
	"go-task-list/utilities"
	"testing"
)

func TestDateCheckPass(t *testing.T) {
	result := utilities.CheckDateStringFormat("2017-02-02")
	if result == false {
		t.Fatal("2017-02-02 should have passes but it didnt!")
	}
}

func TestDateCheckFail1(t *testing.T) {
	result := utilities.CheckDateStringFormat("2017-02esdfsd-02")
	if result == true {
		t.Fatal("2017-02esdfsd-02 shouldn't have passed but it did!")
	}
}

func TestDateCheckFail2(t *testing.T) {
	result := utilities.CheckDateStringFormat("201702")
	if result == true {
		t.Fatal("201702 shouldn't have passed but it did!")
	}
}

func TestDateCheckFail3(t *testing.T) {
	result := utilities.CheckDateStringFormat("")
	if result == true {
		t.Fatal("Empty string shouldn't have passed but it did!")
	}
}

func TestDateCheckFail4(t *testing.T) {
	result := utilities.CheckDateStringFormat("1950-02-02")
	if result == true {
		t.Fatal("2017-02esdfsd-02 shouldn't have passed but it did!")
	}
}

func TestDateCheckFail5(t *testing.T) {
	result := utilities.CheckDateStringFormat("2017-19-02")
	if result == true {
		t.Fatal("2017-02esdfsd-02 shouldn't have passed but it did!")
	}
}

func TestDateCheckFail6(t *testing.T) {
	result := utilities.CheckDateStringFormat("2017-02-44")
	if result == true {
		t.Fatal("2017-02esdfsd-02 shouldn't have passed but it did!")
	}
}
