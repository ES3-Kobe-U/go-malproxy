package usecase

import (
	"context"
	"io/ioutil"

	"github.com/chromedp/chromedp"
)

func (c *Contents) CheckingContextContents(ctx context.Context) error {
	var imageBuf []byte
	if err := chromedp.Run(ctx,
		chromedp.CaptureScreenshot(&imageBuf),
	); err != nil {
		return err
	}
	if err := ioutil.WriteFile("server/templates/img/now.png", imageBuf, 0644); err != nil {
		return err
	}
	return nil
}
