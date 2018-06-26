// Code generated by hero.
// source: /c/Users/ecoip/go/src/market/template/customerlist.html
// DO NOT EDIT!
package template

import (
	"io"

	"github.com/shiyanhui/hero"
)

func CustomerList(customerList []string, w io.Writer) (int, error) {
	_buffer := hero.GetBuffer()
	defer hero.PutBuffer(_buffer)
	_buffer.WriteString(`<!DOCTYPE html>
	<html>
    	<head>
        	<meta charset="utf-8">
    	</head>
    	<body>
        `)
	for _, customer := range customerList {
		_buffer.WriteString(`
        <ul>
            `)
		_buffer.WriteString(`<li>`)
		hero.EscapeHTML(customer, _buffer)
		_buffer.WriteString(`</li>`)
		_buffer.WriteString(`
        </ul>
    `)
	}

	_buffer.WriteString(`
    </body>
</html>`)
	return w.Write(_buffer.Bytes())

}