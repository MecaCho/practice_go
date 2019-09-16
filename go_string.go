//golang中字符串和各种int类型之间的相互转换方式：

//string转成int：
int, err := strconv.Atoi(string)

//string转成int64：
int64, err := strconv.ParseInt(string, 10, 64)


//int转成string：
string := strconv.Itoa(int)

//int64转成string：
string := strconv.FormatInt(int64,10)
