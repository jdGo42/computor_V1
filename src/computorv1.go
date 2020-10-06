/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   computorv1.go                                      :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: jdGo42 <jdGo42@student.42.fr>              +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2020/04/14 21:32:31 by jdGo42            #+#    #+#             */
/*   Updated: 2020/10/06 05:38:10 by jdGo42           ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

main

import (
	"os"		// for Args
	"fmt"		// for Printf Println
	"strings"	// for Split Contains etc.
	"strconv"	// for Atoi
	"regexp"	// for parsing
)

func main () {
	verbose := 0
	if len(os.Args) == 2{
		arg := os.Args[1]
		sliceStr, err, lenSlice := regex(arg, verbose)
		copiedArrayStr := make([]string, lenSlice + 1)
		copy(copiedArrayStr[:], sliceStr)
		arrayFloatReduced := groupNumbers(copiedArrayStr, lenSlice)
		if err == nil{
			printReducedForm(arrayFloatReduced)
			solution(arrayFloatReduced)
		} else {
			fmt.Printf("%v", err)
		}
	} else {
		fmt.Println("Error, please enter a unique string between \"\" to execute program")
	}
	return
}

func solution (arrayFloatReduced [3]float64) (discriminant float64) {
	discriminant = (arrayFloatReduced[1] * arrayFloatReduced[1]) - (4 * arrayFloatReduced[0] * arrayFloatReduced[2])
	if arrayFloatReduced[2] != 0 && discriminant > 0 {
		squareDiscriminant := calculSquare(discriminant)
		fmt.Printf("disc : %g, square discr : %g, sqDis *sqDis : %g\n", discriminant, squareDiscriminant, squareDiscriminant * squareDiscriminant)
		fmt.Printf("Polynomial degree: 2\nDiscriminant is strictly positive(%g), the two solutions are:\n%g\n%g\n", discriminant, (-arrayFloatReduced[1] - squareDiscriminant)/ (2 *  arrayFloatReduced[2]), (-arrayFloatReduced[1] + squareDiscriminant)/ (2 *  arrayFloatReduced[2]))
		fmt.Printf("Which correspond to :(-%g - √(%g)) / (2 * %g) and\n(-%g + √(%g)) / (2 * %g)\n", arrayFloatReduced[1], discriminant, arrayFloatReduced[2], arrayFloatReduced[1], discriminant, arrayFloatReduced[2])
	} else if arrayFloatReduced[2] != 0 && discriminant < 0 {
		fmt.Printf("Polynomial degree: 2\nDiscriminant is strictly negative(%g), the two solutions (into the comlex set) are:\n(-%g - (i *√(%g)) / (%g)\n(-%g + (i *√(%g)) / (%g)\n ",discriminant, arrayFloatReduced[1], -discriminant, 2 * arrayFloatReduced[2], arrayFloatReduced[1], -discriminant, 2 * arrayFloatReduced[2])
		fmt.Printf("Which correspond to :\n(-%g - (i *√(%g)) / (2 * %g)\n(-%g + (i *√(%g)) / (2 * %g)\n ", arrayFloatReduced[1], -discriminant, arrayFloatReduced[2], arrayFloatReduced[1], -discriminant, arrayFloatReduced[2])
	} else if arrayFloatReduced[2] != 0 && discriminant == 0 {
		fmt.Printf("Polynomial degree: 2\nDiscriminant is equal to 0(%g), the unique solution is:\n%g\nWhich correspond to :-%g / (2*%g)\n", discriminant, -arrayFloatReduced[1]/ (2 *  arrayFloatReduced[2]), arrayFloatReduced[1], arrayFloatReduced[2])
	} else if arrayFloatReduced[1] != 0{
		fmt.Printf("Polynomial degree: 1\nThe solution is :\n%g\n", arrayFloatReduced[0] / arrayFloatReduced[1])
	} else if arrayFloatReduced[0] != 0{
		fmt.Printf("Polynomial degree: 0\nThere is no solution\n")
	} else if arrayFloatReduced[0] == 0 {
		fmt.Printf("Polynomial degree: 0\nAll the numbers in the universe and behond are a solution\n")
	}
	return discriminant
}

func regex (arg string, verbose int) (arrayStr []string, err error, lenSlice int) {
	splitWhiteSpace := strings.Fields(arg)
	equalPattern := `(^=$)`
	equalRegex, _ := regexp.Compile(equalPattern)
	simplyPattern := `(^-?[0-9]+$)`
	simplyNumberRegex, _ := regexp.Compile(simplyPattern)
	decimalPattern := `(^-?[0-9]+[.][0-9]+$)`
	decimalNumberRegex, _ := regexp.Compile(decimalPattern)
	multiPattern := `(^\*$)`
	multiRegex, _ := regexp.Compile(multiPattern)
	addOrSubPattern := `(^\+$)|(^\-$)`
	addOrSubRegex, _ := regexp.Compile(addOrSubPattern)
	powerPattern :=  `(^X\^0$)|(^X\^1$)|(^X\^2$)`
	powerRegex, _ := regexp.Compile(powerPattern)
	containsEqualRegex, _ := regexp.Compile(".?=.?")
	splitWhiteSpace = strings.Fields(arg)
	nbEqual, numberCount, multiCount, powerCount := 0, 0, 0, 0
	numberBehondEqual := 0
	for index2, _ := range splitWhiteSpace {
		if verbose == 1 {
			fmt.Printf("\"%s\"\n",splitWhiteSpace[index2])
		}
		validSimplyNumberRegex := simplyNumberRegex.MatchString(splitWhiteSpace[index2])
		validDecimalNumberRegex := decimalNumberRegex.MatchString(splitWhiteSpace[index2])
		validMultiRegex := multiRegex.MatchString(splitWhiteSpace[index2])
		validAddOrSubRegex := addOrSubRegex.MatchString(splitWhiteSpace[index2])
		validPowerRegex := powerRegex.MatchString(splitWhiteSpace[index2])
		validEqualRegex := equalRegex.MatchString(splitWhiteSpace[index2])
		containsEqualNumber := containsEqualRegex.MatchString(splitWhiteSpace[index2])
		if containsEqualNumber {
			nbEqual += 1
		}
		if validSimplyNumberRegex || validDecimalNumberRegex|| validMultiRegex|| validAddOrSubRegex|| validPowerRegex || validEqualRegex{
			if verbose == 1 {
				fmt.Printf("validRegex for %s, at index %d\n", splitWhiteSpace[index2], index2)
			}
			if validSimplyNumberRegex || validDecimalNumberRegex {
				numberCount +=1
			}
			if validMultiRegex {
				multiCount += 1
			}
			if validPowerRegex {
				powerCount += 1
			}
			if (validDecimalNumberRegex || validSimplyNumberRegex || validMultiRegex || validPowerRegex) && nbEqual !=0 {
				numberBehondEqual += 1
			}
		} else {
			err = fmt.Errorf("Houston, we have a problem, the parameter %s is not correct\n", splitWhiteSpace[index2])
			return nil, err, -1
		}
		lenSlice = index2
	}
	if nbEqual != 1 {
		err =fmt.Errorf("Error, Houston, we have a problem, number of equal is not 1, we have found it %d times\n", nbEqual)
		return nil, err, -1
	}
//	fmt.Printf("numbers behond equal : %d\n", numberBehondEqual)
	if numberBehondEqual < 3  || numberBehondEqual % 3 != 0 {
		err = fmt.Errorf("Error, Houston, we have a problem, there is no valid number behond the equal, we need at least a \" 0 * X^[0/1/2]\" to be correct\n")
		return nil, err, -1
	}
	if powerCount != multiCount || powerCount != numberCount || multiCount != numberCount {
		err = fmt.Errorf("Error, Houston, we have a problem, the entry :\"%s\" is not well formated\n", arg)
		return nil, err, -1
	}
//	}
	fmt.Println(splitWhiteSpace)
	return splitWhiteSpace, nil, lenSlice
}

func getValue(arrayStr []string, index int, indexEqual int, firstValue int) (value float64){
	if firstValue == 1 && indexEqual == -1 {
		if arrayStr[0] == "-"{
			value, _ = strconv.ParseFloat(arrayStr[1], 64)
			return -value
		} else {
			value, _  = strconv.ParseFloat(arrayStr[0], 64)
			return value
		}
	} else if firstValue == 1 && indexEqual != -1{
		if arrayStr[indexEqual + 1] == "-"{
			value, _ = strconv.ParseFloat(arrayStr[indexEqual + 2], 64)
			return value
		} else {
			value, _  = strconv.ParseFloat(arrayStr[indexEqual + 1], 64)
			return -value
		} 
	} else if firstValue == 0 && indexEqual == -1 {
		if arrayStr[index - 3] == "-"{
			value, _ = strconv.ParseFloat(arrayStr[index - 2], 64)
			return -value
		} else {
			value, _  = strconv.ParseFloat(arrayStr[index - 2], 64)
			return value
		} 
	} else if firstValue == 0 && indexEqual != -1 {
		if arrayStr[index - 3] == "-"{
			value, _ = strconv.ParseFloat(arrayStr[index - 2], 64)
			return value
		} else {
			value, _  = strconv.ParseFloat(arrayStr[index - 2], 64)
			return -value
		} 
	} else{
		return -1000
	} 
}

func groupNumbers(arrayStr []string, lenSlice int) (arrayFloatReduced [3]float64){
	indexEqual := -1
	firstValue := 1
	copiedArrayStr := make([]string, lenSlice + 1)
	copy(copiedArrayStr[:], arrayStr)
	arrayFloatReduced[0] = 0.0
	arrayFloatReduced[1] = 0.0
	arrayFloatReduced[2] = 0.0
	for index, _ := range copiedArrayStr{
		if copiedArrayStr[index] == "=" {
			indexEqual = index
			firstValue = 1
		}
		if copiedArrayStr[index] == "X^0"{
			arrayFloatReduced[0] += getValue(copiedArrayStr, index, indexEqual, firstValue)
			firstValue = 0
		} else 	if copiedArrayStr[index] == "X^1" {
			arrayFloatReduced[1] += getValue(copiedArrayStr, index, indexEqual, firstValue)
			firstValue = 0
		} else 	if copiedArrayStr[index] == "X^2" {
			arrayFloatReduced[2] += getValue(copiedArrayStr, index, indexEqual, firstValue)
			firstValue = 0
		}
	}
	return arrayFloatReduced
}

func printReducedForm(arrayFloatReduced [3]float64) {
	fmt.Printf("Reduced form : ")
	if arrayFloatReduced[0] != 0 {
		fmt.Printf("%g * X^0 ", arrayFloatReduced[0])
	}
	if arrayFloatReduced[1] != 0 {
		if arrayFloatReduced[0] != 0 {
			fmt.Printf("+ %g * X^1 ", arrayFloatReduced[1])
		} else {
			fmt.Printf("%g * X^1 ", arrayFloatReduced[1])
		}
	}
	if arrayFloatReduced[2] != 0 {
		if arrayFloatReduced[0] != 0  || arrayFloatReduced[1] != 0 {
			fmt.Printf("+ %g * X^2 ", arrayFloatReduced[2])
		} else {
			fmt.Printf("%g * X^2 ", arrayFloatReduced[2])
		}
	} 
	if arrayFloatReduced[0] == 0 && arrayFloatReduced[1] == 0 && arrayFloatReduced[2] == 0 {
		fmt.Printf("0")
	}
	fmt.Printf(" = 0\n")
}

func calculSquare(number float64) (squareNumber float64) {
	// this is the heron algorythm
	// another method possible is to implement Newton algorythm : calcul derivative, take a random number x, make a for loop in which x = x - f(x)/df(x)
	toDecreaseNumber := number
	squareNumber = number
	for increment := 0; increment < 14; increment++{ // 14 to be really close from a perfect precision, sometimes overkill but anyway
		squareNumber = (toDecreaseNumber+(number)/toDecreaseNumber)/2
		toDecreaseNumber = squareNumber
	}
	return squareNumber
}
