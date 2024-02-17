package actionresults

type ErrorActionResult struct {
	error
}

func (action ErrorActionResult) Execute(ctx *ActionContext) error {
	return action.error
}

func NewErrorAction(err error) ActionResult {
	return &ErrorActionResult{err}
}
