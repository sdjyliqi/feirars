package control

import "errors"

func (uc *userCenter) Login(userID, passport string) error {
	isValid, err := uc.UserBasic.ChkPassportValid(uc.db, userID, passport)
	if err != nil {
		return err
	}
	if !isValid {
		return errors.New("passport-invalid")
	}
	return nil
}

func (uc *userCenter) Logout(userID string) error {
	return nil
}
