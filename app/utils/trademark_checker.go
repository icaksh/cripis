package utils

import (
	"fmt"
)

func TrademarkSimilarity(name string) (bool, error) {
	data, err := FetchDataFromApi(name)
	if err != nil {
		return false, err
	}
	for _, record := range data {
		if rec, ok := record.(map[string]interface{}); ok {
			for key, _ := range rec {
				fmt.Println(key)
				break
			}
		}
	}
	return false, nil
}
