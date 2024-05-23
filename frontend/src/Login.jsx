import axios from "axios"
import React, { useState } from "react"
import { toast } from "react-toastify";
import { useHistory } from "react-router-dom";
import { Mainpage } from "./Mainpage";


export const Login = (props) => { //parent componenet send some thing to children
    const [email, setEmail] = useState('')
    const [password, setpassword] = useState('')
    const mainPage = (email) => {
        props.onFormSwitch("mainpage", email)
    }

    const handleSubmit = (e) => {
        e.preventDefault(); // if we dont call page will be reloded and data will be lost
        const regobj = {
            email: email,
            password: password
        }
        fetch("http://localhost:8080/login", {
            method: "POST",
            headers: { 'content-type': 'application/json' },
            body: JSON.stringify(regobj),
            credentials : 'include'
        }).then((res) => {
            if (res.status === 200) {
                toast.success('Login Success');
                mainPage(email)
            } else {
                throw new Error(res.json().message);
            }
        }).catch((err) => {
            toast.error('Login Failed due to :' + err.message);
        });


    }


    return (

        <div className="auth-form-container">
            <form className="login-form" onSubmit={handleSubmit}>
                <label htmlFor="email"> email </label>
                <input value={email} onChange={(e) => setEmail(e.target.value)} type="email" placeholder="youremail@gmail.com" id="email" name="email" />
                <label htmlFor="password">password </label>
                <input value={password} onChange={(e) => setpassword(e.target.value)} type="password" placeholder="*******" id="password" name="password" />
                <p></p>
                <button style={{ backgroundColor: "#776854" }} type="submit">Log In</button>

            </form>
            <br></br>
            <text onClick={() => props.onFormSwitch("register")}>dont have account? click here to sign up</text>

        </div>

    )
}



