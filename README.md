# gowk

Use golang like awk.

The gowk is utility of run golang program in command line, without editor/commands.

## install

`go get github.com/kuwa72/gowk`

## examples

### Hello world.

`gowk -r 'fmt.Println("Hello world")'`

### like cat(1)

`gowk -n -r 'fmt.Println(s[0])'`

### HTTP Server(Hello world)

`gowk -i net/http -r 'http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {fmt.Fprintln(w, "Hello world")});http.ListenAndServe(":8888", nil)'`


## Author

kuwa72 https://github.com/kuwa72 @kuwashima

## License

MIT