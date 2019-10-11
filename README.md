## Requirement
* OS : Windows/Mac OS/Linux
* [Go Compiler](https://golang.org/doc/install)
* [Dep (Go Dependency Management Tool)](https://github.com/golang/dep)
* Directory **PageSpeed (Cron)** in your google drive

## Setting
* Define your **$GOPATH** environment & workspace [Tutorial](https://dasarpemrogramangolang.novalagung.com/3-gopath-dan-workspace.html), don't forget 3 core directory [bin, pkg, src]
* Go to directory **$GOPATH/src**
* Clone this git, directory name is strict (except your change the code)
```
> git clone https://github.com/andhieji3/pagespeed_reader.git pagespeed_reader
```
* Delete All file in **$GOPATH/src/pagespeed_reader/configs/drive** and **$GOPATH/src/pagespeed_reader/configs/spreadsheet**

## Generate API Credential
* Generate credential.json for Drive API [Link](https://developers.google.com/drive/api/v3/quickstart/go), download and move it to **$GOPATH/src/pagespeed_reader/configs/drive**
* Generate credential.json for Sheets API [Link](https://developers.google.com/sheets/api/quickstart/go), download and move it to **$GOPATH/src/pagespeed_reader/configs/spreadsheet**

## Build & Run It
* Download all dependency using DEP
```
> dep ensure
```
* Build using this command
```
go build -o $GOPATH/src/pagespeed_reader/bin/pagespeed_reader $GOPATH/src/pagespeed_reader/cmd/pagespeed_reader/main.go
```
* Run compiled program
```
$GOPATH/src/pagespeed_reader/bin/pagespeed_reader 
```
* After get pagespeed result, it will try connect to google drive API, if token not exist it will give you link.
![image](https://user-images.githubusercontent.com/31300105/66664170-3de19200-ec76-11e9-967c-94b6ad81ef81.png)

* Open link in browser, and allow
![image](https://user-images.githubusercontent.com/31300105/66664275-6a95a980-ec76-11e9-94cc-7f80f9508011.png)

* Copy token hash and paste it to your command line, it will create new credential (token.json)
![image](https://user-images.githubusercontent.com/31300105/66664390-a2045600-ec76-11e9-9a2e-63e6f814e5d3.png)

## Move It (Only for \*nix user)
* Copy it to /user/local/bin directory
```
> cp $GOPATH/src/pagespeed_reader/bin/pagespeed_reader /usr/local/bin/pagespeed_reader
```
* Now, you can call it by type ```pagespeed_reader``` from your
