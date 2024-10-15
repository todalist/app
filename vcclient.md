### 修改vcclient定制
- 由于要加网关，因此要改url地址前缀，用正则改js源码的请求地址为相对路径可解
- sio的请求地址需要在启动时针对每个服务器进行修改，正则直接写在源码中，在升级vcclient需注意
- 删除vcclient中从huggingface下载示例模型的代码