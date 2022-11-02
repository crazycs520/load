package payload

import (
	"database/sql"
	"fmt"

	"github.com/crazycs520/loadgen/cmd"
	"github.com/crazycs520/loadgen/config"
	"github.com/crazycs520/loadgen/util"
	"github.com/spf13/cobra"
)

type FKDeleteCheckSuite struct {
	cfg       *config.Config
	db *sql.DB
}

func NewFKDeleteCheckSuite(cfg *config.Config) cmd.CMDGenerater {
	return &FKDeleteCheckSuite{
		cfg: cfg,
	}
}

func (c *FKDeleteCheckSuite) Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "fk-delete-check",
		Short:        "payload of insert with foreign key check",
		RunE:         c.RunE,
		SilenceUsage: true,
	}
	return cmd
}

func (c *FKDeleteCheckSuite) RunE(cmd *cobra.Command, args []string) error {
	return c.Run()
}

func (c *FKDeleteCheckSuite) prepare() error {
	c.db = util.GetSQLCli(c.cfg)
	prepareSQLs := []string{
		"set @@global.tidb_enable_foreign_key=1",
		"set @@foreign_key_checks=1",
		"drop table if exists t1,t2",
		"create table t1 (id int key, name varchar(10));",
		"create table t2 (id int, pid int, unique index(id), foreign key fk(pid) references t1(id));",
		"insert into t1 values (0, ''), (1, 'a'), (2, 'b'), (3, 'c'), (4, 'd'), (5, 'c'), (6, ''), (7, 'a'), (8, 'b'), (9, 'c'), (10, 'd')",
	}
	for _,sql := range prepareSQLs{
		_, err := c.db.Exec(sql)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *FKDeleteCheckSuite) Run() error {
	err := c.prepare()
	if err != nil {
		fmt.Println("prepare table meet error: ", err)
		return err
	}
	fmt.Println("started")
	cnt :=0
	for {
		cnt++
		c.db.Exec("begin")
		_,err := c.db.Exec("delete from t1")
		if err != nil {
			fmt.Println("exec meet error: ", err)
			return err
		}
		c.db.Exec("rollback")
	}
}

