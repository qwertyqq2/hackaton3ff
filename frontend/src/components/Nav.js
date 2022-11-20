import './Nav.css';
import {Link, useMatch, useResolvedPath} from 'react-router-dom';
import SidePanel from './SidePanel';
import React, { useState } from 'react';
import UserStatus from '../State';

const NavbarWithLink = () => {
    //console.log("Nav ", UserStatus.getStatus());
    const [userActive, setUserActive] = useState(UserStatus.getStatus());

    let isUserEnter = userActive;
    let sidePanel = isUserEnter ? <SidePanel /> : <span className='plug'/>

    let wallet = isUserEnter ? <CustomLink to={"/wallet"}>Wallet</CustomLink> : null

    let enter = isUserEnter ? null : <CustomLink to={"/enter"} onClick={() => setUserActive(!userActive)}>Enter</CustomLink>

    let register = isUserEnter ? null : <CustomLink to={"/register"}>Register</CustomLink>

    let exit = isUserEnter ? <CustomLink to={"/"}>
        <a href='/' onClick={() => setUserActive(!userActive)}>
            <img src="logout.png" className="logout" alt='exit' />
        </a>
    </CustomLink> : null

    return (
        <nav className="nav">
            {sidePanel}
            
            <div className='ulist'>
                <CustomLink to={"/main"}>Main</CustomLink>
                {wallet}
                {enter}
                {register}
                {exit}
            </div>
        </nav>
    )
}

function CustomLink({to, children, ...props}) {
    const resolvedPath = useResolvedPath(to)
    const isActive = useMatch({ path: resolvedPath.pathname, end: true })

    return (
        <li className={isActive ? 'active' : ''}>
            <Link to={to} {...props}>{children}</Link>
        </li>
    )
}

export default NavbarWithLink;