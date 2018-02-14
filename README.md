# hashcash-go
Implementing Simple Hashcash Alg. in Go

To see how does it work:

```bash

git clone https://github.com/rezamt/hashcash-go

cd hashcash-go

go build

./hashcash-demo --complexity 4 # means hash must starts with 0 leading Zeros


```

and output should looks like:

```bash
Hashing: 0000046eb0d73d55b1b6811322e2ecba7158e8f8ea6cc482ed9cc1e2e2c40df0
Block Data:	9D6D1B17974FCF31BF38396A91CD98DF00C7C2B3DB8E157F4BFE9A7854173784C156B73B4EF67BEEF4D34B2B8DCC4DD876F8BD98AD723BB19205C3F89B241CBE9E8A5B550BE64AC34C10FA88DBA1B6C9363E84643E3ACC69AD92B8AFE12C7549E5C05E33
Hash:	0000046EB0D73D55B1B6811322E2ECBA7158E8F8EA6CC482ED9CC1E2E2C40DF0
Nonce:	22889
Complexity:	4
Hashing time:	0.126 sec

```

Note: 

- Blockdata is just a random data for testing purpose.
- Nonce is just an incremental value starting from 1
- Hashing time: the time my application spends to find the hash with n x of Zero at the begining

