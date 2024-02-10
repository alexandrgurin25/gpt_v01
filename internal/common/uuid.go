package common

import (
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

// UUIDFromString преобразует строку в объект UUID.
func UUIDFromString(source string) (*pgtype.UUID, error) {
	if len(source) != 36 {
		return nil, errors.New("cannot parse UUID")
	}

	// Извлекаем шестнадцатеричные значения из строки UUID и декодируем их
	source = source[0:8] + source[9:13] + source[14:18] + source[19:23] + source[24:]
	buf, err := hex.DecodeString(source)
	var buf16 [16]byte
	copy(buf16[:], buf)

	if err != nil {
		return nil, nil
	}

	// Возвращаем объект UUID с декодированными значениями
	return &pgtype.UUID{
		Bytes: buf16,
		Valid: true,
	}, nil
}

// StringFromUUID преобразует объект UUID в строку.
func StringFromUUID(source *pgtype.UUID) string {
	// Форматируем байты UUID в строку формата "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
	return fmt.Sprintf("%x-%x-%x-%x-%x", source.Bytes[0:4], source.Bytes[4:6], source.Bytes[6:8], source.Bytes[8:10], source.Bytes[10:16])
}
