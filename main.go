package main

import "fmt"

func isPushOpcode(opcode string) (bool, int) {
	if opcode >= "60" && opcode <= "7f" {
		// PUSH1(0x60)到PUSH32(0x7f)
		distance := int(opcode[1]) - '0' // 计算基础距离
		if opcode[1] > '9' {             // 对于字母a-f
			distance = 10 + int(opcode[1]-'a')
		}
		return true, distance + 1 // 返回是否为PUSH指令和应跳过的字节（数据）数
	}
	return false, 0
}

func main() {
	bytecode := "608060405234801561001057600080fd5b50600436106100365760003560e01c8063d1d80fdf1461003b578063f8a8fd6d1461006d575b600080fd5b61006b6100493660046100fd565b600080546001600160a01b0319166001600160a01b0392909216919091179055565b005b61006b6000805460408051600481526024810182526020810180516001600160e01b03166303155a6760e21b17905290516001600160a01b0390921692620f42409290916100ba9161012d565b600060405180830381858888f193505050503d80600081146100f8576040519150601f19603f3d011682016040523d82523d6000602084013e505050565b505050565b60006020828403121561010f57600080fd5b81356001600160a01b038116811461012657600080fd5b9392505050565b6000825160005b8181101561014e5760208186018101518583015201610134565b8181111561015d576000828501525b50919091019291505056fea2646970667358221220cc7d6ddf3143c42ef5bc0ccff7f77ef357b677f50a6cca00d9644047be61274b64736f6c63430008070033"
	fmt.Println("原始字节码：", bytecode)
	fmt.Println("字节位置\tOpcode\t\t数据")

	for i := 0; i < len(bytecode); {
		op := bytecode[i : i+2]
		if opcode, exists := opcodes[op]; exists {
			isPush, bytesToPush := isPushOpcode(op)
			if isPush {
				data := ""
				if i+2+bytesToPush*2 <= len(bytecode) {
					data = bytecode[i+2 : i+2+bytesToPush*2]
					fmt.Printf("%x\t\t%s\t\t%s\n", i/2, opcode, data)
					i += 2 + bytesToPush*2 // 纠正了此处的逻辑，正确跳过数据字节
					continue
				}
			}
			fmt.Printf("%x\t\t%s\n", i/2, opcode)
			i += 2
		} else {
			fmt.Printf("%x\t\tUNKNOWN\n", i/2)
			i += 2 // 对于未知指令，同样递增i以避免死循环
		}
	}
}
