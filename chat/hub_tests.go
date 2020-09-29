// Author: yangzq80@gmail.com
// Date: 2020-09-28
//
package chat

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"testing"
	"time"
)

func TestTime(t *testing.T)  {
	log.Println("222")
	fmt.Print(time.Now().Format("2006-01-02 15:04:05"))
}


func TestStr(t *testing.T)  {
	r := strings.NewReader("some io.Reader stream to be read\n")

	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}

}