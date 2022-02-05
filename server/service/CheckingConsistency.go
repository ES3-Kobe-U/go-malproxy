package service

import (
	"context"
	"os"
	"time"

	"github.com/chromedp/chromedp"
)

/* ※WSLでは動かない(Chromeがない)．
Amazon用の処理
*/
func CheckingTheIntegrityOfAmazonInformation(email string, password string) error {
	var res []byte
	ctx, cancel := chromedp.NewContext(context.Background(), chromedp.WithBrowserOption())
	defer cancel()
	err := chromedp.Run(ctx,
		chromedp.Navigate("https://www.amazon.co.jp/stores/page/FD12D065-F94A-4656-B816-4C8121739FAE?ingress=3"),
		chromedp.WaitReady("body"),
		chromedp.Click(`a[data-nav-role='signin']`, chromedp.ByQuery),
		chromedp.Sleep(time.Second*2),
		chromedp.SetValue(`ap_email`, email, chromedp.ByID),
		chromedp.Click(`continue`, chromedp.ByID),
		chromedp.Sleep(time.Second*1),
		chromedp.SetValue(`ap_password`, password, chromedp.ByID),
		chromedp.Click(`signInSubmit`, chromedp.ByID),
		chromedp.Sleep(time.Second*2),
		chromedp.CaptureScreenshot(&res),
	)
	if err != nil {
		return err
	}
	os.WriteFile("./loggedin.png", res, 0644)
	return nil
}
