package src

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func FetchRaw(sbd string) string {
					
	url := fmt.Sprintf("https://diemthi.laodong.vn/tra-cuu-diem-thi-thpt-xem-diem-thi-dai-hoc-2022.html?sbd=%s")
	// url := fmt.Sprintf("https://thanhnien.vn/api/diemthi/get/all?kythi=THPT&year=2022&city=&text=%s&top=no", sbd)
	res, err := http.Get(url)

	if res.StatusCode == 403 {
		errMsg := "https://diemthi.laodong.vn/ might blocked you. Try to change IP address, reduce patch size/delay and run again."
		errMsg += "\nLast SBD: " + string(sbd)
		panic(errMsg)
	}
	if err != nil {
		fmt.Println(sbd, err)
		return ""
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(sbd, err)
		return ""
	}

	return string(body)
}

func FetchScore(sbd string, ch chan *StudentChannel, chFail chan string, chFinish chan bool) {

	htmlBody := FetchRaw(sbd)
	var std *Student

	if len(htmlBody) == 0 {
		std = nil
		chFail <- sbd
	} else {
		std = ParseStudent(&htmlBody)
	}

	ch <- &StudentChannel{
		id:   sbd,
		data: std,
	}
	chFinish <- true
}
