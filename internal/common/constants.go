package common

const (
	PasswordHashRound          = 14
	QuestionsRateLimitInterval = 3600 * 24
	MaxQuestionCount           = 15
	ExpirationTime             = 1.5 * 60 * 60 // 1.5 часа в сек.
	SQLTimestampFormatTemplate = "2006-01-02 15:04:05"
	MaxQuestionCountPremium    = 150
)
