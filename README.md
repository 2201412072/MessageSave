# MessageSave
MessageSave是一款密码管理软件，区别于其他不同的密码管理软件，本软件需要两人解密才能显示出密码，因此大大提升了安全性。
本软件内部存储的密码经过双人加密，因此即使黑客掌握一个人的全部信息、全部权限，如果不知道另外一个人的信息，也无法获取密码，十分的安全。

V1版本
原理介绍：
1、每个都有一个公钥、私钥，其中私钥需要保存好，公钥则可以向别人传播。
2、首先需要添加关联用户的用户名和公钥。
3、当需要密码进行加密时，选择要加密的关联用户，使用其公钥以及自己的公钥进行加密，将结果存储到软件中。
4、当需要解密时，需要将软件中给出的信息发送给对应用户，然后将其返回信息进一步解码，得到最终密码。


功能介绍：
1、有添加关联用户、公钥的功能。
2、能够管理密码，可以查询一个应用对应的关联用户，以及一个关联用户对应的应用密码。
3、能够添加密码。
4、能够保存重要的信息，具体操作同密码。