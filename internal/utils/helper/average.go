package helper

import "terhandle/internal/app/model"



func GetAverageRating(feedbacks []model.Feedback) float64 {
    totalRating := 0
    totalFeedbacks := len(feedbacks)

    for _, feedback := range feedbacks {
        totalRating += feedback.Rating
    }

    averageRating := 0.0
    if totalFeedbacks > 0 {
        averageRating = float64(totalRating) / float64(totalFeedbacks)
    }

    return averageRating
}
