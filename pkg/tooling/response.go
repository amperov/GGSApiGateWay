package tooling

import "fmt"

func SignInResponse(AccessCode, Refresh, Status string) []byte {
	resp := []byte(fmt.Sprintf(`{"access-token": "%s", "refresh-token": "%s", "status": "%s"}`, AccessCode, Refresh, Status))
	return resp
}
func RecoverResponse(ActionID, Status string) []byte {
	resp := []byte(fmt.Sprintf(`{"action-uid": %s, "status": "%s"}`, ActionID, Status))
	return resp
}
