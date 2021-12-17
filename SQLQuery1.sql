--Tạo bảng
Create database bang_test
use bang_test;
create table Project_user (
id bigint not null Identity(1,1) primary key,
project_id bigint not null ,
user_id bigint not null,
created_at timestamp ,
updated_at date
 );

create table Project (
id bigint not null Identity(1,1) primary key,
project_name varchar(50) not null,
category_id  bigint  not null,
project_spend int not null,
project_variance int not null,
revenue_recognised int not null,
created_at timestamp not null,
updated_at date not null,
company_id bigint

);
create table users (
id bigint not null Identity(1,1) primary key,
user_name varchar(50) not null,
email varchar(100) not null,
pass_word varchar(50) not null,
birthday date not null ,
image_url text not null,
is_active bit not null,
created_at timestamp not null,
updated_at date not null,
);
use bang_test
ALTER TABLE Project_user
add  foreign key (project_id) references project (id);

ALTER TABLE Project_user
add  foreign key (user_id) references users (id);

create table categories (
id bigint not null primary key identity(1,1),
category_name varchar(50) not null,
created_at timestamp not null,
updated_at date not null
)

create table companies (
id bigint not null primary key identity(1,1),
company_name varchar(100) not null,
company_code varchar(100) not null,
created_at timestamp not null,
updated_at date not null
)

alter table companies
alter column company_name varchar(50) not null

alter table users
alter column email varchar(50) not null

alter table users
alter column user_name varchar(16) not null

alter table users
alter column pass_word varchar(10) not null

alter table project
add foreign key (category_id) references categories(id);

alter table project
add foreign key (company_id) references companies(id);

alter table project
alter column project_spend int

alter table project
alter column project_variance int

alter table project
alter column revenue_recognised int

alter table users
alter column birthday date

alter table users
alter column image_url varchar


-- CHèn dữ liệu vào bảng

insert into users(
user_name,
email,
pass_word,
birthday,
is_active,
created_at,
updated_at
)
values (
'Dao van B',
'yyyyyyyyyyy@gmail.com',
'12345678',
'2002/12/23',
1,
DEFAULT,
'2021/12/23'
)

Insert into companies(
company_name,
company_code,
created_at,
updated_at
)
values (
'monstar-lab',
'1234',
default,
'2011-11-01'
)

insert into Project (
project_name,
category_id,
project_spend,
project_variance,
revenue_recognised,
created_at,
updated_at,
company_id
)
values (
'project_4',
2,
13,
1234,
23,
default,
'2021/05/30',
3
)

insert into Project_user (
project_id,
user_id,
created_at,
updated_at
)
values (
3,
1,
default,
'2020/12/05'
)

--Viết lệnh sql để có thể lấy thông tin những bản ghi của projects 
--và số lượng user của mỗi projects đó (count user)
select project.id,project_name,project_spend,Project.project_variance,Project.category_id, count(users.id) as "count_user" from Project
left join project_user on Project.id= Project_user.project_id
left join users on Project_user.user_id=USERs.id
group by project.id,project_name,project_spend,Project.project_variance,Project.category_id


--viết lệnh sql để lấy ra danh sách các project của company có company_name = “monstar-lab” 

select * from Project
inner join companies
on Project.company_id=companies.id
where company_name = 'monstar-lab'

--viết lệnh sql lấy ra danh sách các công ty có project có project_spend > 100

select companies.company_name from companies
inner join Project
on Project.company_id=companies.id
where project.project_spend>100

--viết lệnh sql để lấy ra thông tin của user tham gia vào projects
select distinct USERs.id,user_name,email,pass_word,birthday,image_url,USERs.created_at,USERs.updated_at from users 
inner join Project_user
on users.id=Project_user.user_id
inner join Project
on Project_user.project_id=Project.id

--lấy ra danh sách project mà có số lượng user tham gia > 10 , sắp xếp số lượng user tham gia tăng dần

select project_name,
count(users.id) as "count_user" 
from Project
left join project_user on Project.id= Project_user.project_id
left join users on Project_user.user_id=USERs.id
group by project_name
delete fro
--Xoá project mà chưa có user nào tham gia



