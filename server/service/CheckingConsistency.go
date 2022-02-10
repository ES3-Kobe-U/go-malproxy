package service

import (
	"context"
	"fmt"
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
	var res0 []byte
	var res1 []byte
	var res2 []byte
	var res3 []byte
	// var data string
	// var str string = "xxxxxx"
	ctx, cancel := chromedp.NewContext(context.Background(), chromedp.WithBrowserOption())
	defer cancel()

	task1 := chromedp.Tasks{ //タスクリストの作成
		chromedp.Navigate("https://www.amazon.co.jp/stores/page/FD12D065-F94A-4656-B816-4C8121739FAE?ingress=3"),
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
	}

	err := chromedp.Run(ctx, task1)
	if err != nil {
		return err
	}

	task2 := chromedp.Tasks{
		chromedp.CaptureScreenshot(&res3),
	}

	err = chromedp.Run(ctx, task2) //セッションが維持されているかの確認
	if err != nil {
		return err
	}

	// err := chromedp.Run(ctx,
	// 	chromedp.Navigate("https://www.amazon.co.jp/stores/page/FD12D065-F94A-4656-B816-4C8121739FAE?ingress=3"),
	// 	chromedp.WaitReady("body"),
	// 	chromedp.Click(`a[data-nav-role='signin']`, chromedp.ByQuery),
	// 	chromedp.Sleep(time.Second*2),
	// 	chromedp.SetValue(`ap_email`, email, chromedp.ByID),
	// 	chromedp.Click(`continue`, chromedp.ByID),
	// 	chromedp.Sleep(time.Second*1),
	// 	chromedp.SetValue(`ap_password`, password, chromedp.ByID),
	// 	chromedp.Click(`signInSubmit`, chromedp.ByID),
	// 	chromedp.Sleep(time.Second*2),
	// 	chromedp.CaptureScreenshot(&res1),
	// 	chromedp.SetValue(`ap_password`, password, chromedp.ByID),
	// 	chromedp.ActionFunc(func(c context.Context) error {
	// 		fmt.Println("output --> res1.png")
	// 		err := os.WriteFile("./res1.png", res1, 0644)
	// 		if err != nil {
	// 			return err
	// 		}
	// 		return nil
	// 	}),
	// 	chromedp.Sleep(time.Second*10),
	// 	chromedp.ActionFunc(func(c context.Context) error {
	// 		var input string
	// 		fmt.Print("fmt.Scan >")
	// 		fmt.Scanf("%s", &input)
	// 		str = input
	// 		fmt.Println("input:", str)
	// 		return nil
	// 	}),
	// 	chromedp.Sleep(time.Second*10),
	// 	chromedp.SetValue(`auth-captcha-guess`, str, chromedp.ByID),
	// 	chromedp.CaptureScreenshot(&res2),
	// 	chromedp.ScrollIntoView(`footer`),
	// 	chromedp.Text(`h1`, &data, chromedp.NodeVisible, chromedp.ByQuery),
	// 	chromedp.ActionFunc(func(ctx context.Context) error {
	// 		node, err := dom.GetDocument().Do(ctx)
	// 		if err != nil {
	// 			return err
	// 		}
	// 		data, err := dom.GetOuterHTML().WithNodeID(node.NodeID).Do(ctx)
	// 		if err != nil {
	// 			return err
	// 		}
	// 		err = ioutil.WriteFile("result.html", []byte(data), os.ModePerm)
	// 		if err != nil {
	// 			return err
	// 		}
	// 		return nil
	// 	}),
	// )
	// if err != nil {
	// 	return err
	// }
	os.WriteFile("./res0.png", res0, 0644)
	os.WriteFile("./res1.png", res1, 0644)
	os.WriteFile("./res2.png", res2, 0644)
	os.WriteFile("./res3.png", res3, 0644)
	return nil
}

/*
楽天用の処理
*/
func CheckingTheIntegrityOfRakutenInformation(userId string, password string) error {
	ctx, cancel := chromedp.NewContext(context.Background(), chromedp.WithBrowserOption())
	defer cancel()
	//var data string
	err := chromedp.Run(ctx,
		chromedp.Navigate("https://grp01.id.rakuten.co.jp/rms/nid/vc?__event=login&service_id=top"),
		chromedp.WaitReady("body"),
		chromedp.Sleep(time.Second*1),
		chromedp.SetValue(`loginInner_u`, userId, chromedp.ByID),
		chromedp.SetValue(`loginInner_p`, password, chromedp.ByID),
		chromedp.Click(`auto_logout`, chromedp.ByID),
		chromedp.Sleep(time.Second*1),
		chromedp.Click(`#loginInner > p:nth-child(3) > input`, chromedp.BySearch),
		chromedp.Sleep(time.Second*10),
		chromedp.ActionFunc(func(ctx context.Context) error {
			node, err := dom.GetDocument().Do(ctx)
			if err != nil {
				return err
			}
			data, er := dom.GetOuterHTML().WithNodeID(node.NodeID).Do(ctx)
			fmt.Print(data)
			return er
		}),
	)
	if err != nil {
		return err
	}
	return nil
}
