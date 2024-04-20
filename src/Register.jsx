import { Link, useNavigate } from "react-router-dom";
import { toast } from "react-toastify";
import React, {isValidElement, useState} from "react"




export const Register = (probs) => {
    const[id , setid] = useState('')
    const[pass , setPass] = useState('')
    const[name , setName] = useState('')
    let isused = true;

    const IsValidate = () => {
        let isproceed = true;
        
            if(/^[a-zA-Z0-9]+@[a-zA-Z0-9]+\.[A-Za-z]+$/.test(id)){
                    if(/^[a-zA-Z0-9]{8,20}$/.test(pass)){
                        
                    }
                    else{
                        isproceed = false;
                        toast.warning('Please enter the valid password')
                    }
                    }
                    else{
                        isproceed = false;
                        toast.warning('Please enter the valid email')
                    }
        
                     
        return isproceed;
    }


    
    const handleSubmit = (e) => {
        e.preventDefault(); // if we dont call page will be reloded and data will be lost
        let regobj = { id , pass };
            
            if(IsValidate()){

                fetch("http://localhost:8000/user/" + id).then((res) => {
                return res.json();
            }).then((resp) => {
                //console.log(resp)
                if (Object.keys(resp).length != 0) {
                   toast.error("email was used");
                } else {
                    
                }
            }).catch((err) => {

                fetch("http://localhost:8000/user", {
                    method: "POST",
                    headers: { 'content-type': 'application/json' },
                body: JSON.stringify(regobj)
                }).then((res) => {
                        toast.success('Registered successfully.')
                }).catch((err) => {
                    toast.error('Failed :' + err.message);
                });
            });
        
            
            
            }
        
    }











    return (

        <div className="auth-form-container">
        <form className="register-form"  onSubmit = {handleSubmit}>
            <label htmlFor="name"> full name </label>
            <input value={name} onChange={(e)=>setName(e.target.value)} type = "text" placeholder="name" id = "name" name = "name" />
            <label htmlFor="id"> email </label>
            <input value={id} onChange={(e)=>setid(e.target.value) } type = "email" placeholder="youremail@gmail.com" id = "id" name = "id" />
            <label htmlFor="password">password </label>
            <input value = {pass}  onChange={(e)=>setPass(e.target.value)} type = "password" placeholder="*******" id = "password" name = "password" />
            <p></p>
            <button type = "submit">Sign up</button>
        </form>

        <button className="link-button" onClick={() => probs.onFormSwitch("login")}>alreade have account? log in! </button>
        </div>
    )
  }
  
  

  