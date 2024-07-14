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

    this.askQueue = new PriorityQueue<Ask>({
      comparator: (a: Ask, b: Ask): number => {
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
  }

  addMarketEvent = (event: string[]): void => {
    if(event[0] === "BUY") {
      const buy: Buy = {
        price: parseFloat(event[1]),
        qty: parseFloat(event[2]),
        time: new Date()
      }
      this.addBuy(buy)
    }
    else if(event[0] === "ASK") {
      const ask: Ask = {
        price: parseFloat(event[1]),
        qty: parseFloat(event[2]),
        time: new Date()
      }
      this.addAsk(ask)
    }
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
    let largestBuy = this.getLargestBuy()
    while (!!largestBuy && largestBuy.price >= ask.price) {
      if (ask.qty < largestBuy.qty) {
        largestBuy.qty -= ask.qty
        ask.qty = 0
        break;
      } else {
        ask.qty -= largestBuy.qty
        this.popLargestBuy()
        largestBuy = this.getLargestBuy()
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
    if (this.askQueue.length == 0) {
      return null
    }
    return this.askQueue.dequeue()
  }

  getLargestBuy = (): Buy | null => {
    if (this.buyQueue.length == 0) {
      return null
    }
    return this.buyQueue.peek()
  }

  popLargestBuy = (): Buy | null => {
    if (this.buyQueue.length == 0) {
      return null
    }
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
    const event = line.split(" ")
    if(event[0] !== "#") {
      marketData.addMarketEvent(event)
      console.log("event", event)
      console.log("askQueue", marketData.askQueue.length > 0 && marketData.askQueue.peek())
      console.log("buyQueue", marketData.buyQueue.length > 0 && marketData.buyQueue.peek())
    }
  }

}

main()
