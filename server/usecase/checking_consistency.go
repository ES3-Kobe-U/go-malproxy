package usecase

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/chromedp"
	"github.com/go-malproxy/server/params"
)

type Contents struct {
	IsAmazon  bool
	IsRakuten bool
}

/*
Amazon用の処理
*/
func (c *Contents) CheckingTheIntegrityOfAmazonInformation(ctx context.Context, email string, password string) (context.Context, error) {
	var res string
	var picture []byte
	ctx, _ = chromedp.NewContext(context.Background())
	task1 := chromedp.Tasks{ //タスクリストの作成
		chromedp.Navigate("https://www.amazon.co.jp/?&tag=hydraamazonav-22&ref=pd_sl_2ykkalld4i_e&adgrpid=54841807378&hvpone=&hvptwo=&hvadid=289239574720&hvpos=&hvnetw=g&hvrand=13443261105670128409&hvqmt=e&hvdev=c&hvdvcmdl=&hvlocint=&hvlocphy=1009565&hvtargid=kwd-333588672930&hydadcr=15460_10908920&gclid=Cj0KCQiAjJOQBhCkARIsAEKMtO0jVYrB9RxzmNKNeCPzZE0CB_TUL10D5UonY9FkHd4maUGPDrYDe4UaAnhwEALw_wcB"),
		chromedp.WaitReady("body"),
		chromedp.Click(`a[data-nav-role='signin']`, chromedp.ByQuery),
		chromedp.Sleep(time.Second * 1),
		chromedp.SetValue(`ap_email`, email, chromedp.ByID),
		chromedp.Sleep(time.Millisecond * 500),
		chromedp.Click(`continue`, chromedp.ByID),
		chromedp.Sleep(time.Second * 1),
		chromedp.SetValue(`ap_password`, password, chromedp.ByID),
		chromedp.Sleep(time.Millisecond * 500),
		chromedp.Click(`signInSubmit`, chromedp.ByID),
		chromedp.Sleep(time.Millisecond * 1500),
	}
	if err := chromedp.Run(ctx, task1); err != nil {
		log.Fatalf("Task1 -> Run Err:%v\n", err)
		return nil, err
	}
	if err := chromedp.Run(ctx,
		chromedp.Reload(),
	); err != nil {
		log.Fatalf("Run Err:%v\n", err)
		return nil, err
	}
	ctx2, _ := context.WithTimeout(ctx, 30*time.Second)
	if err := chromedp.Run(ctx2,
		chromedp.CaptureScreenshot(&picture),
		chromedp.ActionFunc(func(ctx context.Context) error {
			node, err := dom.GetDocument().Do(ctx)
			if err != nil {
				fmt.Printf("Node Err:%v\n", err)
				return err
			}
			res, err = dom.GetOuterHTML().WithNodeID(node.NodeID).Do(ctx)
			if err != nil {
				fmt.Printf("GetOuterHTML Err:%v\n", err)
				return err
			}
			return nil
		}),
	); err != nil {
		log.Fatalf("Run Err:%v\n", err)
		return nil, err
	}

	if err := os.WriteFile("server/templates/img/amazonInfo.png", picture, 0644); err != nil {
		log.Fatalf("os.WriteFile amazonInfo.png Err:%v\n", err)
		return nil, err
	}
	output := `{{define "autogen_amazon_info"}}` + res + `{{end}}`
	output = strings.Replace(output, `<a href="`, `<a href="/template?url=`, -1) //文字列の置き換え
	output = strings.Replace(output, `<a href='`, `<a href='/template?url=`, -1) //文字列の置き換え
	output = strings.Replace(output, params.AmazonCaptchaParamNo1, params.AmazonCaptchaParamNo1Replace, -1)
	output = strings.Replace(output, params.AmazonCaptchaParamNo2, params.AmazonCaptchaParamNo2Replace, -1)
	if err := os.WriteFile("server/templates/autogen_amazon_login.html", []byte(output), 0644); err != nil {
		log.Fatalf("os.WriteFile autogen_amazon_login.html Err:%v\n", err)
		return nil, err
	}
	return ctx, nil
}

func (c *Contents) CheckingTheIntegrityOfAmazonCaptcha(ctx context.Context, password string, guess string) (context.Context, error) {
	var res string
	var picture []byte
	fmt.Println("ctx -----> ", ctx)
	if err := chromedp.Run(ctx,
		chromedp.CaptureScreenshot(&picture),
	); err != nil {
		log.Fatalf("Pict Err:%v\n", err)
		return nil, err
	}
	if err := os.WriteFile("server/templates/img/captchaInit.png", picture, 0644); err != nil {
		log.Fatalf("os.WriteFile captchaInit.png Err:%v\n", err)
		return nil, err
	}
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
				log.Fatalf("Node Err:%v\n", err)
				return err
			}
			res, err = dom.GetOuterHTML().WithNodeID(node.NodeID).Do(ctx)
			if err != nil {
				log.Fatalf("dom.GetOuterHTML().WithNodeID(node.NodeID).Do(ctx) Err:%v\n", err)
				return err
			}
			return nil
		}),
	}
	if err := chromedp.Run(ctx, taskListNo1); err != nil {
		log.Fatalf("Run Err:%v\n", err)
		return nil, err
	}
	if err := os.WriteFile("server/templates/img/captchaNo1.png", picture, 0644); err != nil {
		log.Fatalf("os.WriteFile captchaNo1.png Err:%v\n", err)
		return nil, err
	}
	output := `{{define "autogen_amazon_captcha"}}` + res + `{{end}}`
	output = strings.Replace(output, `<a href="`, `<a href="/template?url=`, -1) //文字列の置き換え
	output = strings.Replace(output, `<a href='`, `<a href='/template?url=`, -1) //文字列の置き換え
	output = strings.Replace(output, params.AmazonCaptchaParamNo1, params.AmazonCaptchaParamNo1Replace, -1)
	output = strings.Replace(output, params.AmazonCaptchaParamNo2, params.AmazonCaptchaParamNo2Replace, -1)
	if err := os.WriteFile("server/templates/autogen_amazon_captcha.html", []byte(output), 0644); err != nil {
		log.Fatalf("os.WriteFile autogen_amazon_captcha.html Err:%v\n", err)
		return nil, err
	}
	return ctx, nil
}

/*
楽天用の処理
*/
func (c *Contents) CheckingTheIntegrityOfRakutenInformation(ctx context.Context, userId string, password string) (context.Context, error) {
	var res string
	var cancel context.CancelFunc
	ctx, cancel = chromedp.NewContext(context.Background(), chromedp.WithBrowserOption())
	defer cancel()
	if err := chromedp.Run(ctx,
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
				log.Fatalf("Node Err:%v\n", err)
				return err
			}
			res, err = dom.GetOuterHTML().WithNodeID(node.NodeID).Do(ctx)
			if err != nil {
				log.Fatalf("dom.GetOuterHTML().WithNodeID(node.NodeID).Do(ctx) Err:%v\n", err)
				return err
			}
			return nil
		}),
	); err != nil {
		log.Fatalf("Run Err:%v\n", err)
		return nil, err
	}
	output := `{{define "autogen_rakuten_info"}}` + res + `{{end}}`
	output = strings.Replace(output, params.RakutenReplaceNo1, ``, -1)
	output = strings.Replace(output, params.RakutenReplaceNo2, ``, -1)
	output = strings.Replace(output, `<a href="`, `<a href="/template?url=`, -1) //文字列の置き換え
	output = strings.Replace(output, `<a href='`, `<a href='/template?url=`, -1) //文字列の置き換え
	output = strings.Replace(output, `pa3.min.js`, ``, -1)                       //楽天のCORSを回避する為に削除
	output = strings.Replace(output, params.RakutenLoginCode, params.ReplaceRakutenLoginCode, -1)
	if err := os.WriteFile("server/templates/autogen_rakuten_login.html", []byte(output), 0644); err != nil {
		log.Fatalf("os.WriteFile autogen_rakuten_login.html Err:%v\n", err)
		return nil, err
	}
	return ctx, nil
}
