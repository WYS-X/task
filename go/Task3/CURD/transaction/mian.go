package main

import (
	"context"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Balance float64
}
type Transaction struct {
	gorm.Model
	FromAccountId int
	ToAccountId   int
	Amount        float64
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/learn?charset=utf8mb4&parseTime=True&loc=Local"
	db, dberr := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if dberr != nil {
		fmt.Println(dberr)
		panic("数据库链接失败")
	}

	db.AutoMigrate(&Account{})
	db.AutoMigrate(&Transaction{})

	//添加账号和余额
	ctx := context.Background()
	// gorm.G[Account](db).Create(ctx, &Account{Balance: 150})
	// gorm.G[Account](db).Create(ctx, &Account{Balance: 10})

	//开始转账
	amount := float64(100)
	tx := db.Begin()
	defer func() {
		if recover() != nil {
			tx.Rollback()
		}
	}()
	accountA, err1 := gorm.G[Account](tx).Raw("select * from accounts where id = ?", 1).First(ctx)
	accountB, err2 := gorm.G[Account](tx).Raw("select * from accounts where id = ?", 2).First(ctx)
	if err1 != nil || err2 != nil {
		fmt.Println("账号错误")
		return
	}
	if accountA.Balance < amount {
		fmt.Println("账户1余额不足")
		return
	}
	err := gorm.G[any](tx).Exec(ctx, "update accounts set balance = balance - ? where id = ?", amount, accountA.ID)
	if err != nil {
		tx.Rollback()
		fmt.Println("更改账户1余额失败")
		return
	}
	fmt.Println("账户1扣款成功")
	err = gorm.G[any](tx).Exec(ctx, "update accounts set balance = balance + ? where id = ?", amount, accountB.ID)
	if err != nil {
		tx.Rollback()
		fmt.Println("更改账户2余额失败")
		return
	}
	fmt.Println("账户2加钱成功")
	err = gorm.G[any](tx).Exec(ctx, "insert into transactions(from_account_id, to_account_id, amount, created_at) values(?,?,?,?)",
		accountA.ID, accountB.ID, amount, time.Now())
	if err != nil {
		tx.Rollback()
		fmt.Println("创建交易记录失败")
		return
	}
	fmt.Println("交易成功")
	tx.Commit()
}
