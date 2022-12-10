
# 创建用户表
drop table `user`;
create table `user`(
  `id` bigint(20) not null auto_increment primary key,
  `user_id` bigint(20) not null,
  `name` varchar(64) collate utf8_general_ci not null,
  `avatar` varchar(400),
  `sex` int default 0,
  `birthday` varchar(16),
  `password` varchar(64) collate utf8_general_ci not null,
  `phone_number` int8(11),
  `phone_code` int8(4),
  `wx` varchar(100),
  `email` varchar(100),
  `introduction` varchar(200),
  `open_id` varchar(64),
  `token` varchar(300),
  `source` varchar(100),
  `status` int default 0,
  `is_manager` boolean default false,

  `province` varchar(16) collate utf8_general_ci,
  `city` varchar(16) collate utf8_general_ci,
  `district` varchar(16) collate utf8_general_ci,
  `detail_position` varchar(64) collate utf8_general_ci,

  `diamond` bigint(12),
  `coin` bigint(12)
)engine = InnoDB default charset = utf8mb4 collate = utf8mb4_general_ci;

#