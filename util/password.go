package util

import (
	"fmt"
	"regexp"
)

// CheckPasswordStrenth 检查密码强度
func CheckPasswordStrenth(ps string) error {
	if len(ps) < 8 {
		return fmt.Errorf("密码长度不能小于8位")
	}
	num := `[0-9]{1}`
	a_z := `[a-z]{1}`
	A_Z := `[A-Z]{1}`
	symbol := `[!@#~$%^&*()+|_]{1}`
	if b, err := regexp.MatchString(num, ps); !b || err != nil {
		return fmt.Errorf("密码必须含有数字 :%v", err)
	}
	if b, err := regexp.MatchString(a_z, ps); !b || err != nil {
		return fmt.Errorf("密码必须含有小写字母 :%v", err)
	}
	if b, err := regexp.MatchString(A_Z, ps); !b || err != nil {
		return fmt.Errorf("密码必须含有大写字母 :%v", err)
	}
	if b, err := regexp.MatchString(symbol, ps); !b || err != nil {
		return fmt.Errorf("密码必须含有特殊符号 :%v", err)
	}

	return nil
}
