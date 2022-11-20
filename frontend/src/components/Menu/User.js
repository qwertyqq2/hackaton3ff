import React from "react"
import './User.css'
import '../../pages/Main.css'
import Course from "../../Course"
import { Link } from "react-router-dom"
import UserData from "../../UserData"

const User = () => {

  let tmp = [
    {
      name: "Bitcoin",
      img: "bit.png",
      ref: "/",
      cost: "0"
    },
    {
      name: "Ethereum",
      img: "ether.png",
      ref: "/",
      cost: "0"
    }
  ];

  let arr = [];

  for (let i = 0; i < 2; i++) {
    arr.push(
      new Course(tmp[i].name, tmp[i].img, tmp[i].ref, tmp[i].cost)
    );
  }

  let divOutput = arr.map(item => <a className='Market_course'>
    <div className='Market_card'>
      <div className='Market_1row'><img src={item.img} alt="card" />{item.name}</div>
      <div className='Market_1row Market_2row'>{item.cost} 
      <Link to="/wallet"><img src="arrow.png"/></Link> 
      <Link to="/wallet"><img src="arrow.png" className="imgRot"/></Link> 
      <Link to="/wallet"><img src="exchange.png"/></Link> </div>
    </div>
  </a>)

  let testUser = new UserData("platypus", divOutput, "Standart")

  return (
    <div className="user">
      <ul>
        <li>
          <img src="utkonos.jpg" alt="avatar" />
        </li>
        <li>
          <h3>Access level: {testUser.access}</h3>
        </li>
        <li>
          <h2>{testUser.name}</h2>
        </li>
        <li>
          <div className='Markets_container'>
            {testUser.wallets}
          </div>
        </li>
      </ul>
    </div>
  )
}

export default User;