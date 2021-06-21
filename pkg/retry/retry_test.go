package retry

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func TestBackOffRetryMaxTimesMaxDurationUntilNoError(t *testing.T) {
	err := BackOffRetryMaxTimesMaxDurationUntilNoError(2, 1*time.Minute, func() error {
		var body []byte
		url := "https://www.yahoo.com/"
		res, err := http.Get(url)
		if err != nil {
			return err
		}
		defer res.Body.Close()
		body, err = ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}
		fmt.Println(string(body))
		return nil
	})
	if err != nil {
		t.Fatalf("error:%v", err)
	} else {
		t.Log("Test pass")
	}

}
