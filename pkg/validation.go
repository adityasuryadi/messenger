package pkg

import (
	"context"
	"errors"
	"log/slog"
	"reflect"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Validation struct {
	validation *validator.Validate
}

type User struct {
	Email    string `bson:"email"`
	Fullname string `bson:"fullname"`
}

func NewValidation(db *mongo.Database) *Validation {
	validation := validator.New()

	// register custom unique validation
	validation.RegisterValidation("unique", func(fl validator.FieldLevel) bool {
		// // get parameter dari tag struct validate
		table := fl.Param()
		field := strings.ToLower(fl.StructFieldName())
		var exist bool
		var result User
		// var result struct {
		// 	Value float64
		// }
		// err := db.Table(table).Where(""+field+" = ?", fl.Field().String()).Count(&total).Error
		err := db.Collection(table).FindOne(context.TODO(), bson.M{field: fl.Field().String()}).Decode(&result)
		if errors.Is(err, mongo.ErrNoDocuments) {
			exist = false
		} else if err == nil {
			exist = true
		} else if err != nil {
			slog.Error("failed to get data", err)
		}

		// // Return true if the count is zero (i.e., the value is unique)
		return !exist
	})

	// validasi gte than now
	validation.RegisterValidation("gtenow", func(fl validator.FieldLevel) bool {
		bookingDate, _ := time.Parse("2006-01-02", fl.Field().String())
		now := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC)
		diff := bookingDate.Sub(now)
		return diff >= 0
	})

	validation.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		// skip if tag key says it should be ignored
		if name == "-" {
			return ""
		}
		return name
	})

	return &Validation{
		validation: validation,
	}
}

func removeFirstNameSpace(namespace string) string {
	s := strings.Split(namespace, ".")
	if len(s) > 1 {
		arr := make([]string, 0, len(s))
		for i := 1; i < len(s); i++ {
			arr = append(arr, s[i])
		}
		result := strings.Join([]string(arr), ".")
		return result
	}
	return namespace
}

func GetErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "field " + fe.Field() + " tidak boleh kosong"
	case "lte":
		return "harus lebih kecil dari " + fe.Param()
	case "gtenow":
		return "harus lebih besar dari tanggal hari ini"
	case "gte":
		return "harus lebih besar dari " + fe.Param()
	case "email":
		return "format email salah"
	case "unique":
		return fe.Field() + " not avaiable"
	case "min":
		return "minimal " + fe.Param() + " karakter"
	case "max":
		return "max " + fe.Param() + " Kb"
	case "image_validation":
		return "Harus Image"
	case "number":
		return "harus numeric"
	case "eqfield":
		return "field tidak sama dengan " + fe.Param()
	case "maxquota":
		return "kuota antrian habis"
	case "maxquotabooking":
		return "kuota antrian habis"
	case "point":
		return "point tidak mencukupi"
	case "datetime":
		return "format waktu salah"
	}
	return "Unknown error"
}

func (v *Validation) ValidateRequest(request interface{}) error {
	err := v.validation.Struct(request)
	if err != nil {
		return err
	}
	return nil
}

func (v *Validation) ErrorJson(err error) interface{} {

	validationErrors := err.(validator.ValidationErrors)
	out := make(map[string][]string, len(validationErrors))
	for _, fieldError := range validationErrors {
		out[removeFirstNameSpace(fieldError.Namespace())] = append(out[removeFirstNameSpace(fieldError.Namespace())], GetErrorMsg(fieldError))
	}
	return out
}
