package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/istsh/go-sqlboiler-example/models"

	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
)

func main() {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=dev dbname=postgres password=pass sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	now := time.Now()
	// dropTable(db)
	// timeAfterDropTable := time.Now()
	// fmt.Printf("dropTable: %s\n", timeAfterDropTable.Sub(now).String())

	// migrate(db)
	// timeAfterMigrate := time.Now()
	// fmt.Printf("migrate: %s\n", timeAfterMigrate.Sub(timeAfterDropTable).String())

	insert(db)
	timeAfterInsert := time.Now()
	fmt.Printf("insert: %s\n", timeAfterInsert.Sub(now).String())

	selectAndUpdate(db)
	timeAfterSelectAndUpdate := time.Now()
	fmt.Printf("selectAndUpdate: %s\n", timeAfterSelectAndUpdate.Sub(timeAfterInsert).String())

	selectAndDelete(db)
	timeAfterSelectAndDelete := time.Now()
	fmt.Printf("selectAndDelete: %s\n", timeAfterSelectAndDelete.Sub(timeAfterSelectAndUpdate).String())
}

// func dropTable(db *sql.DB) {
// 	// テーブルを削除
// 	db.DropTable(&User{})
// }

// func migrate(db *gorm.DB) {
// 	// テーブルを作成
// 	db.AutoMigrate(&User{})
// }

func insert(db *sql.DB) {
	ctx := context.Background()
	for i := 1; i <= 100; i++ {
		// now := null.NewTime(time.Now(), true)
		u := models.User{
			// CreatedAt: now,
			// UpdatedAt: now,
			Name: null.NewString(fmt.Sprintf("sqlboiler_test_user_%03d", i), true),
		}
		// データを登録
		if err := u.Insert(ctx, db, boil.Infer()); err != nil {
			fmt.Printf("err(id = %d): %s\n", i, err.Error())
			continue
		}
	}
}

func selectAndUpdate(db *sql.DB) {
	ctx := context.Background()
	for i := 1; i <= 100; i++ {
		u, err := models.FindUser(ctx, db, i)
		if err != nil {
			fmt.Printf("err(id = %d): %s\n", i, err.Error())
			continue
		}
		u.Name = null.NewString(fmt.Sprintf("sqlboiler_test_user_%03d_updated", i), true)
		// データを更新
		_, err = u.Update(ctx, db, boil.Infer())
		if err != nil {
			fmt.Printf("err(id = %d): %s\n", i, err.Error())
			continue
		}
	}
}

func selectAndDelete(db *sql.DB) {
	ctx := context.Background()
	for i := 1; i <= 100; i++ {
		u, err := models.FindUser(ctx, db, i)
		if err != nil {
			fmt.Printf("err(id = %d): %s\n", i, err.Error())
			continue
		}
		// データを削除
		_, err = u.Delete(ctx, db)
		if err != nil {
			fmt.Printf("err(id = %d): %s\n", i, err.Error())
			continue
		}
	}
}
