package service

import (
	"io/ioutil"

	"github.com/chromedp/chromedp"
)

func (c *Contents) CheckingContextContents() error {
	var imageBuf []byte
	err := chromedp.Run(*c.CTX,
		chromedp.CaptureScreenshot(&imageBuf),
	)
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile("server/templates/img/now.png", imageBuf, 0644); err != nil {
		return err
	}
	return nil
}
