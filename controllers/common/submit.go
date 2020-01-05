package common

import (
	"Rabbit-OJ-Backend/controllers/auth"
	"Rabbit-OJ-Backend/models/forms"
	"Rabbit-OJ-Backend/services/question"
	SubmissionService "Rabbit-OJ-Backend/services/submission"
	"Rabbit-OJ-Backend/services/user"
	"Rabbit-OJ-Backend/utils/files"
	"errors"
	"fmt"
	"io/ioutil"
)

func CodeSubmit(tid string, submitForm *forms.SubmitForm, authObject *auth.Claims, isContest bool) (string, error) {
	questionJudge, err := question.JudgeInfo(tid)
	if err != nil {
		return "", err
	}

	questionDetail, err := question.Detail(tid)
	if err != nil {
		return "", err
	}

	if questionDetail.Hide && !authObject.IsAdmin {
		return "", errors.New("permission denied")
	}

	fileName, err := files.CodeGenerateFileNameWithMkdir(authObject.Uid)
	if err != nil {
		return "", err
	}

	filePath, err := files.CodePath(fileName)
	if err != nil {
		return "", err
	}

	if err := ioutil.WriteFile(filePath, []byte(submitForm.Code), 0644); err != nil {
		return "", err
	}

	submission, err := SubmissionService.Create(tid, authObject.Uid, submitForm.Language, fileName)
	if err != nil {
		return "", err
	}

	go func() {
		if err := SubmissionService.Starter(
			[]byte(submitForm.Code), submission, questionJudge,
			questionDetail,
			isContest); err != nil {
			fmt.Print(err)
		}
		question.UpdateAttemptCount(tid)
		user.UpdateAttemptCount(authObject.Uid)
	}()

	return submission.Sid, nil
}
