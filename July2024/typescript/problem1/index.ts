import PriorityQueue from "ts-priority-queue"
import * as fs from "fs"
import readline from "readline"

interface Buy {
  price: number
  qty: number
  time: Date
}

interface Ask {
  price: number
  qty: number
  time: Date
}

class MarketData {
  buyQueue: PriorityQueue<Buy>
  askQueue: PriorityQueue<Ask>

  constructor() {
    this.buyQueue = new PriorityQueue<Buy>({
      comparator: (a: Buy, b: Buy): number => {
        if (a.price > b.price) {
          return 1
        }
        if (a.price == b.price) {
          if (a.time <= b.time) {
            return 1
          }
          return -1
        }
        return -1
      }
    })

    this.askQueue = new PriorityQueue<Buy>({
      comparator: (a: Buy, b: Buy): number => {
        if (a.price < b.price) {
          return 1
        }
        if (a.price == b.price) {
          if (a.time <= b.time) {
            return 1
          }
          return -1
        }
        return -1
      }
    })
  }
  
  addBuy = (buy: Buy): void => {
    let smallestAsk = this.getSmallestAsk()
    while (!!smallestAsk && smallestAsk.price <= buy.price) {
      if (buy.qty < smallestAsk.qty) {
        smallestAsk.qty -= buy.qty
        buy.qty = 0
        break;
      } else {
        buy.qty -= smallestAsk.qty
        this.popSmallestAsk()
        smallestAsk = this.getSmallestAsk()
      }
    }
    if(buy.qty > 0) {
      this.buyQueue.queue(buy)
    }
  }
  
  addAsk = (ask: Ask): void => {
    let smallestBuy = this.getSmallestBuy()
    while (!!smallestBuy && smallestBuy.price >= ask.price) {
      if (ask.qty < smallestBuy.qty) {
        smallestBuy.qty -= ask.qty
        ask.qty = 0
        break;
      } else {
        ask.qty -= smallestBuy.qty
        this.popSmallestAsk()
        smallestBuy = this.getSmallestBuy()
      }
    }
    if(ask.qty > 0) {
      this.askQueue.queue(ask)
    }
  }

  getSmallestAsk = (): Ask | null => {
    if (this.askQueue.length == 0) {
      return null
    }
    return this.askQueue.peek()
  }

  popSmallestAsk = (): Ask | null => {
    return this.askQueue.dequeue()
  }

  getSmallestBuy = (): Buy | null => {
    if (this.buyQueue.length == 0) {
      return null
    }
    return this.buyQueue.peek()
  }

  popSmallestBuy = (): Buy | null => {
    return this.buyQueue.dequeue()
  }
}


async function main() {
  let marketData: MarketData = new MarketData()
  const fileStream = fs.createReadStream('input.txt')

  const rl = readline.createInterface({
    input: fileStream,
    crlfDelay: Infinity,
  })

  for await (const line of rl) {
    process
  }
}

main()
