package main

import (
	"os"
	"time"

	"github.com/therecipe/qt/widgets"
)

func main() {

	widgets.NewQApplication(len(os.Args), os.Args)

	// ウィンドウ作成
	window := widgets.NewQMainWindow(nil, 0)
	window.SetWindowTitle("占いアプリ")
	window.SetMinimumSize2(200, 200)

	// レイアウト作成
	layout := widgets.NewQVBoxLayout()

	// ウィジェット作成
	widget := widgets.NewQWidget(nil, 0)
	widget.SetLayout(layout)

	// 日付入力
	dateEdit := widgets.NewQDateEdit(nil)
	layout.AddWidget(dateEdit, 0, 0)

	// ボタン作成
	button := widgets.NewQPushButton2("生年月日を入力したら押してね(•ө•)♡", nil)
	button.ConnectClicked(func(checked bool) {

		// 日付オブジェクトに変換
		timeObject, _ := time.Parse("2000/01/02", dateEdit.Text())

		resultStr := "あなたは"
		switch int(timeObject.Month()) {
		case 1:
			resultStr += "山羊座"
		case 2:
			resultStr += "水瓶座"
		case 3:
			resultStr += "魚座"
		case 4:
			resultStr += "牡羊座"
		case 5:
			resultStr += "牡牛座"
		case 6:
			resultStr += "双子座"
		case 7:
			resultStr += "蟹座"
		case 8:
			resultStr += "しし座"
		case 9:
			resultStr += "乙女座"
		case 10:
			resultStr += "てんびん座"
		case 11:
			resultStr += "蠍座"
		case 12:
			resultStr += "射手座"
		}

		resultStr += "です"
		widgets.QMessageBox_Information(nil, "OK", resultStr, widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
	})
	layout.AddWidget(button, 0, 0)

	// ウィンドウにウィジェットを設定
	window.SetCentralWidget(widget)

	// ウィンドウを表示
	window.Show()

	// メインループ開始
	widgets.QApplication_Exec()
}
