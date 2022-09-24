package bilipush

import (
	"errors"
	"github.com/avast/retry-go"
	"github.com/playwright-community/playwright-go"
	log "github.com/sirupsen/logrus"
)

func GetPage(Browser playwright.Browser) (playwright.Page, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Error("playwright 页面获取失败 ", err)
		}
	}()
	page, err := Browser.NewPage(
		playwright.BrowserNewContextOptions{
			IsMobile:  playwright.Bool(true),
			UserAgent: playwright.String("Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1 Edg/105.0.0.0"),
			Screen: &playwright.BrowserNewContextOptionsScreen{
				Width:  playwright.Int(414),
				Height: playwright.Int(896),
			},
			Viewport: &playwright.BrowserNewContextOptionsViewport{
				Width:  playwright.Int(414),
				Height: playwright.Int(896),
			},
		})
	if err != nil {
		return nil, errors.New("could not create page: %v")
	}
	return page, nil
}

func GetScreenshot(url string) ([]byte, error) {
	var screenshot []byte
	err := retry.Do(func() error {
		pw, err := playwright.Run()
		if err != nil {
			return err
		}
		defer pw.Stop()
		browser, err := pw.Chromium.Launch()
		if err != nil {
			return err
		}
		defer browser.Close()
		page, err := GetPage(browser)
		if err != nil {
			return err
		}
		defer page.Close()
		_, err = page.Goto(url, playwright.PageGotoOptions{
			Timeout:   playwright.Float(10000),
			WaitUntil: playwright.WaitUntilStateNetworkidle,
		})
		if err != nil {
			return err
		}
		if page.URL() == "https://m.bilibili.com/404" {
			return errors.New("404")
		}
		card, err := page.QuerySelector(".dyn-card")
		if card == nil || err != nil {
			return errors.New("动态元素没找到")
		}
		screenshot, err = card.Screenshot(
			playwright.ElementHandleScreenshotOptions{
				Timeout: playwright.Float(10000),
			})
		if len(screenshot) == 0 || err != nil {
			return errors.New("截图失败")
		}
		return nil
	}, retry.Attempts(3), retry.RetryIf(func(err error) bool {
		return err.Error() != "404"
	}), retry.LastErrorOnly(true))
	if err != nil {
		if err.Error() == "404" {
			return nil, errors.New("动态进入审核")
		} else {
			return nil, err
		}
	}
	return screenshot, nil
}
