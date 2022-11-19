import React from 'react';
import Course from '../Course';
import './Main.css';

let tmp = [
    {
        name: "Bitcoin",
        img: "bit.png",
        ref: "/graph",
        cost: 123
    },
    {
        name: "Ethereum",
        img: "ether.png",
        ref: "/",
        cost: 15
    },
    {
        name: "Dogecoin",
        img: "dodge.png",
        ref: "/",
        cost: 1
    },
    {
        name: "Ethereum",
        img: "ether.png",
        ref: "/",
        cost: 15
    },
    {
        name: "Ethereum",
        img: "ether.png",
        ref: "/",
        cost: 15
    },
    {
        name: "Ethereum",
        img: "ether.png",
        ref: "/",
        cost: 15
    },
    {
        name: "Ethereum",
        img: "ether.png",
        ref: "/",
        cost: 15
    },
    {
        name: "Ethereum",
        img: "ether.png",
        ref: "/",
        cost: 15
    },
];

let arr = [];
for (let i = 0; i < 8; i++) {
    arr.push(
        new Course(tmp[i].name, tmp[i].img, tmp[i].ref, tmp[i].cost)
    );
    console.log("success");
}

// for (let i = 0; i < 3; i++) {
//     arr.push(
//         i
//     );
// }

const Main = () => {
    let divOutput = arr.map(item => <a key={Course.id++} href={item.ref} className='Market_course'>
        <div className='Market_card'>
            <div className='Market_1row'><img src={item.img} alt="card"/> {item.name}</div>
            <div className='Market_1row Market_2row'>{item.cost}</div>
        </div>
    </a>)

    return (
        <div>
            <div className='Market_center'><h2>Начать торговлю</h2></div>
            <div className='Markets_container'>
                {divOutput}
            </div>
        </div>
    )
}

export default Main;