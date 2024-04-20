import React, {useState} from "react"
import './App.css';
import { Login } from "./Login";
import { Register } from "./Register";
import Home from "./Home"
import { BrowserRouter,Routes, Route } from "react-router-dom";
import { ToastContainer } from "react-toastify";

function App() {

const [currentForm, setCurrentForm] = useState('login');
const toggleForm = (formName) => {
  setCurrentForm(formName);
}
return (
  <div className="App">
    
    {
 
     currentForm === "login"? <Login onFormSwitch= {toggleForm}/> :currentForm === "register"? <Register  onFormSwitch= {toggleForm}/>:  <Login onFormSwitch= {toggleForm}/> 
    
    }
  </div>
); 
}

export default App;
