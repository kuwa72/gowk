# gowk

Use golang like awk.

The gowk is utility of run golang program in command line, without editor/commands.

## install

`go get github.com/kuwa72/gowk`

## usage

```
Usage of gowk:
   gowk [-v] [-n] [-i package] [-i ...] [-d definition-code] [-b begin-code] [-e end-code] -r codes⏎ 
```

### options

* -h: Show usage.
* -n: Read line and process in main code. Datas expand to variable 's'. s[0] contains full line data. s[1], s[2]... contains word unit data.
* -v: verbose mode. show full source code.
* -i pkg: import package.
* -d script: definition code.
* -r script: main code.
* -b script: codes execute before main.
* -e script: codes execute after main.


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
