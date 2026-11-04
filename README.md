# cookie-conv-c2n

A tool that converts cookies copied from Chrome into Netscape cookie file format.

This is just a temp tool and might have some bugs. Do **NOT** use it in important projects which require stability.

## Usage

1. Use `go build` to compile the code into binary executable file.
2. Create a new text file.
3. Press `F12` in Chrome window, choose `Application` > `Cookies` > (domain name).
4. Select (highlight) all cookies and press `Ctrl+C` to copy them.
5. Paste cookies into the new empty text file and save it.
6. Drag the text file onto the executable file.
