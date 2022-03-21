package service

import (
	"io/ioutil"

	"github.com/chromedp/chromedp"
)

func (c *Contents) CheckingContextContents() error {
	var imageBuf []byte
	if err := chromedp.Run(*c.Parent,
		chromedp.CaptureScreenshot(&imageBuf),
	); err != nil {
		return err
	}
	if err := ioutil.WriteFile("server/templates/img/now.png", imageBuf, 0644); err != nil {
		return err
	}
	return nil
}
