create table `auth` (
`id` int(10) not null auto_increment,
`app_key` varchar(20) default '' comment 'key',
`app_secret`,
primary key (`id`) using btree

)engine=InnoDB default  CHARSET=utf8mb4 comment='认证管理';
insert into auth(id,app_key,app_secret,create_on,create_by,modify_on,modify_by,delete_on,is_del)values
(1,'home-system','zhaopeilin',0,'zpl',0,'',0,0);