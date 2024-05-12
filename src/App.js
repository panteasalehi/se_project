import React, {useState} from "react"
import logo from './logo.svg';
import './App.css';
import { Login } from "./Login";
import { Makepost } from "./Makepost";
import { Register } from "./Register";
import { Mainpage } from "./Mainpage";
import {Myprofilepage} from "./Myprofilepage";
import { BrowserRouter,Routes, Route } from "react-router-dom";
import { ToastContainer } from "react-toastify";

function App() {

const [currentForm, setCurrentForm] = useState('login');
const[email , setEmail] = useState('')
const toggleForm = (formName , email) => {
  setCurrentForm(formName);
  setEmail(email);
}
return (
  <div className="App">
    <ToastContainer></ToastContainer>
    <BrowserRouter>
    <Routes>
        <Route path="/mainpage" element={<Mainpage />} /> 
    </Routes>
    </BrowserRouter>
    {
      
     currentForm === "login"? <Login onFormSwitch= {toggleForm} data = {email}/> :
     currentForm === "register"? <Register  onFormSwitch= {toggleForm} data = {email}/> :
     currentForm === "mainpage"? <Mainpage  onFormSwitch= {toggleForm} data = {email}/> :
     currentForm === "myprofilepage" ? <Myprofilepage onFormSwitch = {toggleForm} data = {email}/> : 
     <Makepost onFormSwitch = {toggleForm} data = {email} />
    }
  </div>
); 
}

export default App;
