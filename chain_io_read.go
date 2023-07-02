package main

import (
	"compress/gzip"
	"encoding/base64"
	"io"
	"os"
	"strings"
)

func main() {
	var r io.Reader
	r = strings.NewReader(data)
	r = base64.NewDecoder(base64.StdEncoding, r)
	r, _ = gzip.NewReader(r)
	io.Copy(os.Stdout, r)
}

const data = `H4sIAAAAAAAA/1SOMW6FMBBEe59iujSIC6RKlSOkXsKAVzHryLtw/i8+uKAdzXt6340SWArFkeUgikYU3ksC9v9qiExtmOT3zxEVkwbxwW1MXzY/ictBd7ZrGRIgNsMrZIbaoqax32hkYj0LOh6Zm7Mc9AFqiL3ZKXhbr2Prz4q1otpn+slaeJJOyCpqz7uHltIrnmMPqza+AgAA//8iRFXyCAEA`
