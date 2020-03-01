package common

import (
	"Rabbit-OJ-Backend/controllers/auth"
	"Rabbit-OJ-Backend/models"
	"Rabbit-OJ-Backend/models/forms"
	"Rabbit-OJ-Backend/services/question"
	SubmissionService "Rabbit-OJ-Backend/services/submission"
	"Rabbit-OJ-Backend/services/user"
	"Rabbit-OJ-Backend/utils/files"
	"errors"
	"fmt"
	"io/ioutil"
)

func CodeSubmit(tid uint32, submitForm *forms.SubmitForm, authObject *auth.Claims, isContest bool) (uint32, error) {
	questionJudge, err := question.JudgeInfo(tid)
	if err != nil {
		return 0, err
	}

	questionDetail, err := question.Detail(tid)
	if err != nil {
		return 0, err
	}

	if !isContest && questionDetail.Hide && !authObject.IsAdmin {
		return 0, errors.New("permission denied")
	}

	fileName, err := files.CodeGenerateFileNameWithMkdir(authObject.Uid)
	if err != nil {
		return 0, err
	}

	filePath, err := files.CodePath(fileName)
	if err != nil {
		return 0, err
	}

	if err := ioutil.WriteFile(filePath, []byte(submitForm.Code), 0644); err != nil {
		return 0, err
	}

	submission, err := SubmissionService.Create(tid, authObject.Uid, submitForm.Language, fileName)
	if err != nil {
		return 0, err
	}

	go func(submission *models.Submission) {
		if err := SubmissionService.Starter(
			[]byte(submitForm.Code), submission, questionJudge,
			questionDetail,
			isContest); err != nil {
			fmt.Print(err)
		}
		question.UpdateAttemptCount(tid)
		user.UpdateAttemptCount(authObject.Uid)
	}(submission)

	return submission.Sid, nil
}
