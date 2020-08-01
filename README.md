# html-selector-checker

CLI tool of html selector checker implementation with Go.



## Install

```bash
go get github.com/shinshin86/html-selector-checker
```



## Usage

```bash
# Target HTML File -> target-file.html (default: index.html)
# Selector -> class="listItem" (default: body)
# Is Debug Mode (default: false)
go run main.go -f="target-file.html" -s=".listItem" -d=true
```

