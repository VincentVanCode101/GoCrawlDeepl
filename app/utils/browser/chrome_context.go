package browser

import (
	"context"
	"errors"

	"github.com/chromedp/chromedp"
)

func buildChromeOptions(headless, disableGPU, startMaximized bool) []chromedp.ExecAllocatorOption {
	options := []chromedp.ExecAllocatorOption{
		chromedp.Flag("headless", headless),
		chromedp.Flag("disable-gpu", disableGPU),
		chromedp.Flag("start-maximized", startMaximized),
		chromedp.NoSandbox,
	}
	return append(chromedp.DefaultExecAllocatorOptions[:], options...)
}

func getChromeExecAllocatorContext(opts []chromedp.ExecAllocatorOption) (context.Context, context.CancelFunc, error) {
	ctx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	if ctx == nil { // Typically you'd check an actual error, not just a nil context, but chromedp.NewExecAllocator() does not return an err
		return nil, nil, errors.New("failed to create Chrome executor allocator")
	}
	return ctx, cancel, nil
}

// GetChromeContext initializes and returns a Chrome browser context for web scraping.
func GetChromeContext() (context.Context, context.CancelFunc, context.CancelFunc, error) {
	opts := buildChromeOptions(true, true, true)
	execAllocCtx, cancelExecAllocator, err := getChromeExecAllocatorContext(opts)
	if err != nil {
		return nil, nil, nil, err
	}
	chromeCtx, cancelChrome := chromedp.NewContext(execAllocCtx)
	if chromeCtx == nil { // Again, you'd typically check an error, not context
		cancelExecAllocator()
		return nil, nil, nil, errors.New("failed to create Chrome context")
	}
	return chromeCtx, cancelChrome, cancelExecAllocator, nil
}
