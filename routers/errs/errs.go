package errs

import (
	"net/http"

	"hito/routers/resp"
)

/*
 * Code Format: [xx][xxxx]
 * The first code is function code, the second code is error code.
 */
var (
	InternalError        = resp.Error{http.StatusInternalServerError, resp.System.String() + "0001", "Internal Error"}
	InvalidToken         = resp.Error{http.StatusUnauthorized, resp.System.String() + "0010", "Invalid token"}
	InvalidParams        = resp.Error{http.StatusBadRequest, resp.System.String() + "0011", "Invalid parameters"}
	CanNotCreateUser     = resp.Error{http.StatusBadRequest, resp.System.String() + "0012", "Can not create user"}
	CanNotLogin          = resp.Error{http.StatusBadRequest, resp.System.String() + "0013", "Invalid id or password"}
	CanNotChangePassword = resp.Error{http.StatusBadRequest, resp.System.String() + "0014", "Invalid old password"}

// InvalidParams                  = resp.Error{http.StatusBadRequest, resp.HMX.String() + "0011", "Invalid parameters"}
// CanNotAssignDataToModel        = resp.Error{http.StatusBadRequest, resp.HMX.String() + "0012", "Can not assign data to model"}
// UserNotFound                   = resp.Error{http.StatusBadRequest, resp.HMX.String() + "0100", "User Not Found"}
// InvalidLogin                   = resp.Error{http.StatusUnauthorized, resp.HMX.String() + "0101", "Invalid id or password"}
// CanNotCreateSession            = resp.Error{http.StatusBadRequest, resp.HMX.String() + "0102", "Can create session"}
// InvalidOldPass                 = resp.Error{http.StatusBadRequest, resp.HMX.String() + "0103", "Invalid old password"}
// CanNotChangePass               = resp.Error{http.StatusBadRequest, resp.HMX.String() + "0104", "Can not change password"}
// CanNotCreateUser               = resp.Error{http.StatusBadRequest, resp.HMX.String() + "0105", "Can not create user"}
// CanNotCreateEmailVerifyToken   = resp.Error{http.StatusBadRequest, resp.HMX.String() + "0106", "Can not create email verify token"}
// EmailVerified                  = resp.Error{http.StatusBadRequest, resp.HMX.String() + "0107", "Email has been verified"}
// CanNotCreatePasswordResetToken = resp.Error{http.StatusBadRequest, resp.HMX.String() + "0108", "Can not create password reset token"}
// CanNotChangePassByToken        = resp.Error{http.StatusBadRequest, resp.HMX.String() + "0109", "Can change pasword by token"}
// CanNotFindModel                = resp.Error{http.StatusBadRequest, resp.HMX.String() + "0120", "Can not find model"}
// CanNotUpdateModel              = resp.Error{http.StatusBadRequest, resp.HMX.String() + "0121", "Can not update model"}
// UserHasAlreadyCreated          = resp.Error{http.StatusBadRequest, resp.HMX.String() + "0122", "Account already created"}
)

// func Msg(err resp.Error, msg string) resp.Error {
// 	err.Message = err.Error() + ": " + msg
// 	return err
// }

func Msg(err error, msg string) error {

	return err
}
