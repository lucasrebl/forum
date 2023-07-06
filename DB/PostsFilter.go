package forum

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type FilterType int

const (
	ByDate FilterType = iota
	ByTheme
	ByLike
	ByComment
)

func FilterPostsByDate(posts []Post, startDate time.Time, endDate time.Time) []Post {
	filteredPosts := make([]Post, 0)

	for _, post := range posts {
		if post.CreationDate.After(startDate) && post.CreationDate.Before(endDate) {
			filteredPosts = append(filteredPosts, post)
		}
	}
	return filteredPosts
}

func Filter(posts []Post, startDateInput, endDateInput string) []Post {

	//START DATE && END DATE -- GET, SPLIT AND Atoi
	startDateInputParts := strings.Split(startDateInput, "-") //split la chaine -- array de [année, mois, jour]

	fmt.Println(startDateInputParts, len(startDateInputParts))

	endDateInputParts := strings.Split(endDateInput, "-") //split la chaine

	if (len(startDateInputParts) == 3) && (len(endDateInputParts) == 3) {
		startYearInt, _ := strconv.Atoi(startDateInputParts[0])  // ->> l'année en int
		startMonthInt, _ := strconv.Atoi(startDateInputParts[1]) // ->> le mois en int
		startDayInt, _ := strconv.Atoi(startDateInputParts[2])   // ->> le jour en int
		startDate := time.Date(startYearInt, time.Month(startMonthInt), startDayInt, 0, 0, 0, 0, time.UTC)
		fmt.Println("START DATE -> ", startDate)

		endYearInt, _ := strconv.Atoi(endDateInputParts[0])  // ->> l'année en int
		endMonthInt, _ := strconv.Atoi(endDateInputParts[1]) // ->> le mois en int
		endDayInt, _ := strconv.Atoi(endDateInputParts[2])   // ->> le jour en int
		endDate := time.Date(endYearInt, time.Month(endMonthInt), endDayInt, 0, 0, 0, 0, time.UTC)
		fmt.Println("END DATE -> ", endDate)

		fmt.Println("Func Filter -> ", FilterPostsByDate(posts, startDate, endDate))

		return FilterPostsByDate(posts, startDate, endDate)
	}
	//

	return nil
}
