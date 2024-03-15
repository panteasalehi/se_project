import axios from "axios"
import React, {useState} from "react"
import { toast } from "react-toastify";


export const Login = (props) => { //parent componenet send some thing to children
    const[email , setEmail] = useState('')
    const[pass , setPass] = useState('')

    const handleSubmit = (e) => {
        e.preventDefault(); // if we dont call page will be reloded and data will be lost

        fetch("http://localhost:8000/user/" + email).then((res) => {
                return res.json();
            }).then((resp) => {
                //console.log(resp)
                if (Object.keys(resp).length === 0) {
                    toast.error('Please Enter valid email');
                } else {
                    if (resp.pass === pass) {
                        toast.success('Success');
                        sessionStorage.setItem('email',email);
                        sessionStorage.setItem('userrole',resp.role);
                        console.log("success!")
                        //usenavigate('/')
                    }else{
                        toast.error('Please Enter valid credentials');
                    }
                }
            }).catch((err) => {
                toast.error('Login Failed due to :' + err.message);
            });
        

    }

    
    return (

        <div className="auth-form-container">
        <form className="login-form" onSubmit = {handleSubmit}>
            <label htmlFor="email"> email </label>
            <input value={email} onChange={(e)=>setEmail(e.target.value)} type = "email" placeholder="youremail@gmail.com" id = "email" name = "email" />
            <label htmlFor="password">password </label>
            <input value = {pass}  onChange={(e)=>setPass(e.target.value)} type = "password" placeholder="*******" id = "password" name = "password" />
            <button type = "submit">Log In</button>
        </form>

        <button className="link-button" onClick={ () => props.onFormSwitch("register")}>dont have account? click here to sign up</button>
        </div>
    )
  }
  
  

  