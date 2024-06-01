# 数悦深者  
本产品是结合了区块链和AI的元计算资源云端整合系统。在对闲置资源实现调度的过程中，高度保护了用户的隐私性，首创性的提出了大语言模型时代的无痕浏览。
![image](https://github.com/nft-maker-one/Intersection_of_metadata_and_AI/assets/121859606/353c60c2-976f-47cc-88e9-5b217b85829c)

## 启动教程
1. 运行数据库文件，利用 Mysql 创建用户的基本信息
2. 开启本地 redis 数据库服务端，使用默认端口号 6379
3. 开启 golang 后端
4. 开启 vue 前端
5. 启动算力节点(可选)
## 原理拆解

### 数悦链(Fancy Digital Chain)
本部分使用 golang 实现，基于 kademila 哈希表在网络中进行节点的发现与搜寻，实现了以下功能
1. 更改 config.yaml 文件中的 API_KEY 配置信息，完成节点初始化
2. 运行算力节点，并通过网络中的广播节点导入网络
3. 广播算力以及对应价格至 UnlimiteMarketplace 智能合约进行发布
![image](https://github.com/nft-maker-one/Intersection_of_metadata_and_AI/assets/121859606/fddb7579-4d34-4a75-b0b9-1ee316099388)


### 无源算力交易平台
本部分使用 golang 和 vue 实现，具体功能如下
1. 实时获取数悦链上 UnlimiteMarketplace 智能合约的信息，并在平台进行发布
2. 通过 UnlimiteMarketplace 质押用户相应资产(这个过程中平台只知道用户的公钥而不会保存任何用户隐私)，保证后续能够支付调用算力节点的费用
3. 检查算力节点所公布的 API_KEY 最近的调度情况，如果调度时间和回复内容与其返回结果一致，判定算力节点完成相应的任务，向其支付先前用户质押的字节
![image](https://github.com/nft-maker-one/Intersection_of_metadata_and_AI/assets/121859606/9bb351ad-0ba9-4600-a11a-38dacb3177fb)


