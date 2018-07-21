package parser

import (
	"testing"
	"io/ioutil"
	"fmt"
)

func TestParseCityList(t *testing.T) {
	contents,err := ioutil.ReadFile("cityList_test_data.html")

	if err != nil {
		panic(err)
	}

	result := ParseCourseList(contents)

	fmt.Println(result)


}
