--题目1：基本CRUD操作
-- 假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
-- 编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
    insert into students(`name`,`age`,`grade`) values ('张三',20,'三年级')
-- 编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
    select `id`,`name`,`age`,`grade` from students where `age` > 18
-- 编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
    update students set `grade` = '四年级' where `name`='张三'
-- 编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
    delete from students where `age` < 15

--题目2：事务语句
--假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表（包含字段 id 主键， from_account_id 转出账户ID，
-- to_account_id 转入账户ID， amount 转账金额）。
-- 编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。
-- 在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。
-- 如果余额不足，则回滚事务。
START TRANSACTION;
SELECT balance FROM accounts WHERE id = A FOR UPDATE;
IF (select balance FROM accounts WHERE id = A) >= 100 THEN
    UPDATE accounts SET balance = balance-100 WHERE id = A
    UPDATE accounts SET balance = balance+100 WHERE id = B
    INSERT INTO transactions(from_account_id,to_account_id,amount) VALUES(A,B,100)
    COMMIT;
else
    ROLLBACK;
END IF;