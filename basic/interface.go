// interface
package basic

import (
	"fmt"
)

type Phone interface {
	call()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, i can call you")
}

func TestInterface() {

	var phone Phone

	phone = new(NokiaPhone)
	phone.call()

}

func testIF() {

}
