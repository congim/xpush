package mysql

var SQLStoreMsg string = "insert into `xpush`.`msgs` (`topic`, `msg_id`, `type`, `payload`, `insert_time`) values ( ?, ?, ?, ?, ?)"
