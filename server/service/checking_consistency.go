package service

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"github.com/go-malproxy/server/params"
	"github.com/sirupsen/logrus"
)

type Contents struct {
	CTX *context.Context
}

/*
Amazon用の処理
*/
func (c *Contents) CheckingTheIntegrityOfAmazonInformation(email string, password string) error {
	var res string
	var picture []byte
	// var cancel context.CancelFunc
	// *c.CTX, cancel = chromedp.NewContext(context.Background(), chromedp.WithBrowserOption())
	// defer cancel()
	*c.CTX, _ = chromedp.NewContext(context.Background(), chromedp.WithDebugf(logrus.Printf))
	task1 := chromedp.Tasks{ //タスクリストの作成
		network.Enable(),
		network.SetExtraHTTPHeaders(network.Headers(map[string]interface{}{
			"User-Agent": "Mozilla",
		})),
		chromedp.Navigate("https://www.amazon.co.jp/?&tag=hydraamazonav-22&ref=pd_sl_2ykkalld4i_e&adgrpid=54841807378&hvpone=&hvptwo=&hvadid=289239574720&hvpos=&hvnetw=g&hvrand=13443261105670128409&hvqmt=e&hvdev=c&hvdvcmdl=&hvlocint=&hvlocphy=1009565&hvtargid=kwd-333588672930&hydadcr=15460_10908920&gclid=Cj0KCQiAjJOQBhCkARIsAEKMtO0jVYrB9RxzmNKNeCPzZE0CB_TUL10D5UonY9FkHd4maUGPDrYDe4UaAnhwEALw_wcB"),
		chromedp.WaitReady("body"),
		chromedp.Click(`a[data-nav-role='signin']`, chromedp.ByQuery),
		chromedp.Sleep(time.Second * 1),
		chromedp.SetValue(`ap_email`, email, chromedp.ByID),
		chromedp.Sleep(time.Second * 1),
		chromedp.Click(`continue`, chromedp.ByID),
		chromedp.Sleep(time.Second * 1),
		chromedp.SetValue(`ap_password`, password, chromedp.ByID),
		chromedp.Sleep(time.Second * 1),
		chromedp.Click(`signInSubmit`, chromedp.ByID),
		chromedp.Sleep(time.Second * 2),
		chromedp.CaptureScreenshot(&picture),
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
	err := chromedp.Run(*c.CTX, task1)
	if err != nil {
		return err
	}
	err = os.WriteFile("server/templates/img/res2.png", picture, 0644)
	//err = os.WriteFile("../../server/templates/img/picture.png", picture, 0644)
	if err != nil {
		return err
	}
	output := `{{define "autogen_amazon_info"}}` + res + `{{end}}`
	output = strings.Replace(output, `<a href="`, `<a href="/template?url=`, -1) //文字列の置き換え
	output = strings.Replace(output, `<a href='`, `<a href='/template?url=`, -1) //文字列の置き換え
	output = strings.Replace(output, params.AmazonCaptchaParamNo1, params.AmazonCaptchaParamNo1Replace, -1)
	output = strings.Replace(output, params.AmazonCaptchaParamNo2, params.AmazonCaptchaParamNo2Replace, -1)
	err = os.WriteFile("server/templates/autogen_amazon_login.html", []byte(output), 0644)
	//err = os.WriteFile("../../server/templates/autogen_amazon_login.html", []byte(output), 0644)
	if err != nil {
		return err
	}
	return nil
}

func (c *Contents) CheckingTheIntegrityOfAmazonCaptcha(password string, guess string) error {
	var res string
	var picture []byte
	if err := chromedp.Run(*c.CTX); err != nil {
		return err
	}
	ctx1, _ := context.WithTimeout(*c.CTX, 30*time.Second)
	ctx1, _ = chromedp.NewContext(ctx1)
	taskListNo1 := chromedp.Tasks{
		chromedp.SetValue(`document.querySelector("#ap_password")`, password, chromedp.ByJSPath),
		chromedp.Sleep(time.Second * 1),
		chromedp.SetValue(`document.querySelector("#auth-captcha-guess")`, guess, chromedp.ByJSPath),
		chromedp.Sleep(time.Second * 1),
		chromedp.CaptureScreenshot(&picture),
		chromedp.Click(`signInSubmit`, chromedp.ByID),
		chromedp.Sleep(time.Second * 2),
		chromedp.ActionFunc(func(ctx context.Context) error {
			node, err := dom.GetDocument().Do(ctx)
			if err != nil {
				fmt.Println("node err")
				return err
			}
			res, err = dom.GetOuterHTML().WithNodeID(node.NodeID).Do(ctx)
			if err != nil {
				fmt.Println("res err")
				return err
			}
			return nil
		}),
	}
	if err := chromedp.Run(ctx1, taskListNo1); err != nil {
		fmt.Println("Run err")
		return err
	}
	if err := os.WriteFile("server/templates/img/captchaNo1.png", picture, 0644); err != nil {
		return err
	}
	output := `{{define "autogen_amazon_captcha"}}` + res + `{{end}}`
	output = strings.Replace(output, `<a href="`, `<a href="/template?url=`, -1) //文字列の置き換え
	output = strings.Replace(output, `<a href='`, `<a href='/template?url=`, -1) //文字列の置き換え
	output = strings.Replace(output, params.AmazonCaptchaParamNo1, params.AmazonCaptchaParamNo1Replace, -1)
	output = strings.Replace(output, params.AmazonCaptchaParamNo2, params.AmazonCaptchaParamNo2Replace, -1)
	if err := os.WriteFile("server/templates/autogen_amazon_captcha.html", []byte(output), 0644); err != nil {
		return err
	}
	return nil
}

/*
楽天用の処理
*/
func (c *Contents) CheckingTheIntegrityOfRakutenInformation(userId string, password string) error {
	//var cancel context.CancelFunc
	*c.CTX, _ = chromedp.NewContext(context.Background(), chromedp.WithBrowserOption())
	//defer cancel()
	var res string
	err := chromedp.Run(*c.CTX,
		chromedp.Navigate("https://grp01.id.rakuten.co.jp/rms/nid/vc?__event=login&service_id=top"),
		chromedp.WaitReady("body"),
		chromedp.SetValue(`document.querySelector("#loginInner_u")`, userId, chromedp.ByJSPath),
		chromedp.SetValue(`document.querySelector("#loginInner_p")`, password, chromedp.ByJSPath),
		chromedp.Click(`document.querySelector("#auto_logout")`, chromedp.ByJSPath),
		chromedp.Click(`document.querySelector("#loginInner > p:nth-child(3) > input")`, chromedp.ByJSPath),
		chromedp.Sleep(time.Second*20),
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
	)
	if err != nil {
		return err
	}
	output := `{{define "autogen_rakuten_info"}}` + res + `{{end}}`
	output = strings.Replace(output, params.RakutenReplaceNo1, ``, -1)
	output = strings.Replace(output, params.RakutenReplaceNo2, ``, -1)
	output = strings.Replace(output, `<a href="`, `<a href="/template?url=`, -1) //文字列の置き換え
	output = strings.Replace(output, `<a href='`, `<a href='/template?url=`, -1) //文字列の置き換え
	output = strings.Replace(output, `pa3.min.js`, ``, -1)                       //楽天のCORSを回避する為に削除
	output = strings.Replace(output, params.RakutenLoginCode, params.ReplaceRakutenLoginCode, -1)
	err = os.WriteFile("server/templates/autogen_rakuten_login.html", []byte(output), 0644)
	//err = os.WriteFile("../../server/templates/autogen_rakuten_login.html", []byte(output), 0644)
	if err != nil {
		return err
	}
	return nil
}
