# 使用しているgoのパッケージ
go get -u -v github.com/nsf/gocode
go get -u -v github.com/rogpeppe/godef
go get -u -v github.com/zmb3/gogetdoc
go get -u -v github.com/golang/lint/golint
go get -u -v github.com/lukehoban/go-outline
go get -u -v sourcegraph.com/sqs/goreturns
go get -u -v golang.org/x/tools/cmd/gorename
go get -u -v github.com/tpng/gopkgs
go get -u -v github.com/newhook/go-symbols
go get -u -v golang.org/x/tools/cmd/guru
go get -u -v github.com/cweill/gotests/...

go get -u -v github.com/therecipe/qt/cmd/...

# トラブルシュート
* could not launch process: exec: “lldb-server”: executable file not found in $PATH
  * xcode-select --install

* qtのインストール
  * brew install qt5
  * QT_HOMEBREW=true
  * $GOPATH/bin/qtsetup

* qtアプリの起動
  * `qtdeploy test desktop`