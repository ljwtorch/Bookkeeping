package utils

import "fmt"

type Account struct {
	//声明必须的字段
	//声明一个字段，保存接收用户输入的选项
	key string
	//声明一个字段，控制是否退出for
	loop bool
	//定义账户的余额 []
	balance float64
	//每次收支的金额
	money float64
	//每次收支的说明
	note string
	//定义个字段，记录是否有收支的行为
	flag bool
	//收支的详情使用字符串来记录
	//当有收支时，只需要对details 进行拼接处理即可
	details string
}

//给结构体绑定对应方法

//编写要给工厂模式的构造方法，返回一个*Account实例
func NewAccount() *Account {
	return &Account{
		key:     "",
		loop:    true,
		balance: 10000.0,
		money:   0.0,
		note:    "",
		flag:    false,
		details: "收支\t账户金额\t收支金额\t说    明",
	}
}

//将显示明细写成一个方法
func (this *Account) showDetails() {
	fmt.Println("-----------------当前收支明细记录-----------------")
	if this.flag {
		fmt.Println(this.details)
	} else {
		fmt.Println("当前没有收支明细... 来一笔吧!")
	}
}

//将登记收入写成一个方法，和*Account绑定
func (this *Account) income() {

	fmt.Print("本次收入金额:")
	fmt.Scanln(&this.money)
	this.balance += this.money // 修改账户余额
	fmt.Print("本次收入说明:")
	fmt.Scanln(&this.note)
	//将这个收入情况，拼接到details变量
	//收入    11000           1000            有人发红包
	this.details += fmt.Sprintf("\n收入\t%v\t\t%v\t\t%v", this.balance, this.money, this.note)
	this.flag = true
}

//将登记支出写成一个方法，和*Account绑定
func (this *Account) pay() {
	fmt.Print("本次支出金额:")
	fmt.Scanln(&this.money)
	//这里需要做一个必要的判断
	if this.money > this.balance {
		fmt.Println("余额的金额不足")
		//break
	}
	this.balance -= this.money
	fmt.Print("本次支出说明:")
	fmt.Scanln(&this.note)
	this.details += fmt.Sprintf("\n支出\t%v\t\t%v\t\t%v", this.balance, this.money, this.note)
	this.flag = true
}

//将退出系统写成一个方法,和*Account绑定
func (this *Account) exit() {
	fmt.Println("你确定要退出吗? y/n")
	choice := ""
	for {

		fmt.Scanln(&choice)
		if choice == "y" || choice == "n" {
			break
		}
		fmt.Println("你的输入有误，请重新输入 y/n")
	}

	if choice == "y" {
		this.loop = false
	}
}

//显示主菜单
func (this *Account) MainMenu() {
	for {
		fmt.Println("\n----------收支记账软件----------")
		fmt.Println("	  1 收支明细")
		fmt.Println("	  2 登记收入")
		fmt.Println("	  3 登记支出")
		fmt.Println("	  4 退出软件")
		fmt.Print("请选择(1-4)：")
		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			this.showDetails()
		case "2":
			this.income()
		case "3":
			this.pay()
		case "4":
			this.exit()
		default:
			fmt.Println("请输入正确的选项..")
		}
		if !this.loop {
			break
		}

	}
}
