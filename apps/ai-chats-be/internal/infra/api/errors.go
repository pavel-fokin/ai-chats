package api

var (
	NotFound = Error{
		Message: "Not found.",
	}
	BadRequest = Error{
		Message: "Bad request.",
	}
	Unauthorized = Error{
		Message: "Unauthorized.",
	}
	InternalError = Error{
		Message: "Internal error.",
	}
	Forbidden = Error{
		Message: "Forbidden.",
	}
	UsernameIsTaken = Error{
		Message: "That username is already taken. Try another one.",
	}
	UsernameOrPasswordIsIncorrect = Error{
		Message: "Username or password is incorrect. Please try again.",
	}
)
