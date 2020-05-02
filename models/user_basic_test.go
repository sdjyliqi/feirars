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
