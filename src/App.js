import React, {useState} from "react"
import logo from './logo.svg';
import './App.css';
import { Login } from "./Login";
import { Register } from "./Register";
import Home from "./Home";
import { BrowserRouter,Routes, Route } from "react-router-dom";
import { ToastContainer } from "react-toastify";

function App() {

const [currentForm, setCurrentForm] = useState('login');
const toggleForm = (formName) => {
  setCurrentForm(formName);
}
return (
  <div className="App">
    <ToastContainer></ToastContainer>
    <BrowserRouter>
    <Routes>
    
    </Routes>
    </BrowserRouter>
    {
      
     currentForm === "login"? <Login onFormSwitch= {toggleForm}/> : <Register  onFormSwitch= {toggleForm}/>
    }
  </div>
); 
}

export default App;
