import NavbarWithLink from './components/Nav';
import Main from './pages/Main';
import Wallet from './pages/Wallet';
import Enter from './pages/Enter';
import Register from './pages/Register';
import { Route, Routes } from 'react-router-dom';
import './App.css';

function App() {

    return (
        <div className='app'>
            <NavbarWithLink />
            <div className='container'>
                <Routes>
                    <Route path="/main" element={<Main />} />
                    <Route path="/wallet" element={<Wallet />} />
                    <Route path="/enter" element={<Enter />} />
                    <Route path="/register" element={<Register />} />
                </Routes>
            </div>
        </div>
    );
};

export default App;