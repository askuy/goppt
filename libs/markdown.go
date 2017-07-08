package libs

import (
	"io"
	"os"
	//"io/ioutil"
)

// renderDoc reads the present file, gets its template representation,
// and executes the template, sending output to w.
func renderMarkdown(w io.Writer, docFile string) error {

	f, err := os.Open(docFile)
	if err != nil {
		return nil
	}
	defer f.Close()
	//fd,err := ioutil.ReadAll(f)
	//io.WriteString(w,string(fd))
	markdownTemplate.Execute(w, nil)

	return nil
}