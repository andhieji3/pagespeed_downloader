## System Requirement
* OS : Windows/Mac OS/Linux
* [Go Compiler](https://golang.org/doc/install)
* [Dep (Go Dependency Management Tool)](https://github.com/golang/dep)

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
![image](https://user-images.githubusercontent.com/31300105/52934192-48b68200-3388-11e9-8481-bba3c4f267dd.png)

* Open link in browser, and allow
![image](https://user-images.githubusercontent.com/31300105/52934245-67b51400-3388-11e9-9954-a37d4df93aad.png)

* Copy token hash and paste it to your command line, it will create new credential (token.json)
![image](https://user-images.githubusercontent.com/31300105/52934287-8f0be100-3388-11e9-9fa6-2a331769f260.png)

## Move It (Only for \*nix user)
* Copy it to /user/local/bin directory
```
> cp $GOPATH/src/pagespeed_reader/bin/pagespeed_reader /usr/local/bin/pagespeed_reader
```
* Now, you can call it by type ```pagespeed_reader``` from your
