package service

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
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
	os.WriteFile("server/templates/img/res0.png", res0, 0644)
	os.WriteFile("server/templates/img/res1.png", res1, 0644)
	os.WriteFile("server/templates/img/res2.png", res2, 0644)
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
	os.WriteFile("server/templates/img/res4.png", res4, 0644)
	os.WriteFile("server/templates/img/res5.png", res5, 0644)
	return nil
}

/*
楽天用の処理
*/
func CheckingTheIntegrityOfRakutenInformation(userId string, password string) error {
	ctx, cancel := chromedp.NewContext(context.Background(), chromedp.WithBrowserOption())
	defer cancel()
	var res string
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
	os.WriteFile("server/templates/img/res6.png", res6, 0644)
	os.WriteFile("server/templates/img/res7.png", res7, 0644)
	os.WriteFile("server/templates/img/res8.png", res8, 0644)
	os.WriteFile("server/templates/img/res9.png", res9, 0644)
	output := `{{define "autogen_rakuten_info"}}` + res + `{{end}}`
	output = strings.Replace(output, `<div class="spacer--xFAdr  block--2PK_L  
        
        
        
        padding-bottom-small--sgTI2
       
      
      
      
      border-bottom-gray--OXbtM
      
     white--3LZcf"><div class="spacer--xFAdr  block--2PK_L  
        padding-left-xlarge--2d9GV
        padding-right-xlarge--LeKQw
        
        padding-bottom-xxsmall--14_zk
        "><div class="container--IckCk type-3--1HhJJ"><span class="clickable-span--15gHd"><div class="icon--2tjYQ icon-left--3FsRA"><div class="text-display--1Iony type-icon--3g0D- size-custom-medium--3iEUT align-left--1hi1x color-information-icon--3Z3gZ  layout-inline--1ajCj"><div class="icon--2sY_j common-info-filled--auWfJ"></div></div></div><div class="text--2TE80"><div class="text-display--1Iony type-body--1W5uC size-small--sv6IW align-left--1hi1x color-gray-darker--1SJFG line-height-medium--2-H3z layout-inline--1ajCj">ウクライナ 緊急支援募金のお知らせ</div></div><div class="icon--2tjYQ icon-right--2F1nI"><div class="text-display--1Iony type-icon--3g0D- size-custom-small--2Y-pv align-right--2ACTn color-gray--1TFBo  layout-inline--1ajCj"><div class="icon--2sY_j common-chevron-right--VZMgW"></div></div></div></span></div></div></div>`, ``, -1)

	output = strings.Replace(output, `<div irc="CommonHeaderMall" data-url="https://www.rakuten.co.jp" data-settings="[
      {
  &quot;tracker&quot;: {
    &quot;params&quot;: {
      &quot;accountId&quot;: 1,
      &quot;serviceId&quot;: 1,
      &quot;pageLayout&quot;: &quot;pc&quot;,
      &quot;pageType&quot;: &quot;top&quot;
    }
  },

  &quot;showSearchBar&quot;: true,
  &quot;showMemberInfoSummary&quot;: false,
  &quot;showSpu&quot;: false,
  &quot;showCartModal&quot;: true,
  &quot;customLogoImageUrl&quot;: &quot;https://r.r10s.jp/com/img/thumb/logo/logo_rakuten_25th.svg&quot;,
      &quot;links&quot;: {
        &quot;top&quot;: &quot;https://corp.rakuten.co.jp/event/anniversary25th/?scid=wi_ich_r25_pc_top_header_25th_logo_v1&quot;
      },
  &quot;withBorder&quot;: false,
  &quot;suggestionUrl&quot; : &quot;https://rdc-api-catalog-gateway-api.rakuten.co.jp/SUI/autocomplete/pc&quot;,
  &quot;useTBasketDomain&quot;: true,
  &quot;api&quot;: {
      &quot;cartApiSid&quot;: 1010,
      &quot;notificationLocId&quot;: 25,
      &quot;url&quot;:&quot;https://api-ichiba-gateway.rakuten.co.jp/graphql-common-bff/graphql&quot;,
      &quot;apikey&quot;: &quot;59093b15781a092c9573ea7032016ddb&quot;,
      &quot;clientId&quot;: &quot;top&quot;,
      &quot;spuViewType&quot;: &quot;top&quot;,
      &quot;spuSource&quot;: &quot;pc&quot;,
      &quot;spuEncoding&quot;: &quot;UTF-8&quot;,
      &quot;spuAcc&quot;: 1,
      &quot;spuAid&quot;: 1
    }
}

    ]">`, ``, -1)
	err = os.WriteFile("server/templates/autogen_rakuten_login.html", []byte(output), 0644)
	if err != nil {
		return err
	}
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
