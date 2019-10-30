package tool

import "fmt"

func SetBitValue(szTemp byte, nPos byte, nValue int) (byte, error) {
	if nPos < 0 || nPos > 7 {
		return 0, fmt.Errorf("位移位数错误")
	}
	if nValue == 1 {
		szTemp |= uint8(1) << nPos // 将nPos的bit位设置为1，其他位不变
	} else if nValue == 0 {
		szTemp &= ^(uint8(1) << nPos) // 将nPos的bit位设置为0，其他位不变
	} else {
		return 0, fmt.Errorf("只能是0或者1")
	}
	return szTemp, nil
}

func GetBitValue(szTemp byte, nPos byte) (byte, error) {
	if nPos < 0 || nPos > 7 {
		return 0, fmt.Errorf("位移位数错误")
	}
	szTemp = (szTemp << (7 - nPos)) >> 7
	return szTemp, nil
}
