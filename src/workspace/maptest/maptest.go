package maptest

func Maptest() map[string]int{
	// マップを定義する
	a1 := map[string]int{
		"x": 100,
		"y": 200,			// 改行する場合はカンマ必須
	}

	// マップに要素を加える
	a1["z"] = 300

	// マップの要素を削除する
	delete(a1, "z")
	return a1
	}