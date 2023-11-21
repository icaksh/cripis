package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)


func FetchDataFromApi(stype string, keyword string, orderState string, page int,) (map[string]interface{},error) {
	a := fiber.AcquireAgent()
	req := a.Request()
	var obj fiber.Map
	a.Debug()
	req.Header.SetMethod(fiber.MethodPost)
	req.SetRequestURI("https://pdki-indonesia-api.dgip.go.id/api/"+stype+"/search2?")
	a.QueryString(
		"keyword="+keyword+"&order_state="+orderState+"&page="+strconv.Itoa(page)+"&type="+stype)
	if(stype!="trademark"){
		a.JSON(fiber.Map{"key":PrivateKey("")})
	}else{
		a.JSON(fiber.Map{"key":PrivateKey(keyword)})
	}
		

	if err := a.Parse(); err != nil {
		return obj, err
	}

	code, body, errs := a.Bytes()

	if errs != nil {
		return obj, errs[0]
	}
	
	if(code!=200){
		return obj, errs[1]
	}

	if err := json.Unmarshal(body, &obj); err != nil {
		return obj,err
	}
	return obj, nil
}

func PrivateKey(value string) string {

	now := time.Now()
	year, month, day := now.Date()

	monthStr := strconv.Itoa(int(month))
	if len(monthStr) == 1 {
		monthStr = "0" + monthStr
	}

	dayStr := strconv.Itoa(day)
	if len(dayStr) == 1 {
		dayStr = "0" + dayStr
	}

	timestamp := fmt.Sprintf("%d%s%s", year, monthStr, dayStr)
	fmt.Println(timestamp)
	combinedString := timestamp + value + os.Getenv("PDKI_SECRET_KEY")

	hash := sha256.New()
	hash.Write([]byte(combinedString))
	hashInBytes := hash.Sum(nil)

	hashString := hex.EncodeToString(hashInBytes)
	return hashString
}