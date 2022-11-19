import Login from "../components/Login/Login";
import React from 'react';
import useToken from "../components/useToken";
import UserStatus from '../State'

// function setToken(userToken) {
//     sessionStorage.setItem('token', JSON.stringify(userToken));
// }

// function getToken() {
//     const tokenString = sessionStorage.getItem('token');
//     const userToken = JSON.parse(tokenString);
//     return userToken?.token
// }

const Enter = () => {
    const { token, setToken } = useToken();
    //const token = getToken();
    //UserStatus.setStatus(!token);

    if (!token) {
        return <Login setToken={setToken} />
    }
    if (token !== undefined) {
        UserStatus.setStatus(true);
    }

    
    //UserStatus.switchStatus();
    //console.log("Enter", UserStatus.getStatus())
}

export default Enter;