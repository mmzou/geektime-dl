package utils

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/cdproto/runtime"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/device"
)

//ColumnPrintToPDF print pdf
func ColumnPrintToPDF(aid int, filename string, cookies map[string]string) error {
	var buf []byte
	// create chrome instance
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	// create a timeout
	ctx, cancel = context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	err := chromedp.Run(ctx,
		chromedp.Tasks{
			chromedp.Emulate(device.IPhone7),
			enableLifeCycleEvents(),
			setCookies(cookies),
			navigateAndWaitFor(`https://time.geekbang.org/column/article/`+strconv.Itoa(aid), "networkIdle"),
			chromedp.ActionFunc(func(ctx context.Context) error {
				time.Sleep(time.Second * 5)
				s := `
					var iconfontDivs = document.querySelectorAll('.iconfont')
					iconfontDivs[0].parentElement.parentElement.style.display='none';
					iconfontDivs[5].parentElement.parentElement.style.display='none';
					iconfontDivs[6].parentElement.parentElement.style.display='none';
					iconfontDivs[7].style.display='none';
					var as = document.getElementsByTagName('a');
					for (var i = 0; i < as.length; ++i){
						if(as[i].innerText === "提建议"){
							as[i].parentNode.parentNode.style.display="none";
							break;
						}
					}
					var bottom = document.querySelector('.bottom-wrapper');
					if(bottom){
						bottom.parentElement.style.display='none'
					}
					[...document.querySelectorAll('ul>li>div>div>div:nth-child(2)>span')].map(e=>e.click());
				`
				_, exp, err := runtime.Evaluate(s).Do(ctx)
				if err != nil {
					return err
				}

				if exp != nil {
					return exp
				}

				return nil
			}),
			/*chromedp.ActionFunc(func(ctx context.Context) error {
				s := `
					var divs = document.getElementsByTagName('div');
					for (var i = 0; i < divs.length; ++i){
						if(divs[i].innerText === "打开APP"){
							divs[i].parentNode.parentNode.style.display="none";
							break;
						}
					}
				`
				_, exp, err := runtime.Evaluate(s).Do(ctx)
				if err != nil {
					return err
				}

				if exp != nil {
					return exp
				}

				return nil
			}),*/

			chromedp.ActionFunc(func(ctx context.Context) error {
				time.Sleep(time.Second * 10)
				var err error
				buf, _, err = page.PrintToPDF().WithPrintBackground(true).Do(ctx)
				return err
			}),
		},
	)

	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, buf, 0644)
}

func setCookies(cookies map[string]string) chromedp.ActionFunc {
	return chromedp.ActionFunc(func(ctx context.Context) error {
		expr := cdp.TimeSinceEpoch(time.Now().Add(180 * 24 * time.Hour))

		for key, value := range cookies {
			success, err := network.SetCookie(key, value).WithExpires(&expr).WithDomain(".geekbang.org").WithHTTPOnly(true).Do(ctx)
			if err != nil {
				return err
			}

			if !success {
				return fmt.Errorf("could not set cookie %q to %q", key, value)
			}
		}
		return nil
	})
}

func enableLifeCycleEvents() chromedp.ActionFunc {
	return func(ctx context.Context) error {
		err := page.Enable().Do(ctx)
		if err != nil {
			return err
		}

		return page.SetLifecycleEventsEnabled(true).Do(ctx)
	}
}

func navigateAndWaitFor(url string, eventName string) chromedp.ActionFunc {
	return func(ctx context.Context) error {
		_, _, _, err := page.Navigate(url).Do(ctx)
		if err != nil {
			return err
		}

		return waitFor(ctx, eventName)
	}
}

// waitFor blocks until eventName is received.
// Examples of events you can wait for:
//     init, DOMContentLoaded, firstPaint,
//     firstContentfulPaint, firstImagePaint,
//     firstMeaningfulPaintCandidate,
//     load, networkAlmostIdle, firstMeaningfulPaint, networkIdle
//
// This is not super reliable, I've already found incidental cases where
// networkIdle was sent before load. It's probably smart to see how
// puppeteer implements this exactly.
func waitFor(ctx context.Context, eventName string) error {
	ch := make(chan struct{})
	cctx, cancel := context.WithCancel(ctx)
	chromedp.ListenTarget(cctx, func(ev interface{}) {
		switch e := ev.(type) {
		case *page.EventLifecycleEvent:
			if e.Name == eventName {
				cancel()
				close(ch)
			}
		}
	})

	select {
	case <-ch:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
