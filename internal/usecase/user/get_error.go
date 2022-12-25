package user

type GetError interface {
	error
	isGetError()
}

// GetErrorの想定されたパターン
var (
	// 通信失敗する場合がある
	_ GetError = GetErrorInfrastructural{}
)

type GetErrorInfrastructural struct {
	Err error
}

func (GetErrorInfrastructural) isGetError() {}

func (e GetErrorInfrastructural) Error() string {
	return e.Err.Error()
}
