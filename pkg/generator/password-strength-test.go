package generator

import(
	"strings"
)

func passwordstren(password string){ 
	var strength string 
	var length int
	var hascap int
	var hasspec int
	var hasnum int 
	var hasnodicword int 
	
	if(len(password) > 8){
		length = 1
	}else { 
		length = 0
	}

	for i := 0; i < len(password); i ++ { 
		if(password[i].iscapital()){//psuedo my boy
				hascap = 1;

			}
		if(password[i].isspecial()){
			hasspec = 1
		}
		if(password[i].isnum()){
			hasnum = 1
		}
	}

	if(length + hascap + hasspec + hasnum + hasnodicword == 4){
		strength = "your password is decent"
	}
	if(length + hascap + hasspec + hasnum +hasnodicword <= 3){
		strength = "your password sucks figure it out"
	}
	if(length + hascap + hasspec + hasnum + hasnodicword == 5 ){
		strength = "eyyyyy lets get it your password is strong"
	}
	return strength
}