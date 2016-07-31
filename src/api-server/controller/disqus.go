package controller

import (
    "github.com/labstack/echo"
    "fmt"
    "net/http"
    "encoding/base64"
    "encoding/json"
    "crypto/hmac"
    "crypto/sha1"
    "text/template"
    "bytes"
    "time"
)

const DISQUS_SECRET_KEY = "kWbsZY1IhAVvIlBCNLzUAQm2hUg4Gw4Cci2fBqzknLbslADvz6b3FEL61gxWPnCm"
const DISQUS_PUBLIC_KEY = "pwYfjd2zE5jyz1sBp2ao1B1LnaeujaJ9ptBx5uxF5kUuFd66ymn6g5oMq0ixIx4C"
const SCRIPT_TEMPLATE = `
    <script type="text/javascript">
        var disqus_config = function() {
            this.page.remote_auth_s3 = "{{.Message}} {{.Signature}} {{.Timestamp}}";
            this.page.api_key = "{{.PublicKey}}";
            this.page.url = 'http://localhost:9000/#/project/f7d5f1ca-be8a-41cb-a9ed-29386c8f9ca4';
            this.page.identifier = 'f7d5f1ca-be8a-41cb-a9ed-29386c8f9ca4';

        };
        (function() {  // DON'T EDIT BELOW THIS LINE
            var d = document, s = d.createElement('script');
            s.src = 'https://bookathon.disqus.com/embed.js';
            s.setAttribute('data-timestamp', +new Date());
            (d.head || d.body).appendChild(s);
        })();
    </script>
`

type DisqusParam struct {
    Message   string
    Signature string
    Timestamp int64
    PublicKey string
}

var Disqus struct {
    Get echo.HandlerFunc
}

func init() {
    fmt.Println("init controller.disqus")
    Disqus.Get = func(c echo.Context) error {
        return c.String(http.StatusOK, GetDisqusSingleSignOn("1", "taylor", "taylorswift@gmail.com"))
    }
}

func RegisterDisqusRoute(api *echo.Group) {
    api.GET("/disqus", Disqus.Get)
}

func GetDisqusSingleSignOn(id string, username string, email string) string {

    var data = map[string]string{
        "id": id,
        "username": username,
        "email": email,
    }
    dataString, _ := json.Marshal(data)
    message := base64.StdEncoding.EncodeToString(dataString)
    timestamp := time.Now().Unix()
    mac := hmac.New(sha1.New, []byte(DISQUS_SECRET_KEY))
    mac.Write([]byte(fmt.Sprintf("%s %d", message, timestamp)))
    signature := mac.Sum(nil)
    param := DisqusParam{
        Message : message,
        Signature: fmt.Sprintf("%x", signature),
        Timestamp: timestamp,
        PublicKey: DISQUS_PUBLIC_KEY,
    }
    var doc bytes.Buffer
    tmpl, _ := template.New("disqus_script").Parse(SCRIPT_TEMPLATE)
    tmpl.Execute(&doc, param)

    return doc.String()
}