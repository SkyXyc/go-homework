package logic

import (
	"calderxu_workshop1_msg_notification/dao"
	"database/sql"
	"github.com/pkg/errors"
)

func MockLogic() error {
	// 如果在业务逻辑代码中遇到查询sql返回的错误，
	// 我认为需要分情况来决定是否将error进行wrap后抛给上层
	// 1、如果是外部传入的sql查询参数，需要wrap这个错误来告知调用方进行参数的检查，所以需要将error和详细的堆栈情况传出
	if err := dao.MockQuerySql(); err != nil {
		return errors.Wrap(err, "query sql failed with args : xxx")
	}

	//2、如果是内部的查询逻辑，查询sql的参数由服务本身生成，那么在合适的时机需要拆出原来的err，与sql.ErrNoRows比较并内部降级处理
	if err := dao.MockQuerySql(); err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			// 内部降级处理，重新获取参数进行重试查询sql
		}
	}

	//来自刚接触后台的萌新的回答，希望老师多加指教～
	return nil
}
