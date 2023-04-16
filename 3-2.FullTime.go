package GunGiNotationGenerator

type WholeGameTime struct {
	Date         string
	StartTime    string
	SetUpEndTime string
	EndTime      string
	FileName     string
}

func (f *WholeGameTime) GetTimeTitle() string {
	return "# " + f.Date + "\n## 開始時間：" + f.StartTime + "\n## 布陣結束時間：" +
		f.SetUpEndTime + "\n## 結束時間：" + f.EndTime + "\n\n"
}
