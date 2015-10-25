# API based on a single "trade" operation.  Neglects difference between "buy" and "sell"

import requests

# Array holding all possible wallets (right now only one currency)
wallet = [0, 0]

def trade( nodeA, nodeB, dx, cur, wallet):
    # trade dx units of currency from nodeA to nodeB
    dz = dx*requests.get('https://blockchain.info/ticker').json()[cur]['15m']
    wallet[nodeA] += dz
    wallet[nodeB] -= dz
    return wallet
