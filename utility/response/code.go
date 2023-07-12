package response

// 使用错误码: return gerror.NewCode(response.CodeNotFound)

var (
	CodeOK       = New(200, "OK", "")
	CodeNotFound = New(404, "Not Found", "Resource does not exist")
	CodeInternal = New(500, "Internal Error", "An error occurred internally")
)
