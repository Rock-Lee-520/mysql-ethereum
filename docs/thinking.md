学习文档：
https://goethereumbook.org/
https://etherscan.io/


### 查询账户
1、查询账户余额、代币余额， 
select * from account where  address="0x71c7656ec7ab88b098defb751b7401b5f6d8976f"


2、查询账户余额、代币余额， 增加区块筛选
select * from account where  address="0x71c7656ec7ab88b098defb751b7401b5f6d8976f" and  block_number=""


3、查询账户待处理余额
select * from account where  address="0x71c7656ec7ab88b098defb751b7401b5f6d8976f" and  PendingBalanceAt=true


### 钱包
1、创建钱包表
create table wallet  
public_key  varchar(255）
private_key  varchar(255）


2、创建钱包地址
insert into wallet public_key="" and   private_key=""


### 交易

此处省略。。。。。。



