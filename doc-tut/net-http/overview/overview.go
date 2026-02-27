package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

/*
package http provided HTTP clienta and server implementations
Get, Head, Post, PostForm make HTTP (or HTTPS) requests
*/

func main() {
	resp, err := http.Get("http://example.com")
	if err != nil {
		// handle error
		fmt.Printf("%v\n", err)
	}
	// caller must close the response body when finished with it
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	fmt.Println(string(body))

	/*
	   <!doctype html>
	   <html>
	   <head>
	       <title>Example Domain</title>

	       <meta charset="utf-8" />
	       <meta http-equiv="Content-type" content="text/html; charset=utf-8" />
	       <meta name="viewport" content="width=device-width, initial-scale=1" />
	       <style type="text/css">
	       body {
	           background-color: #f0f0f2;
	           margin: 0;
	           padding: 0;
	           font-family: -apple-system, system-ui, BlinkMacSystemFont, "Segoe UI", "Open Sans", "Helvetica Neue", Helvetica, Arial, sans-serif;

	       }
	       div {
	           width: 600px;
	           margin: 5em auto;
	           padding: 2em;
	           background-color: #fdfdff;
	           border-radius: 0.5em;
	           box-shadow: 2px 3px 7px 2px rgba(0,0,0,0.02);
	       }
	       a:link, a:visited {
	           color: #38488f;
	           text-decoration: none;
	       }
	       @media (max-width: 700px) {
	           div {
	               margin: 0 auto;
	               width: auto;
	           }
	       }
	       </style>
	   </head>

	   <body>
	   <div>
	       <h1>Example Domain</h1>
	       <p>This domain is for use in illustrative examples in documents. You may use this
	       domain in literature without prior coordination or asking for permission.</p>
	       <p><a href="https://www.iana.org/domains/example">More information...</a></p>
	   </div>
	   </body>
	   </html>
	*/
}

func post() {
	resp, err := http.Post("http://example.com/upload", "image/png", nil)
	_ = err
	_ = resp
}

func postForm() {
	resp, err := http.PostForm("http://example.com/form",
		url.Values{"key": {"value"}, "id": {"123"}},
	)
	_ = err
	_ = resp
}
