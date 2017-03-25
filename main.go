package main

import runtime "github.com/ncodes/cocoon/core/runtime/golang"

var log = runtime.GetLogger()

// App is an example cocoon code
type App struct {
}

// OnInit method initializes the app
func (app *App) OnInit(link *runtime.Link) error {
	log.Info("App is initializing!")
	log.Info("GetCocoonID(): ", runtime.GetCocoonID())
	log.Info("GetID(): ", runtime.GetID())
	return nil
}

// OnInvoke process invoke transactions
func (app *App) OnInvoke(link *runtime.Link, txID, function string, params []string) (interface{}, error) {
	log.Info("Invoked")
	return nil, nil
}

func main() {

	runtime.Run(new(App))
	log.Info("Cocoon Code Stopped")

	// fmt.Println("Hello Friend")
	// res, err := http.Get("https://google.com.ng")
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

	// output, err := exec.Command("bash", "-c", "iptables -S").Output()
	// if err != nil {
	// 	fmt.Println("Err:", err, string(output))
	// 	return
	// }

	// fmt.Println(len(output), string(output))
}
