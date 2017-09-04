package define

// 参数必选和可选值定义
const (
	PARMS_MUST_STR       = ""
	PARMS_UNMUST_STR     = "No must"
	PARMS_MUST_INT       = 0
	PARMS_UNMUST_INT     = -1
	PARMS_MUST_FLOAT64   = 0.0
	PARMS_UNMUST_FLOAT64 = -1.0
)

type Out struct {
	Ret  int                    `json:"ret"`
	Msg  string                 `json:"msg"`
	Data map[string]interface{} `json:"data"`
}
