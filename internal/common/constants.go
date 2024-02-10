package common

const PasswordHashRound = 14

const QuestionsRateLimitInterval = 3600 * 24

const MaxQuestionCount = 10

var SecretKey = []byte("secret")

const ExpirationTime = 1.5 * 60 * 60 // 1.5 часа в сек.

const SQLTimestampFormatTemplate = "2006-01-02 15:04:05"
