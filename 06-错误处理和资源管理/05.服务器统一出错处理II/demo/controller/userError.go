package controller

// 可以给在页面上展示给用户看的error
type UserError interface {
	error
	Message() string
}