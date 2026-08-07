package main

type Config struct {
	ConnectionString string
}

func main() {

	i := 0
	for i = 0; i < 20; i++ {
		haveConfig := GetConfOnce()
		println(i, ": ", &haveConfig)
	}

}
