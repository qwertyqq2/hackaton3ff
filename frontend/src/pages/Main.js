import React, { useState } from 'react';
import Course from '../Course';
import './Main.css';

const Main = () => {

    let tmp = [
        {
            name: "Bitcoin",
            img: "bit.png",
            ref: "https://trade.dydx.exchange/trade/BTC-USD",
            cost: "$16,657"
        },
        {
            name: "Ethereum",
            img: "ether.png",
            ref: "https://trade.dydx.exchange/trade/ETH-USD",
            cost: "$1,216.8"
        },
        {
            name: "Dogecoin",
            img: "dodge.png",
            ref: "https://trade.dydx.exchange/trade/DOGE-USD",
            cost: "$0.0848"
        },
        {
            name: "Uniswap",
            img: "uniswap.png",
            ref: "https://trade.dydx.exchange/trade/UNI-USD",
            cost: "$5.793"
        },
        {
            name: "Litecoin",
            img: "litecoin.png",
            ref: "https://trade.dydx.exchange/trade/LTC-USD",
            cost: "$63.4"
        }
    ];

    let arr = [];
    for (let i = 0; i < 5; i++) {
        arr.push(
            new Course(tmp[i].name, tmp[i].img, tmp[i].ref, tmp[i].cost)
        );
    }

    let divOutput = arr.map(item => <a key={Course.id++} href={item.ref} className='Market_course'>
        <div className='Market_card'>
            <div className='Market_1row'><img src={item.img} alt="card" />{item.name}</div>
            <div className='Market_1row Market_2row'>{item.cost}</div>
        </div>
    </a>)

    return (
        <div>
            <div className='Market_center'><h2>Start trading</h2></div>
            <div className='Markets_container'>
                {divOutput}
            </div>
        </div>
    )
}

export default Main;