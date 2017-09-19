package main

/*
#include <pwd.h>
#include <sys/types.h>
*/
import "C"
import "fmt"

// http://www.gnu.org/software/libc/manual/html_node/NSS-Modules-Interface.html
const (
	NssStatusSuccess     = 1
	NssStatusNotFound    = 0
	NssStatusUnavailable = -1
	NssStatusTryAgain    = -2
)

func main() {}

//export _nss_tls_getpwnam_r
func _nss_tls_getpwnam_r(name *C.char, pwd *C.struct_passwd, buffer *C.char, bufsize C.size_t, result **C.struct_passwd) C.int {
	return NssStatusNotFound
}

//export _nss_tls_getpwuid_r
func _nss_tls_getpwuid_r(uid C.__uid_t, pwd *C.struct_passwd, buffer *C.char, bufsize C.size_t, result **C.struct_passwd) C.int {
	fmt.Printf("Got request for UID %d\n\n", uid)
	return NssStatusNotFound
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
