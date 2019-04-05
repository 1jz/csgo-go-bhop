package lib

import (
	"bytes"
	"unsafe"
)

func GetProcessID(process string) (uint32, bool) {
	//https://docs.microsoft.com/en-us/windows/desktop/toolhelp/taking-a-snapshot-and-viewing-processes

	var snap HANDLE
	var pe32 PROCESSENTRY32

	// Take a snapshot of all processes in the system.
	snap = CreateToolhelp32Snapshot(TH32CS_SNAPALL, 0)
	pe32.DwSize = uint32(unsafe.Sizeof(pe32))

	//check first process
	exit := Process32First(snap, &pe32)
	parsed := parseint8(pe32.SzExeFile[:])
	if !exit {
		CloseHandle(snap) // clean the snapshot object
		return 0, false
	} else {
		for i := true; i; i = Process32Next(snap, &pe32) {
			parsed = parseint8(pe32.SzExeFile[:])
			if parsed == process {
				//fmt.Println(parsed, pe32.Th32ProcessID)

				return pe32.Th32ProcessID, true
			} else {
				//fmt.Println(parsed)
			}
		}
	}
	return 0, false
}

func GetModule(module string, PID uint32) (MODULEENTRY32, bool, unsafe.Pointer) {
	var me32 MODULEENTRY32
	var snap HANDLE

	snap = CreateToolhelp32Snapshot(TH32CS_SNAPMODULE|TH32CS_SNAPMODULE32, PID)

	me32.DwSize = uint32(unsafe.Sizeof(me32))

	//check first module
	exit := Module32First(snap, &me32)
	parsed := parseint8(me32.SzModule[:])
	if !exit {
		CloseHandle(snap) // clean the snapshot object
		return me32, false, unsafe.Pointer(me32.ModBaseAddr)
	} else {
		for i := true; i; i = Module32Next(snap, &me32) {
			parsed = parseint8(me32.SzModule[:])
			if parsed != module {
				//fmt.Println(parsed)
			} else {
				//fmt.Println(parsed)
				return me32, true, unsafe.Pointer(me32.ModBaseAddr)
			}
		}
	}
	return me32, false, unsafe.Pointer(me32.ModBaseAddr)
}

func parseint8(arr []uint8) string {
	n := bytes.Index(arr, []uint8{0})
	return string(arr[:n])
}

func OffsetAddr(hProcess HANDLE, baseAddr uintptr, offAddrs []uintptr) uintptr {
	var finalAddr uintptr

	return finalAddr
}
