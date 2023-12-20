package commonerr

type Err struct {
	Code     int    `json:"code"`
	ErrorMsg string `json:"msg"`
}

// 中间件中的报错 -10 开头
var (
	ErrAccessToken = Err{-100001, "token 错误:"}
)

// 通用的error 100开头
var (
	ErrParams          = Err{1000001, "参数错误"}
	ErrLocalCacheEmpty = Err{1000002, "本地数据为空"}
	ErrDataExists      = Err{1000003, "数据已存在"}
	ErrFileSize        = Err{1000004, "文件大小错误"}
	ErrFileType        = Err{1000005, "文件类型错误"}
	ErrParamsJson      = Err{1010006, "the value of params must be json"}
)

// 提醒业务使用的error 101开头
var (
	ErrTimeWordsEmpty      = Err{1010001, "没有发现表示时间的词语"}
	ErrTimeWords2TimeStamp = Err{1010002, "没有转换成有效的时间戳或时间已过期:%d"}
	ErrTimeWordsHasParsed  = Err{1010003, "时间词已解析"}
)

// Job中的报错 102 开头
var (
	ErrEmqMsgType = Err{1020001, "err meqtt msg type"}
)
var errList = []Err{
	ErrParams,
}

func GetErrorCode(errMsg string) int {
	for _, err := range errList {
		if err.ErrorMsg == errMsg {
			return err.Code
		}
	}
	return 0
}

func AppenErrorMsg(err Err, appendMsg string) Err {
	err.ErrorMsg += appendMsg
	return err
}
