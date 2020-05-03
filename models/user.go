package models

import (
	"fmt"
	"errors"
)


type User struct {
	Id int64
	Nick_name string
	Invite_code string
	Password string
	Balance float64
	Verify int
}

func AddUser(nick_name,invite_code,pwd string) (error) {
	if _,exit := userExit(invite_code);exit {
		return errors.New("user exit")
	}

	user := User{Nick_name:nick_name,Invite_code:invite_code,Password:pwd}
	id,err := Omer.Insert(&user)

	if err ==nil && id > -1 {
		return nil
	}

	return errors.New("user create failed")
}


func userExit(invite_code string)(error,bool)  {
	user := User{}
	err := Omer.QueryTable("user").Filter("invite_code",invite_code).One(&user)

	//err ！= nil表示有错，用户此时不存在，== nil表示用户存在
	return err,err == nil
}

func UpdateUser(user *User) (error) {
	_,err := Omer.Update(user)
	return err
}

func QueryUserById(id int64)(User,error)  {
	user := User{Id:id}
	err := Omer.Read(&user)
	return user,err
}

func UserInfo(id int64) (*User ,error) {
	user := User{Id:id}
	err := Omer.Read(&user)
	return &user,err
}

func Users() (* []User) {
	users := make([]User, 0)
	_,err := Omer.Raw("select * from user").QueryRows(&users)
	if err == nil {
		return &users;
	}

	return  nil;
	
}

func init() {
	fmt.Println("user init")
}
