package main

import (
	"fmt"
	"os/exec"
)

func main() {
	output, err := exec.Command("bash", "-c", "iptables -S").Output()
	if err != nil {
		fmt.Println("Err:", err, string(output))
		return
	}

	fmt.Println(len(output), string(output))

	// res, err := http.Get("http://localhost:8500/v1/catalog/services")
	// if err != nil {
	// 	fmt.Println("Err: ", err)
	// 	return
	// }

	// defer res.Body.Close()
	// fmt.Println(res.StatusCode)
	// _, err = io.Copy(os.Stdout, res.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
