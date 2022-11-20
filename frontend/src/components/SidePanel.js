import Menu from "./Menu/Menu";
import '../App.css';
import React, { useState } from 'react';

const SidePanel = () => {
    const [menuActive, setMenuActive] = useState(false);
    const items = [
    {value: "Главная", href: '/main'}, 
    {value: "Услуги", href: '/service'},
    {value: "Магазин", href: '/shop'}];

    return (
        <div className='app'>
            <div className='burger-btn' onClick={() => setMenuActive(!menuActive)}>
                <span/>
            </div>
            <Menu active={menuActive} setActive={setMenuActive} header={"Меню"} items={items}/>
        </div>
    )
}

export default SidePanel;