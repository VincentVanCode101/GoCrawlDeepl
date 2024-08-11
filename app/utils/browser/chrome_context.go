package browser

import (
	"context"

	"github.com/chromedp/chromedp"
)

func buildChromeOptions(headless, disableGPU, startMaximized bool) []chromedp.ExecAllocatorOption {
	options := []chromedp.ExecAllocatorOption{
		chromedp.Flag("headless", headless),
		chromedp.Flag("disable-gpu", disableGPU),
		chromedp.Flag("start-maximized", startMaximized),
	}
	return append(chromedp.DefaultExecAllocatorOptions[:], options...)
}

func getChromeExecAllocatorContext(opts []chromedp.ExecAllocatorOption) (context.Context, context.CancelFunc) {
	return chromedp.NewExecAllocator(context.Background(), opts...)
}

// GetChromeContext initializes and returns a Chrome browser context for web scraping.
func GetChromeContext() (context.Context, context.CancelFunc, context.CancelFunc) {
	opts := buildChromeOptions(true, true, true)
	execAllocCtx, cancelExecAllocator := getChromeExecAllocatorContext(opts)
	chromeCtx, cancelChrome := chromedp.NewContext(execAllocCtx)
	return chromeCtx, cancelChrome, cancelExecAllocator
}
