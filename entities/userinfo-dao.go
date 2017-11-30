package entities

func Save(u *UserInfo) error {
	session := engine.NewSession()
	defer session.Close()
	err := session.Begin()
	if err != nil {
		return err
	}
	_, err = session.Table("userinfo").Insert(u)
	if err != nil {
		session.Rollback()
		return err
	}
	err = session.Commit()
	if err != nil {
		return err
	}
	return nil
}

func FindAll() []UserInfo {
	ulist := make([]UserInfo, 0)
	err := engine.Table("userinfo").Find(&ulist)
	checkErr(err)
	return ulist
}

func FindByID(id int) *UserInfo {
	user := new(UserInfo)
	_, err := engine.Table("userinfo").Where("uid=?", id).Get(user)
	checkErr(err)
	return
