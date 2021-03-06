package log

import "fmt"

var Color colorFuncSet

//Black        0;30     Dark Gray     1;30
//Red          0;31     Light Red     1;31
//Green        0;32     Light Green   1;32
//Brown/Orange 0;33     Yellow        1;33
//Blue         0;34     Light Blue    1;34
//Purple       0;35     Light Purple  1;35
//Cyan         0;36     Light Cyan    1;36
//Light Gray   0;37     White         1;37
//
//RED='\033[0;31m'
//NC='\033[0m' # No Color
//printf "I ${RED}love${NC} Stack Overflow\n"

type colorFuncSet struct{}

func (t colorFuncSet) LightPurple(v interface{}) string {
	return fmt.Sprintf("\033[35m%v\033[0m", v)
}

func (t colorFuncSet) LightCyan(v interface{}) string {
	return fmt.Sprintf("\033[35m%v\033[0m", v)
}

func (t colorFuncSet) Red(v interface{}) string {
	return fmt.Sprintf("\033[31m%v\033[0m", v)
}

func (t colorFuncSet) Green(v interface{}) string {
	return fmt.Sprintf("\033[32m%v\033[0m", v)
}

func (t colorFuncSet) Blue(v interface{}) string {
	return fmt.Sprintf("\033[34m%v\033[0m", v)
}

func (t colorFuncSet) Yellow(v interface{}) string {
	return fmt.Sprintf("\033[33m%v\033[0m", v)
}
