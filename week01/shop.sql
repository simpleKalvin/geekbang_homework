CREATE table `goods`
(
    `id`     int not null primary key,
    `name`   varchar(32) comment '产品名称',
    `money`     decimal(10, 2) comment '产品价格',
    `amount` int comment '库存'
);


CREATE table `order`
(
    `id`        int not null primary key,
    `order_no` int not null default 0 unique key,
    `money`     decimal(10, 2) comment '订单金额',
    `status`    int comment '0 未支付 1 已支付 2 已退款'
);