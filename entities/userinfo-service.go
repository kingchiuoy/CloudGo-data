package entities

//UserInfoAtomicService .
type UserInfoAtomicService struct{}

//UserInfoService .
var UserInfoService = UserInfoAtomicService{}

// Save .
func (*UserInfoAtomicService) Save(user *UserInfo) error {
	return Save(user)
}

// FindAll .
func (*UserInfoAtomicService) FindAll() []UserInfo {
	return FindAll()
}

// FindByID .
func (*UserInfoAtomicService) FindByID(id int) *UserInfo {
	return FindByID(id)
}
