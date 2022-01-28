package myquote

import "rsc.io/quote"

func TestQuote() string {
	return quote.Hello() + "\n" + quote.Go() + "\n" + quote.Opt() + "\n" + quote.Glass()
}
