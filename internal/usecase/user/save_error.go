package user

type SaveError interface {
	error
	isSaveError()
}

// GetErrorの想定されたパターン
var (
	// 通信失敗する場合がある
	_ SaveError = SaveErrorInfrastructural{}
)

type SaveErrorInfrastructural struct {
	Err error
}

func (SaveErrorInfrastructural) isSaveError() {}

func (e SaveErrorInfrastructural) Error() string {
	return e.Err.Error()
}
