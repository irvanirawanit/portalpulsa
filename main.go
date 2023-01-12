package portalpulsa

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

type PortalPulsa struct {
	PortalUserId string
	PortalKey    string
	PortalSecret string
}

func (p *PortalPulsa) logToFile(req *http.Request, writerform *multipart.Writer, resbody []byte) {
	// create folder if not exist
	if _, err := os.Stat("portalpulsa"); os.IsNotExist(err) {
		os.Mkdir("portalpulsa", 0755)
	}

	currendate := time.Now().Format("2006-01-02")
	// Create log file
	file, err := os.OpenFile("portalpulsa/"+currendate+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	requestFieldsString := ""
	for key, value := range req.Header {
		requestFieldsString += fmt.Sprintf("%s: %s \r ", key, value)
	}

	// Write to log file
	timenow := time.Now().Format("2006-01-02 15:04:05")
	if _, err := file.WriteString(timenow + " " + req.Method + " " + req.URL.String() + " \n " + requestFieldsString + "  " + string(resbody) + " \n"); err != nil {
		log.Fatal(err)
	}
}

func (p *PortalPulsa) Harga(code string) map[string]interface{} {
	if code == "" {
		code = "pulsa"
	}
	url := "https://portalpulsa.com/api/connect/"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("inquiry", "HARGA")
	_ = writer.WriteField("code", code)
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("portal-userid", p.PortalUserId) // "Pxxxxxx"
	req.Header.Add("portal-key", p.PortalKey)       // "xxx..."
	req.Header.Add("portal-secret", p.PortalSecret) // "xxx..."

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	resbody := make(map[string]interface{})
	json.Unmarshal(body, &resbody)

	return resbody
}

func (p *PortalPulsa) CekSaldo() map[string]interface{} {
	url := "https://portalpulsa.com/api/connect/"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("inquiry", "s")
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("portal-userid", p.PortalUserId) // "Pxxxxxx"
	req.Header.Add("portal-key", p.PortalKey)       // "xxx..."
	req.Header.Add("portal-secret", p.PortalSecret) // "xxx..."

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	resbody := make(map[string]interface{})
	json.Unmarshal(body, &resbody)

	return resbody
}

func (p *PortalPulsa) TopUp(code string, phone string, idcust string, trxid_api string, no string) map[string]interface{} {
	url := "https://portalpulsa.com/api/connect/"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("inquiry", "I")         // konstan
	_ = writer.WriteField("code", code)           // kode produk
	_ = writer.WriteField("phone", phone)         // nohp pembeli
	_ = writer.WriteField("idcust", idcust)       // Diisi jika produk memerlukan IDcust seperti: Unlock/Aktivasi Voucher, Game Online (FF, ML, PUBG, dll)
	_ = writer.WriteField("trxid_api", trxid_api) // Trxid / Reffid dari sisi client
	_ = writer.WriteField("no", no)               // untuk isi lebih dari 1x dlm sehari, isi urutan 1,2,3,4,dst
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("portal-userid", p.PortalUserId) // "Pxxxxxx"
	req.Header.Add("portal-key", p.PortalKey)       // "xxx..."
	req.Header.Add("portal-secret", p.PortalSecret) // "xxx..."

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	resbody := make(map[string]interface{})
	json.Unmarshal(body, &resbody)

	// log to file
	p.logToFile(req, writer, body)

	return resbody
}

func (p *PortalPulsa) BeliTokenPLN(code string, phone string, idcust string, trxid_api string, no string) map[string]interface{} {
	url := "https://portalpulsa.com/api/connect/"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("inquiry", "s")         // konstan
	_ = writer.WriteField("code", code)           // kode produk
	_ = writer.WriteField("phone", phone)         // nohp pembeli
	_ = writer.WriteField("idcust", idcust)       // nomor meter atau id pln
	_ = writer.WriteField("trxid_api", trxid_api) // Trxid / Reffid dari sisi client
	_ = writer.WriteField("no", no)               // untuk isi lebih dari 1x dlm sehari, isi urutan 1,2,3,4,dst
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("portal-userid", p.PortalUserId) // "Pxxxxxx"
	req.Header.Add("portal-key", p.PortalKey)       // "xxx..."
	req.Header.Add("portal-secret", p.PortalSecret) // "xxx..."

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	resbody := make(map[string]interface{})
	json.Unmarshal(body, &resbody)

	return resbody
}

func (p *PortalPulsa) StatusTransaksi(trxid_api string) map[string]interface{} {
	url := "https://portalpulsa.com/api/connect/"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("inquiry", "STATUS")
	_ = writer.WriteField("trxid_api", trxid_api) // Trxid atau Reffid dari sisi client saat transaksi pengisian
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("portal-userid", p.PortalUserId) // "Pxxxxxx"
	req.Header.Add("portal-key", p.PortalKey)       // "xxx..."
	req.Header.Add("portal-secret", p.PortalSecret) // "xxx..."

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	resbody := make(map[string]interface{})
	json.Unmarshal(body, &resbody)

	return resbody
}

func (p *PortalPulsa) RequestSaldoDeposit(bank string, nominal string) map[string]interface{} {
	url := "https://portalpulsa.com/api/connect/"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("inquiry", "D")
	_ = writer.WriteField("bank", bank)
	_ = writer.WriteField("nominal", nominal)
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("portal-userid", p.PortalUserId) // "Pxxxxxx"
	req.Header.Add("portal-key", p.PortalKey)       // "xxx..."
	req.Header.Add("portal-secret", p.PortalSecret) // "xxx..."

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	resbody := make(map[string]interface{})
	json.Unmarshal(body, &resbody)

	return resbody
}
