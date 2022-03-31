package main

// 有序集合
// ZAdd - 添加一个或者多个元素到集合，如果元素已经存在则更新分数
// ZCard - 返回集合元素个数
// ZCount - 统计某个分数范围内的元素个数
// ZIncrBy - 增加元素的分数
// ZRange,ZRevRange - 返回集合中某个索引范围的元素，根据分数从小到大排序
// ZRangeByScore,ZRevRangeByScore - 根据分数范围返回集合元素，元素根据分数从小到大排序，支持分页。
// ZRem - 删除集合元素
// ZRemRangeByRank - 根据索引范围删除元素
// ZRemRangeByScore - 根据分数范围删除元素
// ZScore - 查询元素对应的分数
// ZRank, ZRevRank - 查询元素的排名
