// inout params
// author: baoqiang
// time: 2019/1/21 下午4:10
package learndb

import (
	"database/sql"
	"log"
	"context"
	_ "github.com/go-sql-driver/mysql"
	//_ "github.com/go-goracle/goracle"
	//_ "github.com/ziutek/mymysql/godrv"
)

//var MysqlUrl = "mysql://root:00@127.0.0.1:3306/hello?option=open&charset=utf8"
var MysqlUrl = "root:00@tcp(localhost:3306)/hello"
//var MysqlUrl = "tcp:127.0.0.1:3306*go/root/00"
//var MysqlUrl = "tcp://127.0.0.1:3306/hello/xiao/88?charset=utf8&keepalive=1200"

func RunInout(){
//		sqltextcreate := `
//CREATE PROCEDURE abinout
//   @aid INT,
//   @bid INT OUTPUT
//AS
//BEGIN
//   SELECT @bid = @aid + @bid;
//END;
//`

sqltextcreate := `
create procedure abinout(in aid int, in bid int) 
begin
SELECT @bid = @aid + @bid;
end;
`


		sqltextdrop := `DROP PROCEDURE abinout;`
		sqltextrun := `abinout`


		db, err := sql.Open("mysql", MysqlUrl)
		//db, err := sql.Open("mymysql", MysqlUrl)
		if err != nil {
			log.Fatalf("failed to open driver %v", err)
		}
		defer db.Close()

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		db.ExecContext(ctx, sqltextdrop)
		_, err = db.ExecContext(ctx, sqltextcreate)
		if err != nil {
			log.Fatal(err)
		}
		var bout int64 = 3
		_, err = db.ExecContext(ctx, sqltextrun,
			sql.Named("aid", 5),
			sql.Named("bid", sql.Out{Dest: &bout}),
		)
		defer db.ExecContext(ctx, sqltextdrop)
		if err != nil {
			log.Fatal(err)
		}

		if bout != 8 {
			log.Fatalf("expected 8, got %d", bout)
		}
	}
