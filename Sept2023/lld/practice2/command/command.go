package command

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/vipulvpatil/practice2/custom_errors"
	"github.com/vipulvpatil/practice2/data"
)

func CreateOwesFromCommand(commandString string, eh *custom_errors.ErrorHandler) []data.Owes {
	parts := strings.Split(commandString, " ")
	parts = parts[1:]
	if len(parts) < 6 {
		eh.HandleError(errors.New("invalid command"))
		return nil
	}
	payer := parts[0]
	paidAmount, err := strconv.ParseFloat(parts[1], 64)
	if err != nil {
		eh.HandleError(err)
		return nil
	}
	splitNo, err := strconv.Atoi(parts[2])
	if err != nil {
		eh.HandleError(err)
		return nil
	}
	owers := []string{}
	i := 3
	for ; i < 3+splitNo; i++ {
		owers = append(owers, parts[i])
	}

	splitStyle := parts[i]
	owedAmount := calculateOwedAmount(splitStyle, paidAmount, splitNo, parts[i+1:])

	if len(owedAmount) != len(owers) {
		eh.HandleError(errors.New("something is wrong"))
		return nil
	}

	fmt.Println(payer, owers, owedAmount)
	owes := []data.Owes{}
	for i := range owers {
		owes = append(owes, data.NewOwes(owers[i], payer, owedAmount[i]))
	}
	return owes
}

func calculateOwedAmount(splitStyle string, paidAmount float64, splitNo int, args []string) []float64 {
	owed := []float64{}
	switch splitStyle {
	case "EQUAL":
		each := paidAmount / float64(splitNo)
		eachRoundedDown := math.Floor(each*100) / 100
		for i := 0; i < splitNo; i++ {
			owed = append(owed, eachRoundedDown)
		}
		sum := eachRoundedDown * float64(splitNo)
		if sum < paidAmount {
			owed[0] = owed[0] + paidAmount - sum
			owed[0] = math.Round(owed[0]*100) / 100
		}
	case "PERCENT":
		percentSum := 0.0
		sum := 0.0
		for i := range args {
			p, err := strconv.ParseFloat(args[i], 64)
			if err == nil {
				roundedDown := math.Floor(paidAmount*p) / 100
				owed = append(owed, roundedDown)
				percentSum = percentSum + p
				sum = sum + roundedDown
			}
		}
		if percentSum != 100 {
			return nil
		}
		if sum < paidAmount {
			owed[0] = owed[0] + paidAmount - sum
			owed[0] = math.Round(owed[0]*100) / 100
		}
	case "EXACT":
		sum := 0.0
		for i := range args {
			a, err := strconv.ParseFloat(args[i], 64)
			if err == nil {
				owed = append(owed, a)
				sum = sum + a
			}
		}
		if sum != paidAmount {
			return nil
		}
	}

	return owed
}
