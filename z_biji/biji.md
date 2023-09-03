## golang 快捷键
多选: 长按ALT  
光标放到每个单词开头: ctrl + <-  
块状选择: Alt+Shift  
多行首字母大写: 块状选择 -> 光标前置 -> `shift + ->选中 -> Ctrl+Shift+U  


git remote add origin https://github.com/lilihaooo/gin_blog.git


## git
git init  
git checkout -b main 创建并切换到mian分支  
git add .  
<--  
1. git reset -- HEAD go.mod 从暂存区移除go.mod文件按  
-->

git commit -m  提交后就会git branch 查看当前分支main, 如果之前没有创建的话就是默认master

git log 提交记录

git remote add origin git@github.com:lilihaooo/gin_blog.git 建立远程连接
git remote -v 查看远程连接


撤销提交
git revert --hard B
A -- B -- C -- D (main)

git reset --hard B 慎用
A -- B -- C (main)


## fmt包
1. fmt.Sprintf("%s:%s", s.Host, s.Port)格式化输入字符串,    
   * %s 输出string为字符串
   * %d 输出整型为字符串 
   * %T 输出类型
2. fmt.Print("请输入用户名: ") // 不换行
3. fmt.Println("")   // 换行
4. fmt.Scanln(&userModel.UserName)   // 可以输入空值, 回车后进入下一条
5. fmt.Scan(&userModel.UserName)     // 输入空值回车后依然再本条, 直到非空




## gorm 
1. db.Table 临时指定表名

## 标签:
1. AnimalID int64     `gorm:"column:beast_id"`         // 将列名设为 `beast_id`


## validate
1. IsShow bool   `json:"is_show" validate:"required,oneof=1 2" label:"是否展示"` // 是否展示
2. BannerSortList []bannerSort `validate:"dive"` //进入嵌套结构体中验证



## 结构体转为map
1. "github.com/fatih/structs"
   // 将结构体转为map
   crMap := structs.Map(&cr)


## jwt
https://blog.csdn.net/xmcy001122/article/details/126501371

