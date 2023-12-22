package utils

import (
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2"
	"os"
	"time"
)

type SiteVerifyResponse struct {
	Success     bool      `json:"success"`
	Score       float64   `json:"score"`
	Action      string    `json:"action"`
	ChallengeTS time.Time `json:"challenge_ts"`
	Hostname    string    `json:"hostname"`
	ErrorCodes  []string  `json:"error-codes"`
}

type SiteVerifyRequest struct {
	RecaptchaResponse string `json:"g-recaptcha-response"`
}

const siteVerifyURL = "https://www.google.com/recaptcha/api/siteverify"

func CheckRecaptcha(secret, response string) error {

	a := fiber.AcquireAgent()
	req := a.Request()
	var obj SiteVerifyResponse
	a.Debug()
	req.Header.SetMethod(fiber.MethodPost)
	req.SetRequestURI(siteVerifyURL)
	a.JSON(fiber.Map{"secret": os.Getenv("RECAPTCHA_SECRET")})
	a.QueryString(
		"secret=" + secret + "&response=" + response)

	if err := a.Parse(); err != nil {
		return err
	}

	code, body, errs := a.Bytes()

	if errs != nil {
		return errs[0]
	}

	if code != 200 {
		return errs[1]
	}

	if err := json.Unmarshal(body, &obj); err != nil {
		return err
	}
	if !obj.Success {
		return errors.New("unsuccessful recaptcha verify request")
	}

	// Check response score.
	if obj.Score < 0.5 {
		return errors.New("lower received score than expected")
	}

	// Check response action.
	if obj.Action != "login" {
		return errors.New("mismatched recaptcha action")
	}

	return nil
}
