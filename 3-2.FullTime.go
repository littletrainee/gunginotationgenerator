package GunGiNotationGenerator

type WholeGameTime struct {
	Date      string
	StartTime string
	EndTime   string
	FileName  string
}

func (f *WholeGameTime) GetTimeTitle() string {
	return "# " + f.Date + "&nbsp;&nbsp;&nbsp;&nbsp; 開始時間：" + f.StartTime +
		"&nbsp;&nbsp;&nbsp;&nbsp;結束時間：" + f.EndTime + "\n\n"
}
