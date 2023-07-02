package main

import (
	"compress/gzip"
	"encoding/base64"
	"io"
	"os"
)

func main() {
	var w io.Writer = os.Stdout
	w = base64.NewEncoder(base64.StdEncoding, w)
	gw := gzip.NewWriter(w)
	gw.Write([]byte(data2))
	gw.Close()
}

const data2 = `Great fleas have little fleas 
  upon their backs to bite 'em.
And little fleas have lesser fleas,
  and so ad infinitum.
And the greate fleas themselves, in turn,
  have greater fleas to go on;
While these again have greater still,
  and greater still, and so on.`
