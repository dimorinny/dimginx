package http

const (
	stringSeparator  = "\r\n"
	requestSeparator = "\r\n\r\n"
	dateTimeFormat   = "Mon, _2 Jan 2006 15:04:05 GMT"
	serverName       = "DimGINX"
	defaultFile      = "/index.html"
	notFoundFile     = "/404.html"
	httpVersion      = "1.1"
)

var exts = map[string]string{
	"txt":  "application/text",
	"html": "text/html",
	"json": "application/json",
	"jpg":  "image/jpeg",
	"jpeg": "image/jpeg",
	"png":  "image/png",
	"js":   "text/javascript",
	"css":  "text/css",
	"gif":  "image/gif",
	"swf":  "application/x-shockwave-flash",
}
