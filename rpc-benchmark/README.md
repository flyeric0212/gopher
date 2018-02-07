GoLang RRC 框架新能对比

message: 600 bytes 左右

**go_rpc:**

0ms delay   100 clients  1000000 requests

TPS: 约40000
total: 23s左右


**grpc:**

0ms delay   100 clients  1000000 requests

TPS: 约12000
total: 70 ~ 80s
cpu占用高

**rpcx:**

0ms delay   100 clients  1000000 requests

TPS: 约31000
total: 30s左右


jsonrpcx: