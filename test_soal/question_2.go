package test_soal

func QuestionNumber2(str string) string {
	// Without this, the program generates slice bound error
	if len(str) <= 1 {
		return str
	}
	return QuestionNumber2(string(str[1:])) + string(str[0])
}