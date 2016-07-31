package nganluong

import (
    "strings"
    "gitlab.com/hs-api-go/util/stringutil"
    "fmt"
    "net/url"
)

const (
    SITE_URL = "http://localhost:8000"
    NGANLUONG_URL = "https://www.nganluong.vn/checkout.php"
    MERCHANT_ID = "46592"
    MERCHANT_PASS = "65bcc74e96d315edf154027fe494a66c"
    RECEIVER = "minhanhhere@gmail.com"
    CURRENCY = "vnd"
    RETURN_URL = SITE_URL + "/payment/success"
    CANCEL_URL = "%s/payment/new/%s?action=cancel&package_id=%s"
    AFFILIATE_CODE = ""
)

type BuyerInfo struct {
    Name    string
    Email   string
    Phone   string
    Address string
}

func (buyer *BuyerInfo) getInfoString() string {
    param := []string{buyer.Name, buyer.Email, buyer.Phone, buyer.Address}
    return strings.Join(param, "*|*")
}

func CreateNganLuongPaymentUrl(price float64, orderCode string, orderDesc string, buyerinfo BuyerInfo, projectId string, packageId string) string {

    transInfo := projectId + "_" + packageId

    url, _ := url.Parse(NGANLUONG_URL)
    query := url.Query()
    query.Set("merchant_site_code", MERCHANT_ID)
    query.Set("return_url", RETURN_URL)
    query.Set("cancel_url", fmt.Sprintf(CANCEL_URL, SITE_URL, projectId, packageId))
    query.Set("receiver", RECEIVER)
    query.Set("currency", CURRENCY)
    query.Set("affiliate_code", AFFILIATE_CODE)
    query.Set("quantity", "1")
    query.Set("tax", "0")
    query.Set("discount", "0")
    query.Set("fee_cal", "0")
    query.Set("fee_shipping", "0")
    query.Set("price", stringutil.ValueOfFloat64(price))
    query.Set("order_code", orderCode)
    query.Set("order_description", orderDesc)
    query.Set("transaction_info", transInfo)
    query.Set("buyer_info", buyerinfo.getInfoString())
    query.Set("secure_code", CreateSecureCode(buyerinfo, orderCode, orderDesc, price, transInfo))

    url.RawQuery = query.Encode()
    fmt.Println(url)
    return url.String()
}

func VerifyNganLuongResult(transInfo string, orderCode string, price string, paymentId string, paymentType string, errorText string, secureCode string) bool {
    param := []string{
        transInfo, orderCode, price, paymentId, paymentType, errorText, MERCHANT_ID, MERCHANT_PASS,
    }
    raw := strings.Join(param, " ")
    return stringutil.MD5String(raw) == secureCode
}

func CreateSecureCode(buyerinfo BuyerInfo, orderCode string, orderDesc string, price float64, transInfo string) string {
    param := []string{
        MERCHANT_ID, RETURN_URL, RECEIVER, transInfo, orderCode, stringutil.ValueOfFloat64(price), CURRENCY,
        "1", "0", "0", "0", "0", orderDesc, buyerinfo.getInfoString(), AFFILIATE_CODE, MERCHANT_PASS,
    }
    rawSecure := strings.Join(param, " ")
    return stringutil.MD5String(rawSecure)
}