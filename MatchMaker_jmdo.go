package main

import (
	"fmt"
	"strconv"
)

const prefansw1 = 1
const prefansw5 = 5

const weight1 int = 1
const weight2 int = 2
const weight3 int = 3

const trueLove = 90
const maybeFreinds = 78
const runAway = 70

func validate(x string, q1 string) int {
	xi, err := strconv.Atoi(x)
	if err != nil {
		fmt.Println("")
		fmt.Println("Please type an integer.")
		fmt.Println(q1)
		fmt.Scanln(&x)
		xi = validate(x, q1)
	}
	if xi > 5 || xi < 1 {
		fmt.Println("")
		fmt.Println("You must select a number between 1 and 5. Please Try again.")
		fmt.Println(q1)
		fmt.Scanln(&x)
		xi = validate(x, q1)
	}
	return xi
}

func main() {
	fmt.Println("Hello Welcome to Match Maker. In this program I will give you five statements.")
	fmt.Println("You will need to select a number 1 - 5 with one being strongly disagree and 5 being strongly agree.")
	fmt.Println("After that, I will give you a compatability score to see if we are a good match.")
	fmt.Println("Good luck!")

	var q1 string = "Soccer is the best sport in the world."
	fmt.Println("")
	fmt.Println(q1)
	var answer1 string
	var answer1i int
	var question1_comp int
	fmt.Scanln(&answer1)
	answer1i = validate(answer1, q1)
	question1_comp = (prefansw5 - answer1i)
	if question1_comp < 0 {
		question1_comp = question1_comp * -1
	}
	var q1_weigthed_score int
	q1_weigthed_score = question1_comp * weight3

	var q2 string = "Going to the gym is great and a necessity in my schedule."
	fmt.Println("")
	fmt.Println(q2)
	var answer2 string
	var question2_comp int
	var answer2i int
	fmt.Scanln(&answer2)
	answer2i = validate(answer2, q2)
	question2_comp = (prefansw5 - answer2i)
	if question2_comp < 0 {
		question2_comp = question2_comp * -1
	}
	var q2_weigthed_score int
	q2_weigthed_score = question2_comp * weight2

	var q3 string = "Traveling the world is not a necessity for me."
	fmt.Println("")
	fmt.Println(q3)
	var answer3 string
	var question3_comp int
	var answer3i int
	fmt.Scanln(&answer3)
	answer3i = validate(answer3, q3)
	question3_comp = (prefansw1 - answer3i)
	if question3_comp < 0 {
		question3_comp = question3_comp * -1
	}
	var q3_weigthed_score int
	q3_weigthed_score = question3_comp * weight3

	var q4 string = "Living in Europe sounds like an amazing idea."
	fmt.Println("")
	fmt.Println(q4)
	var answer4 string
	var question4_comp int
	var answer4i int
	fmt.Scanln(&answer4)
	answer4i = validate(answer4, q4)
	question4_comp = (prefansw5 - answer4i)
	if question4_comp < 0 {
		question4_comp = question4_comp * -1
	}
	var q4_weigthed_score int
	q4_weigthed_score = question4_comp * weight2

	var q5 string = "Playing Video Games is childish and a waste of time."
	fmt.Println("")
	fmt.Println(q5)
	var answer5 string
	var question5_comp int
	var answer5i int
	fmt.Scanln(&answer5)
	answer5i = validate(answer5, q5)
	question5_comp = (prefansw1 - answer5i)
	if question5_comp < 0 {
		question5_comp = question5_comp * -1
	}
	var q5_weigthed_score int
	q5_weigthed_score = question5_comp * weight1

	var totalComp int = question1_comp + question2_comp + question3_comp + question4_comp + question5_comp
	var final_comp int = 100 - q1_weigthed_score - q2_weigthed_score - q3_weigthed_score - q4_weigthed_score - q5_weigthed_score

	fmt.Println("")
	fmt.Println("Your results are in...")
	fmt.Println("Your Question 1 compatibility score is: ", question1_comp, ", and the weigthed score is: ", q1_weigthed_score)
	fmt.Println("Your Question 2 compatibility score is: ", question2_comp, ", and the weigthed score is: ", q2_weigthed_score)
	fmt.Println("Your Question 3 compatibility score is: ", question3_comp, ", and the weigthed score is: ", q3_weigthed_score)
	fmt.Println("Your Question 4 compatibility score is: ", question1_comp, ", and the weigthed score is: ", q4_weigthed_score)
	fmt.Println("Your Question 5 compatibility score is: ", question1_comp, ", and the weigthed score is: ", q5_weigthed_score)
	fmt.Println("")
	fmt.Println("Total Compatibility Score is: ", totalComp)
	fmt.Println("Your final weigthed compatibility score is: ", final_comp, "%")

	if final_comp >= trueLove {
		fmt.Println("")
		fmt.Println("Congratualtions! we are a great match! Looks like you are my true love :))).")
	} else if final_comp >= maybeFreinds && final_comp < trueLove {
		fmt.Println("")
		fmt.Println("We might not be the best match, but we can maybe be friends :).")
	} else if final_comp < maybeFreinds || final_comp <= runAway {
		fmt.Println("")
		fmt.Println("I am sorry but we are not a match. I will run away now.")
	}

}
