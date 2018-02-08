GoLang RRC 框架新能对比

本人MacPro  
CPU: 2.7 GHz Intel Core i5   
Memory: 8G
System: macOS Sierra 10.12.6

message size: 600 bytes 左右

**go_rpc:**

0ms delay   100 clients  1000000 requests
TPS: 约51000
total: 19s左右

10ms delay   100 clients  1000000 requests
TPS: 约7600
total: 131s左右

**rpcx:**

0ms delay   100 clients  1000000 requests
TPS: 约40000
total: 24s左右

10ms delay   100 clients  1000000 requests
TPS: 约7600
total: 130s左右

**grpc:**

0ms delay   100 clients  1000000 requests
TPS: 约15000
total: 63s左右

10ms delay   100 clients  1000000 requests
TPS: 约8300
total: 119s左右

**jsonrpc:**

0ms delay   100 clients  1000000 requests
TPS: 约13700
total: 72s左右

10ms delay   100 clients  1000000 requests
TPS: 约8400
total: 118s左右