package actionresults

import "encoding/json"

type JsonActionResult struct {
	data any
}

func (action JsonActionResult) Execute(ctx *ActionContext) error {
	ctx.ResponseWriter.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(ctx.ResponseWriter)
	return encoder.Encode(action.data)
}

func NewJsonAction(data any) ActionResult {
	return &JsonActionResult{data: data}
}
