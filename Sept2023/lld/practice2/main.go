package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/vipulvpatil/practice2/command"
	"github.com/vipulvpatil/practice2/custom_errors"
	"github.com/vipulvpatil/practice2/data"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	eh := &custom_errors.ErrorHandler{}
	users := []*data.User{
		data.CreateUser("u1", "user1", "user1@sample.com", "1234567890"),
		data.CreateUser("u2", "user2", "user2@sample.com", "1234567890"),
		data.CreateUser("u3", "user3", "user3@sample.com", "1234567890"),
		data.CreateUser("u4", "user4", "user4@sample.com", "1234567890"),
	}

	usersMap := make(map[string][]*data.Owes)
	owesList := []*data.Owes{}
	for i := 0; i < len(users); i++ {
		for j := i + 1; j < len(users); j++ {
			owes := data.NewOwes(users[i].Id(), users[j].Id(), 0)
			owesList = append(owesList, &owes)
			usersMap[users[i].Id()] = append(usersMap[users[i].Id()], &owes)
			usersMap[users[j].Id()] = append(usersMap[users[j].Id()], &owes)
		}
	}

	for scanner.Scan() {
		commandText := scanner.Text()
		if commandText == "SHOW" {
			for _, o := range owesList {
				fmt.Println(o)
			}
			continue
		}

		owes := command.CreateOwesFromCommand(commandText, eh)
		for _, o := range owes {
			src := o.Src()
			for _, o2 := range usersMap[src] {
				data.MergeInto(o2, &o)
			}
		}
	}
}
