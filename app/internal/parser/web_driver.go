package parser

import (
	"fmt"
	"log"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

const (
	maxTries = 10
)

func NewWebDriver() selenium.WebDriver {
	caps := selenium.Capabilities{
		"browserName": "chrome",
	}

	chrCaps := chrome.Capabilities{
		W3C: true,
	}
	caps.AddChrome(chrCaps)

	var wd selenium.WebDriver
	var err error

	fmt.Println("сейчас будет ЦИКЛ")
	i := 0
	for i < maxTries {
		wd, err = selenium.NewRemote(caps, "http://chrome:4444/wd/hub")
		if err != nil {
			log.Println("ошибка при создании драйвера:", err)
			i++
			continue
		}
		break
	}

	if wd == nil {
		log.Fatal("не удалось подключиться к драйверу")
	}

	log.Println("подключение в драйверу произведено")

	return wd
}
