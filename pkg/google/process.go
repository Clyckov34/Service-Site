package google

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type ResponseCaptcha struct {
	Success     bool        `json:"success"`
	Score       float64     `json:"score"`
	Action      string      `json:"action"`
	ChallengeTs time.Time   `json:"challenge_ts"`
	Hostname    string      `json:"hostname"`
	ErrorCode   interface{} `json:"error-codes"`
}

//CaptchaCheck капча проверка на стороне google
func CaptchaCheck(ctx *gin.Context) error {
	if len(os.Getenv("GOOGLE_KEY_SITE")) > 0 && len(os.Getenv("GOOGLE_KEY_SERVER")) > 0 && len(os.Getenv("GOOGLE_SCORE")) > 0 {
		key := strings.TrimSpace(ctx.PostForm("g-recaptcha-response"))

		resCap, err := getResponse(key, ctx.ClientIP())
		if err != nil {
			return err
		}
	
		score, err := strconv.ParseFloat(os.Getenv("GOOGLE_SCORE"), 64)
		if err != nil {
			return err
		}
	
		if !resCap.Success || resCap.Score < score {
			return errors.New("вы не прошли проверку reCAPTCHA v3")
		}
	
		return nil
	} else {
		return errors.New("не соотвествуют параметры данные")
	}
}

//getResponse google капча проверка на стороне google запрос
func getResponse(key, ip string) (ResponseCaptcha, error) {
	data := url.Values{
		"secret":   {os.Getenv("GOOGLE_KEY_SERVER")},
		"response": {key},
		"remoteip": {ip},
	}

	resBody, err := http.PostForm("https://www.google.com/recaptcha/api/siteverify", data)
	if err != nil {
		return ResponseCaptcha{}, err
	}
	defer resBody.Body.Close()

	body, err := ioutil.ReadAll(resBody.Body)
	if err != nil {
		return ResponseCaptcha{}, err
	}

	var res ResponseCaptcha
	err = json.Unmarshal(body, &res)
	if err != nil {
		return ResponseCaptcha{}, err
	}

	return res, nil
}
