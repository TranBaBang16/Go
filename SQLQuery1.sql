-- Tạo bảng 
Create database training_sql ;
use training_sql;
create table project_users (
id bigint not null primary key auto_increment,
project_id bigint not null ,
user_id bigint not null,
created_at timestamp default current_timestamp not null,
updated_at timestamp default current_timestamp not null
 );

create table projects (
id bigint not null auto_increment primary key,
project_name varchar(50) not null,
category_id  bigint  not null,
project_spend int not null,
project_variance int not null,
revenue_recognised int not null,
created_at timestamp default current_timestamp not null,
updated_at  timestamp default current_timestamp not null,
company_id bigint not null 
);
create table users (
id bigint not null auto_increment primary key,
user_name varchar(50) not null,
email varchar(100) not null,
pass_word varchar(50) not null,
birthday date not null ,
image_url text not null,
is_active bool not null,
created_at timestamp default current_timestamp not null,
updated_at timestamp default current_timestamp not null
);

ALTER TABLE Project_users
add  constraint project_users_project_fk
foreign key (project_id) references projects (id);

ALTER TABLE project_users
add  constraint project_users_users_id
foreign key (user_id) references users (id);

create table categories (
id bigint not null primary key auto_increment,
category_name varchar(50) not null,
created_at timestamp default current_timestamp not null,
updated_at timestamp default current_timestamp not null
);

create table companies (
id bigint not null primary key auto_increment,
company_name varchar(100) not null,
company_code varchar(100) not null,
created_at timestamp default current_timestamp not null,
updated_at timestamp default current_timestamp not null
);

alter table companies
modify column company_name varchar(50) not null;

alter table users
modify column email varchar(50) not null;

alter table users
modify column user_name varchar(16) not null;

alter table users
modify column pass_word varchar(10) not null;

alter table projects
add constraint project_category_fk
foreign key (category_id) references categories(id);

alter table projects
add constraint projects_companies_fk foreign key (company_id) references companies(id);

alter table projects
modify column project_spend int;

alter table projects
modify column project_variance int;

alter table projects
modify column revenue_recognised int;

alter table users
modify column birthday date;

alter table users
modify column image_url varchar(100) ;


-- CHèn dữ liệu vào bảng

insert into users(
user_name,
email,
pass_word,
birthday,
is_active
)
values (
'Tran B',
'ttbbbbb@gmail.com',
'12345678',
'1997/01/16',
true
);

select * from users;

Insert into companies(
company_name,
company_code
)
values (
'lll-lab',
'1122'
);

Insert into categories(
category_name
)
values (
'mmm'
);

insert into projects (
project_name,
category_id,
project_spend,
project_variance,
revenue_recognised,
company_id
)
values (
'project_5',
4,
132,
123,
23,
4
);


insert into project_users (
project_id,
user_id
)
values (
6,
2
);

-- Viết lệnh sql để có thể lấy thông tin những bản ghi của projects 
-- và số lượng user của mỗi projects đó (count user)

select projects.* ,count(project_users.user_id) as "count_user" from projects
inner join project_users on projects.id = project_users.project_id
group by id ;

-- viết lệnh sql để lấy ra danh sách các project của company có company_name = “monstar-lab” 

select * from projects
inner join companies
on projects.company_id = companies.id
where company_name = 'monstar-lab';

-- viết lệnh sql lấy ra danh sách các công ty có project có project_spend > 100

select distinct companies.company_name from companies
inner join projects
on projects.company_id=companies.id
where projects.project_spend>100;

-- viết lệnh sql để lấy ra thông tin của user tham gia vào projects
select distinct users.* from users 
inner join project_users on users.id = project_users.user_id
group by id;

-- lấy ra danh sách project mà có số lượng user tham gia > 10 , sắp xếp số lượng user tham gia tăng dần

select project_name, count(user_id) as "count_user" 
from projects
inner join project_users on projects.id= project_users.project_id
group by project_name
having count_user > 10
order by count_user asc;


-- Xoá project mà chưa có user nào tham gia

SET SQL_SAFE_UPDATES = 0;

delete from projects
where project_name not in (
select distinct project_name from(
select distinct projects.* from Projects
inner join project_users on projects.id = project_users.project_id
) as t
)
