package service

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/chromedp"
)

// ※WSLでは動かない(Chromeがない)．

/*
Amazon用の処理
*/
func CheckingTheIntegrityOfAmazonInformation(email string, password string) error {
	var res string
	var res0 []byte
	var res1 []byte
	var res2 []byte
	var res3 []byte
	ctx, cancel := chromedp.NewContext(context.Background(), chromedp.WithBrowserOption())
	defer cancel()
	task1 := chromedp.Tasks{ //タスクリストの作成
		chromedp.Navigate("https://www.amazon.co.jp/?&tag=hydraamazonav-22&ref=pd_sl_2ykkalld4i_e&adgrpid=54841807378&hvpone=&hvptwo=&hvadid=289239574720&hvpos=&hvnetw=g&hvrand=13443261105670128409&hvqmt=e&hvdev=c&hvdvcmdl=&hvlocint=&hvlocphy=1009565&hvtargid=kwd-333588672930&hydadcr=15460_10908920&gclid=Cj0KCQiAjJOQBhCkARIsAEKMtO0jVYrB9RxzmNKNeCPzZE0CB_TUL10D5UonY9FkHd4maUGPDrYDe4UaAnhwEALw_wcB"),
		chromedp.WaitReady("body"),
		chromedp.CaptureScreenshot(&res0),
		chromedp.Click(`a[data-nav-role='signin']`, chromedp.ByQuery),
		chromedp.Sleep(time.Second * 2),
		chromedp.SetValue(`ap_email`, email, chromedp.ByID),
		chromedp.Click(`continue`, chromedp.ByID),
		chromedp.Sleep(time.Second * 1),
		chromedp.SetValue(`ap_password`, password, chromedp.ByID),
		chromedp.CaptureScreenshot(&res1),
		chromedp.Click(`signInSubmit`, chromedp.ByID),
		chromedp.Sleep(time.Second * 2),
		chromedp.CaptureScreenshot(&res2),
		chromedp.ActionFunc(func(ctx context.Context) error {
			node, err := dom.GetDocument().Do(ctx)
			if err != nil {
				return err
			}
			res, err = dom.GetOuterHTML().WithNodeID(node.NodeID).Do(ctx)
			if err != nil {
				return err
			}
			return nil
		}),
	}
	err := chromedp.Run(ctx, task1)
	if err != nil {
		err := AmaoznCaptcha(email, password)
		if err != nil {
			return err
		}
	}
	task2 := chromedp.Tasks{
		chromedp.CaptureScreenshot(&res3),
	}
	err = chromedp.Run(ctx, task2) //セッションが維持されているかの確認
	if err != nil {
		return err
	}
	os.WriteFile("./res0.png", res0, 0644)
	os.WriteFile("./res1.png", res1, 0644)
	os.WriteFile("./res2.png", res2, 0644)
	output := `{{define "autogen_amazon_info"}}` + res + `{{end}}`
	err = os.WriteFile("server/templates/autogen_amazon_login.html", []byte(output), 0644)
	if err != nil {
		return err
	}
	return nil
}

func AmaoznCaptcha(email string, password string) error {
	var res4 []byte
	var res5 []byte
	var data string
	var str string = "xxxxxx"
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
		chromedp.CaptureScreenshot(&res4),
		chromedp.SetValue(`ap_password`, password, chromedp.ByID),
		chromedp.ActionFunc(func(c context.Context) error {
			fmt.Println("output --> res1.png")
			err := os.WriteFile("./res1.png", res4, 0644)
			if err != nil {
				return err
			}
			return nil
		}),
		chromedp.Sleep(time.Second*10),
		chromedp.ActionFunc(func(c context.Context) error {
			var input string
			fmt.Print("fmt.Scan >")
			fmt.Scanf("%s", &input)
			str = input
			fmt.Println("input:", str)
			return nil
		}),
		chromedp.Sleep(time.Second*10),
		chromedp.SetValue(`auth-captcha-guess`, str, chromedp.ByID),
		chromedp.CaptureScreenshot(&res5),
		chromedp.ScrollIntoView(`footer`),
		chromedp.Text(`h1`, &data, chromedp.NodeVisible, chromedp.ByQuery),
		chromedp.ActionFunc(func(ctx context.Context) error {
			node, err := dom.GetDocument().Do(ctx)
			if err != nil {
				return err
			}
			data, err := dom.GetOuterHTML().WithNodeID(node.NodeID).Do(ctx)
			if err != nil {
				return err
			}
			err = ioutil.WriteFile("result.html", []byte(data), os.ModePerm)
			if err != nil {
				return err
			}
			return nil
		}),
	)
	if err != nil {
		return err
	}
	os.WriteFile("./res4.png", res4, 0644)
	os.WriteFile("./res5.png", res5, 0644)
	return nil
}

/*
楽天用の処理
*/
func CheckingTheIntegrityOfRakutenInformation(userId string, password string) error {
	ctx, cancel := chromedp.NewContext(context.Background(), chromedp.WithBrowserOption())
	defer cancel()
	//var data string
	var res6 []byte
	var res7 []byte
	var res8 []byte
	var res9 []byte
	err := chromedp.Run(ctx,
		chromedp.Navigate("https://grp01.id.rakuten.co.jp/rms/nid/vc?__event=login&service_id=top"),
		chromedp.WaitReady("body"),
		chromedp.Sleep(time.Second*1),
		chromedp.CaptureScreenshot(&res6),
		chromedp.SetValue(`document.querySelector("#loginInner_u")`, userId, chromedp.ByJSPath),
		chromedp.SetValue(`document.querySelector("#loginInner_p")`, password, chromedp.ByJSPath),
		chromedp.CaptureScreenshot(&res7),
		chromedp.Click(`document.querySelector("#auto_logout")`, chromedp.ByJSPath),
		chromedp.CaptureScreenshot(&res8),
		chromedp.Click(`document.querySelector("#loginInner > p:nth-child(3) > input")`, chromedp.ByJSPath),
		chromedp.Sleep(time.Second*12),
		chromedp.CaptureScreenshot(&res9),
		chromedp.ActionFunc(func(ctx context.Context) error {
			node, err := dom.GetDocument().Do(ctx)
			if err != nil {
				return err
			}
			data, err := dom.GetOuterHTML().WithNodeID(node.NodeID).Do(ctx)
			if err != nil {
				return err
			}
			err = ioutil.WriteFile("../templates/autogen_rakuten.html", []byte(data), os.ModePerm)
			if err != nil {
				return err
			}
			return nil
		}),
	)
	if err != nil {
		return err
	}
	os.WriteFile("./res6.png", res6, 0644)
	os.WriteFile("./res7.png", res7, 0644)
	os.WriteFile("./res8.png", res8, 0644)
	os.WriteFile("./res9.png", res9, 0644)
	return nil
}

/*
楽天サイトに入るだけの処理
*/
func GoRakuten() error {
	ctx, cancel := chromedp.NewContext(context.Background(), chromedp.WithBrowserOption())
	defer cancel()
	err := chromedp.Run(ctx,
		chromedp.Navigate("https://rakuten.co.jp"),
		chromedp.WaitReady("body"),
		chromedp.Sleep(time.Second*1),
		chromedp.ActionFunc(func(ctx context.Context) error {
			node, err := dom.GetDocument().Do(ctx)
			if err != nil {
				return err
			}
			data, err := dom.GetOuterHTML().WithNodeID(node.NodeID).Do(ctx)
			if err != nil {
				return err
			}
			err = ioutil.WriteFile("autogen_rakuten.html", []byte(data), os.ModePerm)
			if err != nil {
				return err
			}
			return nil
		}),
	)
	if err != nil {
		return err
	}
	return nil
}
