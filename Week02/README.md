# Week02 作业
## 题目：

我们在数据库操作的时候，比如 `dao` 层中当遇到一个 `sql.ErrNoRows` 的时候，是否应该 `Wrap` 这个 `error`，抛给上层。为什么，应该怎么做请写出代码？
## 我的思考：



**1,要不要 wrap 抛给上层：**

一般情况下只有最底层的函数需要 wrap。因为这个 error 发生在 dao 层。而 dao 层一般也是一个业务最底层的地方。所以应该 Wrap 然后抛给上层。



**2,如何做：**

可以 wrap 一些查询语句，在 debug 时方便定位问题。

见下方代码实现，同时在同级目录下有相应的 go 文件。



## 代码实现：

> Dao —> service —> api

**调用顺序：api —> service —> dao 发生错误，层层返回。**

**1.最基本的查询方法中不进行 wrap**

**2.但在 dao 层进行查询时出现错误，则 wrap 携带信息返回上层**

**3. service 层直接返回从dao 获取的信息**

**4. api 层判断 err,发现不为空则 log.Fatal(err),err 为空则正常返回数据**

``` go
// ./dao/dao.go

package dao

import (
	"week2/model"

	"github.com/pkg/errors"
)

// GetEntity .
func GetEntity(query int) (*model.Entity, error) {
	result := &model.Entity{}
	result, err := queryEntityFromDBByID(query)
	if err != nil {
		return nil, errors.Wrap(err, "record not found")
	}
	return result, nil
}

// queryEntityFromDBByID 模拟从数据库获取数据
func queryEntityFromDBByID(id int) (*model.Entity, error) {
	// 省略逻辑
	return &model.Entity{}, nil
}


```


``` go
// ./service/service.go

package service

import (
	"week2/dao"
	"week2/model"
)

// GetEntity .
func GetEntity(query int) (*model.Entity, error) {
	return dao.GetEntity(query)
}

```

``` go
// ./api/api.go
package api

import (
	"log"
	"week2/model"
	"week2/service"
)

// GetEntity .
func GetEntity() (*model.Entity, error) {
	testID := 1
	entity, err := service.GetEntity(testID)
	if err != nil {
		log.Fatal(err)
	}
	return entity, nil
}

```
