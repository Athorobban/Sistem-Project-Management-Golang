package types

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type UUIDArray []uuid.UUID

func (a *UUIDArray) Scan(value interface{}) error {

	//{213123asdad,asdasda,123123}
	var str string

	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	default:
		return errors.New("failed to parse UUIDArray: unsupport data type")
	}

	str = strings.TrimPrefix(str, "{")
	str = strings.TrimSuffix(str, "}")
	parts := strings.Split(str, ",")

	//make([]T,length , capacity)
	*a = make(UUIDArray, 0, len(parts))
	for _, s := range parts {
		s = strings.TrimSpace(strings.Trim(s, `"`)) // akan menghapus spasi dan "
		if s == "" {
			continue
		}
		u, err := uuid.Parse(s)
		if err != nil {
			return fmt.Errorf("invalid UUID in Array : %v", err)
		}
		*a = append(*a, u)

	}
	return nil

}

//{"550e8400-e29b-41d4-a716-446655440000","123e4567-e89b-12d3-a456-426614174000"}

func (a UUIDArray) Value() (driver.Value, error) {
	if len(a) == 0 {
		return "{}", nil
	}
	postgreFormat := make([]string, 0, len(a))
	for _, value := range a {
		postgreFormat = append(postgreFormat, fmt.Sprintf(`"%s"`, value.String()))
	}
	return "{" + strings.Join(postgreFormat, ",") + "}", nil

}

func (UUIDArray) GormDataType() string {
	return "uuid[]"
}