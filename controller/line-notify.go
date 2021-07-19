package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type lineNotifyController interface {
	SendNotify(c *fiber.Ctx) error
}

type LineNotify struct {
	Project string `json:"project" xml:"project" form:"project"`
	Level   string `json:"level" xml:"level" form:"level"`
	Url     string `json:"url" xml:"url" form:"url"`
	Message string `json:"message" xml:"message" form:"message"`
}

type NotifyPayload struct {
	Message string `json:"message"`
	Age     uint8
}

const urlLINEAPI = "https://notify-api.line.me/api"

func NewLineNotify() lineNotifyController {
	return &LineNotify{}
}

func (l *LineNotify) SendNotify(c *fiber.Ctx) error {
	var body = new(LineNotify)
	var token = c.Params("token")
	if err := c.BodyParser(&body); err != nil {
		return err
	}

	message := "\n" + "Project: " + body.Project + "\n" + "Level: " + body.Level + "\n" + "Url: " + body.Url + "\n" + "Message: " + body.Message + "\n"

	params := url.Values{}
	params.Add("message", message)

	req, err := http.NewRequest(http.MethodPost, urlLINEAPI+"/notify", strings.NewReader(params.Encode()))
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	if err != nil {
		fmt.Printf("%v", err)
		return c.Status(500).JSON(err)
	}
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Printf("%v", err)
		return c.Status(500).JSON(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		respBody, _ := ioutil.ReadAll(resp.Body)
		return c.Status(resp.StatusCode).JSON(fiber.Map{
			"body": respBody,
		})
	} else {
		return c.Status(resp.StatusCode).JSON(fiber.Map{
			"body": resp.Body.Close(),
		})
	}
}
