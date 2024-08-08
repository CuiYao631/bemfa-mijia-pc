
import './App.css';

import {BrowserRouter,Routes,Route} from "react-router-dom"
import Login from "./components/login/login";
import Home  from "./components/home/home";
import CreateUser  from "./components/createUser/createUser";



function App() {


    return (
        <div id="App">
            <BrowserRouter>
                <Routes>
                    <Route path="/" element={<Login/>}/>
                    <Route path="/home" element={<Home/>}/>
                    <Route path="/createUser" element={<CreateUser/>}/>
                </Routes>
            </BrowserRouter>
        </div>
    )
}

export default App
