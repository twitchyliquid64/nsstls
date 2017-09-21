package main

/*
#include <pwd.h>
#include <sys/types.h>
*/
import "C"

// http://www.gnu.org/software/libc/manual/html_node/NSS-Modules-Interface.html
const (
	NssStatusSuccess     = 1
	NssStatusNotFound    = 0
	NssStatusUnavailable = -1
	NssStatusTryAgain    = -2
)

func main() {}

// initialize returns false if startup failed.
func initialize() bool {
	if initSuccessful {
		return true
	}
	if initError != nil {
		return false
	}
	initError = LoadConfig(configPath)
	if initError == nil {
		initSuccessful = true
		return true
	}
	fatal("init", initError)
	return false
}

func set(pwd *C.struct_passwd, User user) {
	pwd.pw_uid = C.__uid_t(User.UID)
	pwd.pw_name = C.CString(User.Username)
	if User.Directory == "" {
		pwd.pw_dir = C.CString("/home/" + User.Username)
	} else {
		pwd.pw_dir = C.CString(User.Directory)
	}
	if User.Shell == "" {
		pwd.pw_shell = C.CString("/bin/bash")
	} else {
		pwd.pw_shell = C.CString(User.Shell)
	}
	pwd.pw_gid = C.__gid_t(User.GID)
	pwd.pw_passwd = C.CString("x")
	pwd.pw_gecos = C.CString(User.Gecos)
}

//export _nss_tls_getpwnam_r
func _nss_tls_getpwnam_r(name *C.char, pwd *C.struct_passwd, buffer *C.char, bufsize C.size_t, result **C.struct_passwd) C.int {
	if !initialize() {
		return NssStatusUnavailable
	}
	resp, err := getUserByName(C.GoString(name), configuration.Token)
	if err != nil {
		fatal("GETPWNAM_R", err)
		return NssStatusTryAgain
	}
	set(pwd, resp.User)
	result = &pwd // nolint: ineffassign
	return NssStatusSuccess
}

//export _nss_tls_getpwuid_r
func _nss_tls_getpwuid_r(uid C.__uid_t, pwd *C.struct_passwd, buffer *C.char, bufsize C.size_t, result **C.struct_passwd) C.int {
	if !initialize() {
		return NssStatusUnavailable
	}
	resp, err := getUserByUID(int(uid), configuration.Token)
	if err != nil {
		fatal("GETPWNAM_R", err)
		return NssStatusTryAgain
	}
	set(pwd, resp.User)
	result = &pwd // nolint: ineffassign
	return NssStatusSuccess
}

//export _nss_tls_setpwent
func _nss_tls_setpwent() C.int {
	return NssStatusNotFound
}

//export _nss_tls_endpwent
func _nss_tls_endpwent() {
}

//export _nss_tls_getpwent_r
func _nss_tls_getpwent_r(pwd *C.struct_passwd, buffer *C.char, bufsize C.size_t, result **C.struct_passwd) C.int {
	return NssStatusNotFound
}
