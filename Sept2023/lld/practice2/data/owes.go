package data

import "fmt"

type Owes struct {
	// srcUser owes destUser amount if amount positive
	// destUser owes srcUser -amount if amount negative
	srcUser  string
	destUser string
	amount   float64
}

func NewOwes(srcUser string, destUser string, amount float64) Owes {
	return Owes{
		srcUser:  srcUser,
		destUser: destUser,
		amount:   amount,
	}
}

func MergeInto(owes1, owes2 *Owes) {
	if owes1.srcUser == owes2.srcUser && owes1.destUser == owes2.destUser {
		owes1.amount = owes1.amount + owes2.amount
	} else if owes1.srcUser == owes2.destUser && owes1.destUser == owes2.srcUser {
		owes1.amount = owes1.amount - owes2.amount
	}
}

func (o *Owes) Src() string {
	return o.srcUser
}

func (o Owes) String() string {
	if o.amount > 0 {
		return fmt.Sprintf("%s owes %s: %.2f", o.srcUser, o.destUser, o.amount)
	} else if o.amount < 0 {
		return fmt.Sprintf("%s owes %s: %.2f", o.destUser, o.srcUser, -1*o.amount)
	}
	return ""
}
