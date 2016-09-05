package eventemitter

import (
	"fmt"
	"testing"
)

func TestEventEmitter(t *testing.T) {
	ee := NewEventEmitter()
	ee.On("test1", func(event *Event) {
		fmt.Println("name:", event.Name)
		fmt.Println("data1:", event.Data["data1"])
		fmt.Println("data2:", event.Data["data2"])
	})

	map1 := make(map[string]interface{})
	map1["data1"] = "mydata1"
	map1["data2"] = "mydata2"

	ee.Emit("test1", map1)
}
