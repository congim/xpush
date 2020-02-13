package mysql

var SQLStoreMsg string = "insert into `msgs` (`topic`, `id`, `original_id`, `type`, `payload`, `insert_time`) values ( ?, ?, ?, ?, ?, ?)"

//var SQLPullMsg string = "SELECT `id`, `original_id`, `type`, `payload`, `insert_time` FROM `msgs` WHERE `topic`=? ORDER BY `insert_time` desc LIMIT ?,?;"

var SQLPullMsg string = "SELECT id, original_id, type, payload, insert_time FROM msgs WHERE topic=? AND insert_time <=? ORDER BY insert_time DESC LIMIT ?,?;"
