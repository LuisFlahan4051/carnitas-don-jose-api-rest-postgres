package commonFunctions

import (
	"bytes"
	"errors"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/LuisFlahan4051/carnitas-don-jose-api-rest-postgres/database"
	"github.com/LuisFlahan4051/carnitas-don-jose-api-rest-postgres/models"
	"github.com/chai2010/webp"
)

const (
	ROOT       = 1
	ADMIN      = 2
	SUPERVISOR = 3
	CASHIER    = 4
	WAITER     = 5
	COOK       = WAITER
)

func valueIsNull(fieldType reflect.Kind, value reflect.Value) bool {
	switch fieldType {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return value.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return value.Float() == 0
	case reflect.Bool:
		return !value.Bool()
	case reflect.String:
		return value.String() == "" || value.String() == "0001-01-01T00:00:00Z" || value.String() == "0001-01-01 00:00:00 +0000 UTC"
	case reflect.Slice:
		return value.Len() == 0
	case reflect.Struct:
		return value.IsZero()
	case reflect.Interface:
		return value.IsZero()
	case reflect.Ptr:
		return value.IsNil()
	}
	return false
}

func getValueField(fieldType reflect.Kind, value reflect.Value) interface{} {
	switch fieldType {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return value.Uint()
	case reflect.Float32, reflect.Float64:
		return value.Float()
	case reflect.Bool:
		return value.Bool()
	case reflect.String:
		return value.String()
	case reflect.Slice:
		return value.Interface()
	case reflect.Struct:
		return value.Interface()
	case reflect.Interface:
		return value.Interface()
	case reflect.ValueOf(time.Time{}).Kind():
		return value.Interface()
	case reflect.Ptr:
		return value.Elem()
	}
	return nil
}

func GetStructFieldsNotNull(anyStruct interface{}) ([]string, map[string]interface{}) {
	var fields []string
	valueFields := make(map[string]interface{})

	_struct := reflect.ValueOf(anyStruct)

	for i := 0; i < _struct.NumField(); i++ {
		_field := _struct.Type().Field(i)
		value := _struct.Field(i)

		fieldType := _field.Type.Kind()

		switch fieldType {
		case reflect.Struct, reflect.Interface:
			subStruct := _struct.Field(i).Interface()
			subFields, subValueFields := GetStructFieldsNotNull(subStruct)
			fields = append(fields, subFields...)
			for subKey, subValue := range subValueFields {
				valueFields[subKey] = subValue
			}
		}

		valueFieldIsNull := valueIsNull(fieldType, value)

		jsonTag := _field.Tag.Get("json")
		if !valueFieldIsNull && jsonTag != "" {
			jsonTagOptions := strings.Split(jsonTag, ",")
			fieldName := jsonTagOptions[0]
			fields = append(fields, fieldName)

			valueField := getValueField(fieldType, value)
			if valueField != nil {
				valueFields[fieldName] = valueField
			}
		}
	}
	return fields, valueFields
}

func GetStructFields(anyStruct interface{}) ([]string, map[string]interface{}) {
	var fields []string
	valueFields := make(map[string]interface{})

	_struct := reflect.ValueOf(anyStruct)
	for i := 0; i < _struct.NumField(); i++ {
		_field := _struct.Type().Field(i)
		value := _struct.Field(i)
		fieldType := _field.Type.Kind()

		switch fieldType {
		case reflect.Struct, reflect.Interface:
			subStruct := _struct.Field(i).Interface()
			subFields, subValueFields := GetStructFields(subStruct)
			fields = append(fields, subFields...)
			for subKey, subValue := range subValueFields {
				valueFields[subKey] = subValue
			}
		}

		jsonTag := _struct.Type().Field(i).Tag.Get("json")
		if jsonTag != "" {
			jsonTagOptions := strings.Split(jsonTag, ",")
			fieldName := jsonTagOptions[0]
			fields = append(fields, fieldName)

			valueField := getValueField(fieldType, value)
			if valueField != nil {
				valueFields[fieldName] = valueField
			}
		}
	}

	return fields, valueFields
}

func GetStructFieldsNotSlices(anyStruct interface{}) ([]string, map[string]interface{}) {
	var fields []string
	valueFields := make(map[string]interface{})

	_struct := reflect.ValueOf(anyStruct)
	for i := 0; i < _struct.NumField(); i++ {
		_field := _struct.Type().Field(i)
		value := _struct.Field(i)
		fieldType := _field.Type.Kind()

		switch fieldType {
		case reflect.Struct, reflect.Interface:
			subStruct := _struct.Field(i).Interface()
			subFields, subValueFields := GetStructFieldsNotSlices(subStruct)
			fields = append(fields, subFields...)
			for subKey, subValue := range subValueFields {
				valueFields[subKey] = subValue
			}
		}

		jsonTag := _struct.Type().Field(i).Tag.Get("json")
		if fieldType != reflect.Slice && jsonTag != "" {
			jsonTagOptions := strings.Split(jsonTag, ",")
			fieldName := jsonTagOptions[0]
			fields = append(fields, fieldName)
			valueField := getValueField(fieldType, value)
			if valueField != nil {
				valueFields[fieldName] = valueField
			}
		}
	}

	return fields, valueFields
}

// --------------------------------------------------------

func SavePictureAsWebp(file io.Reader, filePath string, fileName string) error {
	data, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	imageData, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return err
	}

	var fileConverted *os.File

	if _, err := os.Stat(filePath + fileName); os.IsNotExist(err) {
		os.MkdirAll(filePath, 0777)
		fileConverted, err = os.Create(filePath + fileName)
		if err != nil {
			return err
		}
	}

	err = webp.Encode(fileConverted, imageData, nil)
	if err != nil {
		return err
	}
	return nil
}

func FileIsImage(fileName string) bool {
	extension := filepath.Ext(CleanSpaces(strings.ToLower(fileName)))
	switch extension {
	case ".jpg", ".jpeg", ".png":
		return true
	default:
		return false
	}
}

// --------------------------------------------------------

// validate (request) vars and returns userID, accessLevel, error
func Authentication(request *http.Request, maxAccessLevel uint) (uint, uint, error) {
	if request.URL.Query().Get("user") == "" || request.URL.Query().Get("password") == "" {
		return 0, 0, errors.New("need credentials to access this resource")
	}

	userId, accessLevel, err := ValidateUser(request.URL.Query().Get("user"), request.URL.Query().Get("password"))

	if err != nil {
		return 0, 0, err
	}

	if accessLevel > maxAccessLevel {
		return 0, 0, errors.New("you don't have access to this resource")
	}

	return userId, accessLevel, nil
}

// if user exists returns userID, accessLevel, error
func ValidateUser(input string, password string) (uint, uint, error) {

	key := CleanSpaces(input)
	var query string

	var err error

	query = "SELECT id, password, verified FROM users WHERE username = $1"
	err = errors.New("user does not exist: incorrect username")

	if IsANumber(key) {
		query = "SELECT user.id, user.password, user.verified FROM users user, user_phones phone " +
			"WHERE user.id = phone.user_id AND phone.main = true AND phone.phone = $1"
		err = errors.New("user does not exist: incorrect number phone")
	}

	if IsAMail(key) {
		query = "SELECT user.id, user.password, user.verified FROM users user, user_mails mail " +
			"WHERE user.id = mail.user_id AND mail.main = true AND mail.email = $1"
		err = errors.New("user does not exist: incorrect email")
	}

	var bufferId uint
	var bufferPassword string
	var bufferVerified bool
	db := database.Connect()
	errQuery := db.QueryRow(query, key).Scan(&bufferId, &bufferPassword, &bufferVerified)

	if errQuery != nil {
		return 0, 0, err
	}

	if bufferPassword != password {
		return 0, 0, errors.New("incorrect password")
	}

	if !bufferVerified {
		return 0, 0, errors.New("user is not verified")
	}

	access_level, errRole := GetUserMaxAccessLevel(bufferId)

	return bufferId, access_level, errRole
}

func GetUserMaxAccessLevel(id uint) (uint, error) {

	db := database.Connect()
	defer db.Close()

	var access_level uint
	query := "SELECT role.access_level, inherit.deleted_at FROM roles role, inherit_user_roles inherit, users us" +
		" WHERE us.id = $1 AND us.id = inherit.user_id AND inherit.role_id = role.id"
	rows, err := db.Query(query, id)

	if err != nil {
		return 0, errors.New("user does not have a role")
	}

	for rows.Next() {
		var level uint
		var deleted_at *time.Time
		rows.Scan(&level, &deleted_at)

		if level > access_level && deleted_at == nil {
			access_level = level
		}
	}

	return access_level, nil
}

// Save a log in the database
func SaveServerActionLog(serverLog models.ServerLogs) {
	db := database.Connect()
	defer db.Close()

	query := "INSERT INTO server_logs (transaction, user_id, branch_id, root, created_at) VALUES ($1, $2, $3, $4, $5)"
	_, err := db.Exec(query, serverLog.Transaction, serverLog.UserID, serverLog.BranchID, serverLog.Root, time.Now())

	if err != nil {
		log.Println("Error saving server action log: " + err.Error())
	}

	log.Println(serverLog.Transaction + " Root: " + strconv.FormatBool(*serverLog.Root))
}

// Prints a log error and sends it to the client
func Logcatch(writer http.ResponseWriter, status int, err error) {
	if err != nil {
		log.Println(err.Error())
		writer.WriteHeader(status)
		writer.Write([]byte(err.Error()))
	}
}

// --------------------------------------------------------

func IsANumber(input string) bool {
	_, err := strconv.Atoi(input)
	if err != nil {
		return false
	} else {
		return true
	}
}

func IsAMail(input string) bool {
	if strings.Contains(input, "@") {
		return true
	} else {
		return false
	}
}

func CleanSpaces(stringToClean string) string {
	result := strings.ReplaceAll(stringToClean, " ", "")
	return result
}
