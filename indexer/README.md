# StarkNet Indexer

```bash
sudo apt install sqlite3 
```

```bash
sqlite3 index.db < sql/erc20.sql
```

```bash
starknet-devnet
```

```bash
starknet deploy --no_wallet --gateway_url http://localhost:5050 --contract erc20.json --inputs 88772897131020503736870254 1229215828 18 1000000000000000000000 0 2636807177332077360661536399285498270125017543149115376957032708973395455206
```

```bash
insert into erc20 values(1, 'Index Token','IDXT', 18, 1000000000000000000000, 0, 193704331102769107990457894326770276842539641509620827278058837869738428138);
```
