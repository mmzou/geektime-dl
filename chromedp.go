package main

import (
	"context"
	"io/ioutil"
	"log"
	"time"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/cdproto/runtime"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/device"
)

func main() {

	var buf []byte

	// create chrome instance
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	// create a timeout
	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	// navigate to a page, wait for an element, click
	var example string

	chromedp.Flag("disable-images", true)
	err := chromedp.Run(ctx,

		chromedp.Tasks{
			chromedp.Emulate(device.IPhone7),
			enableLifeCycleEvents(),
			navigateAndWaitFor(`https://time.geekbang.org/column/article/225604`, "networkIdle"),
			chromedp.ActionFunc(func(ctx context.Context) error {
				s := `
					document.querySelector('.iconfont').parentElement.parentElement.style.display='none';
					document.querySelector('.bottom-wrapper').parentElement.style.display='none';
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
			chromedp.ActionFunc(func(ctx context.Context) error {
				time.Sleep(time.Second * 5)
				var err error
				buf, _, err = page.PrintToPDF().WithPrintBackground(true).Do(ctx)
				return err
			}),
		},

		// chromedp.Emulate(device.IPhone7),
		//访问打开必应页面
		// chromedp.Navigate(`https://time.geekbang.org/column/article/225604`),
		// chromedp.Text(`#app > div._3ADRghFH_0 > div > div._50pDbNcP_0._3O_7qs2p_0._2q1SuvsS_0 > div._50pDbNcP_0 > h1`, &example),
		// chromedp.ActionFunc(func(ctx context.Context) error {
		// 	time.Sleep(time.Second * 5)
		// 	var err error
		// 	buf, _, err = page.PrintToPDF().WithPrintBackground(true).Do(ctx)
		// 	return err
		// }),
	)
	if err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile("4.pdf", buf, 0644); err != nil {
		log.Fatal(err)
	}
	log.Printf("example: %s", example)
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
