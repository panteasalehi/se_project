import React, {useState} from "react"
import logo from './logo.svg';
import './App.css';
import { Login } from "./Login";
import { Paymentpage } from "./Paymentpage";
import { Makepost } from "./Makepost";
import { Register } from "./Register";
import { Support } from "./Support";
import { Notes } from "./Notes";
import { Mainpage } from "./Mainpage";
import {Myprofilepage} from "./Myprofilepage";
import {Postdetail} from "./Postdetail";
import {Myfavourites} from "./Myfavourites"
import {Myposts} from "./myPosts"
import { BrowserRouter,Routes, Route } from "react-router-dom";
import { ToastContainer } from "react-toastify";

function App() {

const [currentForm, setCurrentForm] = useState('login');
const[email , setEmail] = useState('')
const[id , setId] = useState('')
const[title , settitle] = useState('')
const toggleForm = (formName , email , id , title) => {
  setCurrentForm(formName);
  setEmail(email);
  setId(id);
  settitle(title);
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
      currentForm === "notes"? <Notes onFormSwitch= {toggleForm} data = {email}/> :
      currentForm === "register"? <Register  onFormSwitch= {toggleForm} data = {email}/> :
     currentForm === "support"? <Support  onFormSwitch= {toggleForm} data = {email}/> :
     currentForm === "paymentpage"? <Paymentpage  onFormSwitch= {toggleForm} data = {email}/> :
     currentForm === "mainpage"? <Mainpage  onFormSwitch= {toggleForm} data = {email}/> :
     currentForm === "Myfavourites"? <Myfavourites  onFormSwitch= {toggleForm} data = {email}/> :
     currentForm === "Myposts"? <Myposts  onFormSwitch= {toggleForm} data = {email}/> :
     currentForm === "myprofilepage" ? <Myprofilepage onFormSwitch = {toggleForm} data = {email}/> :
     currentForm === "postdetail" ? <Postdetail onFormSwitch = {toggleForm} data = {email} data2 = {id} data3 = {title}/> : 
     <Makepost onFormSwitch = {toggleForm} data = {email} />
    }
  </div>
); 
}

export default App;
