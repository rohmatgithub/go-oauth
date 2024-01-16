package service

import "go-oauth/common"

func InsertI18NMessage(language string) string {
	return GenerateI18NMessage("SUCCESS_INSERT_MESSAGE", language)
}

func UpdateI18NMessage(language string) string {
	return GenerateI18NMessage("SUCCESS_UPDATE_MESSAGE", language)
}

func ListI18NMessage(language string) string {
	return GenerateI18NMessage("SUCCESS_LIST_MESSAGE", language)
}

func CountI18NMessage(language string) string {
	return GenerateI18NMessage("SUCCESS_COUNT_MESSAGE", language)
}

func DeleteI18NMessage(language string) string {
	return GenerateI18NMessage("SUCCESS_DELETE_MESSAGE", language)
}

func ViewI18NMessage(language string) string {
	return GenerateI18NMessage("SUCCESS_VIEW_MESSAGE", language)
}
func GenerateI18NMessage(messageID string, language string) (output string) {
	return common.GenerateI18NServiceMessage(common.CommonBundle, messageID, language, nil)
}
