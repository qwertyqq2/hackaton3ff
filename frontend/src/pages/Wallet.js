import React, { useState } from 'react';

const Wallet = () => {
    return (
        <div>
            <div>
                <h3>Send</h3>
                <input type="text" id="name" name="name" required value="where"></input><br/>
                <input type="text" id="name" name="name" required value="count"></input>
            </div>
            <div>
                <h3>Receive</h3>
                <input type="text" id="name" name="name" required value="where"></input><br/>
                <input type="text" id="name" name="name" required value="count"></input>
            </div>
            <div>
                <h3>Exchange</h3>
                <input type="text" id="name" name="name" required value="where"></input><br/>
                <input type="text" id="name" name="name" required value="count"></input>
            </div>
        </div>
    )
}

export default Wallet;