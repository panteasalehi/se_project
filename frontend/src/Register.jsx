import { Link, useNavigate } from "react-router-dom";
import { toast } from "react-toastify";
import React, { useState } from "react"

export const Register = (props) => {
    const [email, setid] = useState('')
    const [password, setPass] = useState('')
    const [name, setName] = useState('')






    const handleSubmit = (e) => {
        e.preventDefault(); // if we dont call page will be reloded and data will be lost
        let regobj = { email, password, name };

        console.log(JSON.stringify(regobj));
        fetch("http://45.147.97.39:8080/api/v1/signup/", {
            method: "POST",
            headers: { 'content-type': 'application/json' },
            body: JSON.stringify(regobj),
            credentials: 'include'
        }).then((res) => {
            if (res.ok) {
                toast.success('Registered successfully.');
                props.onFormSwitch("mainpage", email);
            } else {
                return res.json().then(err => {
                    throw new Error(err.message);
                });
            }
        }).catch((err) => {
            toast.error('Failed: ' + err.message);
        });

    }
    return (

        <div className="auth-form-container">
            <form className="register-form" onSubmit={handleSubmit}>
                <label htmlFor="name"> full name </label>
                <input value={name} onChange={(e) => setName(e.target.value)} type="text" placeholder="name" id="name" name="name" />
                <label htmlFor="id"> id </label>
                <input value={email} onChange={(e) => setid(e.target.value)} type="id" placeholder="youremail@gmail.com" id="id" name="email" />
                <label htmlFor="password">password </label>
                <input value={password} onChange={(e) => setPass(e.target.value)} type="password" placeholder="*******" id="password" name="password" />
                <br></br>
                <button style={{ backgroundColor: "#776854" }} type="submit">Sign up</button>
            </form>
            <br></br>
            <text onClick={() => props.onFormSwitch("login")}>alreade have account? log in! </text>
        </div>
    )
}



