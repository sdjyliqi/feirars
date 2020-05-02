package models

//
//type UserBasic struct {
//	Id       int    `json:"id" xorm:"not null pk autoincr INT(11)"`
//	Passport string `json:"passport" xorm:"VARCHAR(64)"`
//	Phone    string `json:"phone" xorm:"VARCHAR(16)"`
//	Status   int    `json:"status" xorm:"comment('0 未激活状态  1 正常用户状态   2 已禁止登陆状态') TINYINT(4)"`
//	Email    string `json:"email" xorm:"VARCHAR(64)"`
//}
//
//func (t UserBasic) TableName() string {
//	return "user_basic"
//}
//
//func (t UserBasic) GetUserBasicByID(client *xorm.Engine, id int) (*UserBasic, error) {
//	var item *UserBasic
//	_, err := client.Id(id).Get(item)
//	if err != nil {
//		glog.Errorf("[mysql]Get the item for id %+d from table %s failed,err:%+v", id, item.TableName(), err)
//		return nil, err
//	}
//	return item, nil
//}
//
////NewUserBasic ...新插入一条记录
////todo  为了调试方便，先临时存明文，后续选择MD5（passport + salt）形式处理
//func (t UserBasic) NewUserBasic(client *xorm.Engine, phone, passport, email string) (int64, error) {
//	phoneHasExisted, err := t.chkPhoneHasExisted(client, phone)
//	if phoneHasExisted {
//		return 0, utils.UserCenterRegisterPhoneExist
//	}
//	item := &UserBasic{
//		Passport: passport,
//		Phone:    phone,
//		Status:   utils.UserStatusUnregister,
//		Email:    email,
//	}
//	id, err := client.Insert(item)
//	if err != nil {
//		glog.Errorf("[mysql]Insert the item to table %s failed,err:%+v", item, item.TableName(), err)
//		return 0, err
//	}
//	return id, nil
//}
//
////chkPhoneHasExisted ... 判断手机号是否已经绑定
//func (t UserBasic) chkPhoneHasExisted(client *xorm.Engine, phone string) (bool, error) {
//	item := &UserBasic{
//		Phone: phone,
//	}
//	existed, err := client.Get(item)
//	if err != nil {
//		glog.Errorf("[mysql]Get the item  from the table %s failed,err:%+v", item, item.TableName(), err)
//		return false, err
//	}
//	return existed, nil
//}
