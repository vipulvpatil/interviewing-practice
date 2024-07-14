Let's implement an orderbook.

This is to better understand the concepts of an orderbook and also to brush up my typescript.

Reference API used.
https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md#order-book

Phase 1.
Implement these functions. 
1. Add a Bid.
-> Executes the bid if a matching ask exists.
2. Add an Ask.
-> Executes the bid if a matching bid exists.
3. Calculate current Price for selling n units.
-> This will return the value if n units are to be sold immediately
4. Calculate current Price for buying n units.
-> This will return the value if n units are to be bought immediately

Data examples
Bid {
  price: string
  qty: string
}

Ask  {
  price: string
  qty: string
}
