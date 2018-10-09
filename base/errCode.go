package base

// 10000~20000系统错误
// 20000~30000操作错误
const (
	OpenFileErr      = 10000
	ReadFileErr      = 10001
	PrintHtmlErr     = 10002
	TypeConvertErr   = 10003
	DelFileErr       = 10004
	SaveFileErr      = 10005
	JsonErr          = 10006
	PostBodyReadFail = 10007

	NoBackupFile      = 20000
	NoOwnThing        = 20001
	AlreadyMaxed      = 20002
	MaterialNotEnough = 20003
	AlreadyGrowed     = 20004
	NotGrowed         = 20005
	BagNotEnough      = 20006
	PlaceNothing      = 20007

	ParaInvalid = 40000

	HeroDead     = 80000
	ProgramPanic = 90000
)

type DEAD struct {
	Reason string `json:"reason"`
}
