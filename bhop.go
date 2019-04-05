package main

import (
	"fmt"
	"time"
	"unsafe"

	"./lib"
)

func main() {
	var (
		procName      = "csgo.exe"
		modName       = "client_panorama.dll"
		jump          = uintptr(0x6)
		dwLocalPlayer uintptr
		dwForceJump   uintptr
		oFlag         uintptr
		holder        uint32
		baseP         uintptr
	)

	dwLocalPlayer = 0x00CD2764
	oFlag = 0x104
	dwForceJump = 0x5186978

	proc, success := lib.GetProcessID(procName)

	if success == true {

		fmt.Printf("Success: PID of %s is %d\n", procName, proc)

		mod, success, addr := lib.GetModule(modName, proc)

		if success == true {

			base := uintptr(unsafe.Pointer(addr))
			dwForceJump = dwForceJump + base

			fmt.Println("Success: addr of", modName, "is", addr)

			process, err := lib.OpenProcess(lib.PROCESS_ALL_ACCESS, false, proc)

			fmt.Println("===============================")
			lib.ReadProcessMemory(process, lib.LPCVOID(base+dwLocalPlayer), &baseP, unsafe.Sizeof(holder))
			prAddr := fmt.Sprintf("%X", baseP)
			fmt.Println("Found local player pointer at:", prAddr)
			fmt.Println("===============================")

			if err == nil {
				fmt.Println(err)
			} else {

				var buffer uintptr

				for {
					if lib.GetAsyncKeyState(32) > 0 {

						lib.ReadProcessMemory(process, lib.LPCVOID(baseP+oFlag), &buffer, 1)

						if buffer == 1 || buffer == 7 {
							lib.WriteProcessMemory(process, dwForceJump, unsafe.Pointer(&jump), 1)
						}
					}
					time.Sleep(1)
				}
			}
		} else {
			fmt.Printf("Error: %s module not found. returned: \"%d\"\n", modName, mod)
		}
	} else {
		fmt.Printf("Error: %s process not found. returned: \"%d\"\n", procName, proc)
	}
}
