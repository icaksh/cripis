package utils

import (
	"fmt"
)

func IP_Duplicate(name string, stype string) (bool, error){
	data, err := FetchDataFromApi(stype, name, "asc" ,1)
	if err != nil {
		return false, err
	}
	for _, record := range data{
		if rec, ok := record.(map[string]interface{}); ok {
			for key, _ := range rec{
				fmt.Println(key)
				// if string(key[]) == strings.ToUpper(name){
				// 	return true, nil
				// }
				break
			}
		}
	}
	return false, nil
}