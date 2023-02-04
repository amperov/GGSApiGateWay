package tooling

import "fmt"

func SignInResponse(AccessCode, Status string) []byte {
	resp := []byte(fmt.Sprintf(`{"access-code": "%s", "status": "%s"}`, AccessCode, Status))
	return resp
}
func RecoverResponse(ActionID, Status string) []byte {
	resp := []byte(fmt.Sprintf(`{"action-uid": %s, "status": "%s"}`, ActionID, Status))
	return resp
}
