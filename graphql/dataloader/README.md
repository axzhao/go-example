# N + 1 Problem

`SELECT id FROM Parent`
and then executing a query for each record:
`SELECT * FROM Child WHERE parent_id = ?`

# dataloader

1. 批处理操作 Batching, 将多次查询合并成一次来减少查询次数。
2. 内存级别的缓存 Cache




