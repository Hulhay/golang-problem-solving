package gradingstudents

import "math"

func GradingStudents(grades []int32) []int32 {
	finalGrades := []int32{}
	for _, grade := range grades {
		rounded := math.Ceil(float64(grade) / float64(5))
		rounded5Time := rounded * 5
		if grade >= 38 {
			if int32(rounded5Time)-grade < 3 {
				finalGrades = append(finalGrades, int32(rounded5Time))
			} else {
				finalGrades = append(finalGrades, grade)
			}
		} else {
			finalGrades = append(finalGrades, grade)
		}
	}
	return finalGrades
}
