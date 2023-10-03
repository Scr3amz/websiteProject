package utils

import "github.com/Scr3amz/websiteProject/pkg/models"

/* Структура-оболочка для передачи сразу нескольких динамических элементов*/
type templateData struct {
	Note *models.Note
	Notes []*models.Note
}

func CreateData() *templateData {
	 return &templateData{}
}